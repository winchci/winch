package actions

import (
	"path/filepath"
	"strings"
)

const (
	DefaultBranch = "master"
)

type ActionRef struct {
	Protocol     string
	Server       string
	Organization string
	Repository   string
	Tag          string
	Path         string
}

func (a *ActionRef) String() string {
	var parts []string
	if len(a.Server) > 0 {
		parts = append(parts, a.Server)
	}
	if len(a.Organization) > 0 {
		parts = append(parts, a.Organization)
	}
	if len(a.Repository) > 0 {
		parts = append(parts, a.Repository)
	}
	if len(a.Path) > 0 {
		parts = append(parts, a.Path)
	}
	tag := ""
	if len(a.Tag) > 0 {
		tag = "@" + a.Tag
	}
	return a.Protocol + filepath.Join(parts...) + tag
}

func (a *ActionRef) LocalString() string {
	var parts []string
	if len(a.Server) > 0 {
		parts = append(parts, a.Server)
	}
	if len(a.Organization) > 0 {
		parts = append(parts, a.Organization)
	}
	if len(a.Repository) > 0 {
		parts = append(parts, a.Repository)
	}
	if len(a.Tag) > 0 {
		parts = append(parts, a.Tag)
	}
	if len(a.Path) > 0 {
		parts = append(parts, a.Path)
	}
	return filepath.Join(parts...)
}

func ParseActionRef(s string) *ActionRef {
	a := new(ActionRef)

	if strings.HasPrefix(s, "http://") {
		s = strings.TrimPrefix(s, "http://")
		a.Protocol = "http://"
	} else if strings.HasPrefix(s, "https://") {
		s = strings.TrimPrefix(s, "https://")
		a.Protocol = "https://"
	} else {
		a.Protocol = "https://"
	}

	parts := strings.Split(s, "@")
	if len(parts) > 1 {
		s = parts[0]
		a.Tag = parts[1]
	} else {
		a.Tag = DefaultBranch
	}

	parts = strings.Split(s, "/")
	if len(parts) > 0 && strings.Contains(parts[0], ".") {
		a.Server = parts[0]
		parts = parts[1:]
	} else {
		a.Server = "github.com"
	}
	if len(parts) > 0 {
		a.Organization = parts[0]
		parts = parts[1:]
	}
	if len(parts) > 0 {
		a.Repository = parts[0]
		parts = parts[1:]
	}
	if len(parts) > 0 {
		a.Path = parts[0]
		parts = parts[1:]
	}

	return a
}
