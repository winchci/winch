package actions

import (
	"context"
	"fmt"
	"github.com/mholt/archiver/v3"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

type Manager struct {
	actions map[string]*ActionDefinition
	dir string
}

func NewManager() *Manager {
	tempDir, err := os.UserCacheDir()
	if err != nil {
		return nil
	}

	tempDir = filepath.Join(tempDir, "winch")

	_ = os.MkdirAll(tempDir, 0755)

	return &Manager{
		actions: make(map[string]*ActionDefinition),
		dir:     tempDir,
	}
}

func (m *Manager) GetActions() []*ActionDefinition {
	var actions []*ActionDefinition
	for _, action := range m.actions {
		actions = append(actions, action)
	}
	return actions
}

func (m *Manager) Load(ctx context.Context, ref *ActionRef) (*ActionDefinition, error) {
	if len(ref.Tag) == 0 {
		ref.Tag = "master"
	}

	log.Printf("fetching %s...", ref.String())

	if a, ok := m.actions[ref.String()]; ok {
		log.Printf("found %s", ref.String())
		return a, nil
	}

	cached, err := filepath.Abs(path.Join(m.dir, ref.LocalString()))
	if err != nil {
		return nil, err
	}

	options := []string{
		path.Join(ref.Server, ref.Organization, ref.Repository, ref.Path, "winch-action.yml"),
		path.Join(ref.Server, ref.Organization, ref.Repository, "winch-action.yml"),
		path.Join(ref.Organization, ref.Repository, ref.Path, "winch-action.yml"),
		path.Join(ref.Organization, ref.Repository, "winch-action.yml"),
		path.Join(cached, "winch-action.yml"),
	}
	for _, option := range options {
		if _, err = os.Stat(option); err == nil {
			out, err := LoadActionDefinition(ctx, option)
			if err != nil {
				return nil, err
			}

			log.Printf("found %s", ref.String())
			m.actions[ref.String()] = out
			return out, nil
		}
	}

	file, err := ioutil.TempFile(m.dir, "winch_action_*.tar.gz")
	if err != nil {
		return nil, err
	}

	defer func(){
		file.Close()
		os.Remove(file.Name())
	}()

	server := GetActionServerForRef(ref)
	if server == nil {
		return nil, fmt.Errorf("no provider available to fetch %s", ref.String())
	}

	for _, url := range server.GetDownloadURL(ref) {
		var req *http.Request

		req, err = http.NewRequestWithContext(ctx, "GET", url, nil)
		if err != nil {
			continue
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			continue
		} else if resp.StatusCode != 200 {
			err = fmt.Errorf("could not fetch %s: %s", url, http.StatusText(resp.StatusCode))
			continue
		}

		_, err = io.Copy(file, resp.Body)
		if err != nil {
			continue
		}
		break
	}
	if err != nil {
		return nil, err
	}

	file.Close()

	err = archiver.Unarchive(file.Name(), m.dir)
	if err != nil {
		return nil, err
	}

	err = os.MkdirAll(cached, 0755)
	if err != nil {
		return nil, err
	}

	err = os.RemoveAll(cached)
	if err != nil {
		return nil, err
	}

	err = os.Rename(filepath.Join(m.dir, "winch-"+ref.Repository+"-"+ref.Tag), cached)
	if err != nil {
		return nil, err
	}

	cached, err = filepath.Abs(path.Join(cached, "winch-action.yml"))
	if err != nil {
		return nil, err
	}

	out, err := LoadActionDefinition(ctx, cached)
	if err != nil {
		return nil, err
	}

	log.Printf("found %s", ref.String())

	m.actions[ref.String()] = out
	return out, nil
}
