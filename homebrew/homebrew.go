package homebrew

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/winchci/winch/config"
	"github.com/winchci/winch/templates"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func WriteHomebrew(ctx context.Context, cfg *config.Config, t *config.HomebrewConfig, version string, file string) error {
	if !t.IsEnabled() {
		return nil
	}

	if len(file) == 0 {
		file = t.GetFile()
	}

	if len(file) == 0 {
		file = "formula.rb"
	}

	f, err := os.Create(filepath.Join(cfg.BasePath, file))
	if err != nil {
		return err
	}

	defer f.Close()

	vars := t.Variables
	if vars == nil {
		vars = make(map[string]interface{})
	}
	if _, ok := vars["Name"]; !ok {
		vars["Name"] = cfg.Name
	}
	if _, ok := vars["Description"]; !ok {
		vars["Description"] = cfg.Description
	}
	if _, ok := vars["Organization"]; !ok {
		vars["Organization"] = t.Organization
	}
	if _, ok := vars["Repository"]; !ok {
		vars["Repository"] = fmt.Sprintf("homebrew-%s", vars["Name"])
	}
	if _, ok := vars["Language"]; !ok {
		vars["Language"] = cfg.Language
	}
	if _, ok := vars["Homepage"]; !ok {
		vars["Homepage"] = cfg.Repository
	}
	if _, ok := vars["Version"]; !ok {
		vars["Version"] = version
	}
	if _, ok := vars["Install"]; !ok {
		vars["Install"] = t.Install
	}
	if _, ok := vars["Test"]; !ok {
		vars["Test"] = t.Test
	}
	if _, ok := vars["DependsOn"]; !ok {
		vars["DependsOn"] = t.DependsOn
	}
	if _, ok := vars["Asset"]; !ok {
		vars["Asset"] = t.Asset
	}
	if _, ok := vars["Url"]; !ok && len(t.Url) > 0 {
		vars["Url"] = t.Url
	}
	if _, ok := vars["Url"]; !ok {
		vars["Url"] = fmt.Sprintf("%s/releases/download/v%s/%s", cfg.Repository, vars["Version"], vars["Asset"])
	}

	var data []byte
	if strings.HasPrefix(vars["Url"].(string), "http") {
		fmt.Printf("homebrew: downloading '%s'\n", vars["Url"].(string))

		req, err := http.NewRequestWithContext(ctx, "GET", vars["Url"].(string), nil)
		if err != nil {
			return err
		}

		req.SetBasicAuth(t.Organization, os.Getenv("GITHUB_TOKEN"))

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		// Check server response
		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("homebrew: bad status: %s", resp.Status)
		}

		data, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
	} else {
		fmt.Printf("homebrew: opening file '%s'\n", vars["Url"].(string))
		data, err = ioutil.ReadFile(vars["Url"].(string))
		if err != nil {
			return err
		}
	}

	bytes := sha256.Sum256(data)
	sha := hex.EncodeToString(bytes[:])
	vars["Sha256"] = sha

	err = templates.Load(cfg.BasePath, t.GetTemplate()).Execute(f, vars)
	if err != nil {
		return err
	}

	f.Close()

	b, err := ioutil.ReadFile(filepath.Join(cfg.BasePath, file))
	if err != nil {
		return err
	}
	fmt.Println(string(b))

	return nil
}
