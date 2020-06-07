package actions

import (
	"fmt"
	"os"
)

type ActionServer interface {
	GetDownloadURL(ref *ActionRef) []string
}

type GenericHttpServer struct {}

func (g *GenericHttpServer) GetDownloadURL(ref *ActionRef) []string {
	root := fmt.Sprintf("%s%s/", ref.Protocol, ref.Server)
	return []string{
		root + fmt.Sprintf("%s/%s/%s", ref.Organization, ref.Repository, ref.Path),
		root + fmt.Sprintf("%s/%s/%s/winch-action.yml", ref.Organization, ref.Repository, ref.Path),
	}
}

type GitHubActionServer struct {}

func (g *GitHubActionServer) GetDownloadURL(ref *ActionRef) []string {
	gt := os.Getenv("GITHUB_TOKEN")
	if len(gt) > 0 {
		gt = gt + "@"
	}

	return []string{
		fmt.Sprintf("https://%sgithub.com/%s/winch-%s/archive/master.tar.gz", gt, ref.Organization, ref.Repository),
	}
}

var servers map[string]ActionServer

func GetActionServerForRef(ref *ActionRef) ActionServer {
	return servers[ref.Server]
}

func init() {
	servers = make(map[string]ActionServer)
	servers["github.com"] = new(GitHubActionServer)
}
