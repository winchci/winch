package winch

import (
	"context"
	"fmt"
	"github.com/coreos/go-semver/semver"
	"github.com/google/go-github/v28/github"
	"github.com/mholt/archiver/v3"
	"github.com/switch-bit/winch/version"
	"golang.org/x/oauth2"
	"io/ioutil"
	"mime"
	"os"
	"path"
	"strings"
)

type GitHub struct {
	client *github.Client
	owner  string
	repo   string
}

func NewGitHub(ctx context.Context, url string) (*GitHub, error) {
	token := os.Getenv("GITHUB_TOKEN")
	if len(token) == 0 {
		return nil, fmt.Errorf("GITHUB_TOKEN not set")
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	url = strings.TrimPrefix(url, "https://github.com/")
	parts := strings.SplitN(url, "/", 3)

	return &GitHub{
		client: github.NewClient(tc),
		owner:  parts[0],
		repo:   parts[1],
	}, nil
}

func (g GitHub) GetTags(ctx context.Context) (map[string]string, error) {
	tags := make(map[string]string)

	t, _, err := g.client.Repositories.ListTags(ctx, g.owner, g.repo, nil)
	if err != nil {
		return tags, nil
	}

	for _, tag := range t {
		sha := tag.GetCommit().GetSHA()
		s := tag.GetName()
		newVersion := semver.New(s[1:])
		if _, ok := tags[sha]; !ok {
			tags[sha] = s
		} else {
			existingVersion := semver.New(tags[sha][1:])
			if existingVersion.LessThan(*newVersion) {
				tags[sha] = s
			}

		}
	}

	return tags, nil
}

func (g GitHub) GetCommits(ctx context.Context) ([]*Commit, error) {
	var commits []*Commit

	c, _, err := g.client.Repositories.ListCommits(ctx, g.owner, g.repo, &github.CommitsListOptions{
		ListOptions: github.ListOptions{
			PerPage: 100,
		},
	})
	if err != nil {
		return commits, nil
	}

	for _, commit := range c {
		commits = append(commits, &Commit{
			Hash:    commit.GetSHA(),
			When:    commit.GetCommit().GetAuthor().GetDate(),
			Message: ParseMessage(strings.TrimSpace(commit.GetCommit().GetMessage())),
		})
	}

	return commits, nil
}

func (g GitHub) CreateRelease(ctx context.Context, tag string, body string) (*github.RepositoryRelease, error) {
	rel, _, err := g.client.Repositories.GetReleaseByTag(ctx, g.owner, g.repo, tag)
	if err != nil {
		rel, _, err = g.client.Repositories.CreateRelease(ctx, g.owner, g.repo, &github.RepositoryRelease{
			TagName:         github.String(tag),
			TargetCommitish: nil,
			Name:            github.String(tag),
			Body:            github.String(body),
			Draft:           github.Bool(false),
			Prerelease:      github.Bool(false),
		})
		if err != nil {
			return nil, err
		}
	}

	return rel, nil
}

func (g GitHub) UploadAsset(ctx context.Context, relID int64, artifact string) error {
	if i, err := os.Stat(artifact); err == nil {
		if i.IsDir() {
			dir, err := ioutil.TempDir("", version.Name)
			if err != nil {
				return err
			}

			defer os.RemoveAll(dir)

			artifact = path.Join(dir, "artifacts.tgz")

			err = archiver.Archive([]string{i.Name()}, artifact)
			if err != nil {
				return err
			}
		}

		file, err := os.Open(artifact)
		if err != nil {
			return err
		}

		_, _, err = g.client.Repositories.UploadReleaseAsset(ctx, g.owner, g.repo, relID, &github.UploadOptions{
			Name:      artifact,
			MediaType: mime.TypeByExtension(artifact),
		}, file)
		if err != nil {
			return err
		}
	}

	return nil
}
