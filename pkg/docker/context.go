/*
winch - Universal Build and Release Tool
Copyright (C) 2021 Ketch Kloud, Inc.

This program is free software: you can redistribute it and/or modify it under the terms of the GNU
General Public License as published by the Free Software Foundation, either version 3 of the License,
or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even
the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public
License for more details.

You should have received a copy of the GNU General Public License along with this program. If not,
see <https://www.gnu.org/licenses/>.
*/

package docker

import (
	"archive/tar"
	"compress/gzip"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/c2h5oh/datasize"
	"github.com/codeskyblue/dockerignore"
	"github.com/winchci/winch/version"
	"io"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type ContextProvider struct {
	context map[string]string
	dir     string
}

func NewContextProvider() (*ContextProvider, error) {
	dir, err := os.MkdirTemp("", version.Name)
	if err != nil {
		return nil, err
	}

	return &ContextProvider{
		dir:     dir,
		context: make(map[string]string),
	}, nil
}

func (p *ContextProvider) GetContext(contextPath string) (string, error) {
	if p, ok := p.context[contextPath]; ok {
		if fi, err := os.Stat(p); err == nil {
			fmt.Println("Reusing pre-built context")
			fmt.Printf("Context size: %s\n", datasize.ByteSize(fi.Size()).HumanReadable())
			return p, nil
		}
	}

	var err error
	var patterns []string

	if _, err = os.Stat(".dockerignore"); err == nil {
		patterns, err = dockerignore.ReadIgnoreFile(".dockerignore")
		if err != nil {
			return "", err
		}
	} else if _, err = os.Stat(".gitignore"); err == nil {
		patterns, err = dockerignore.ReadIgnoreFile(".gitignore")
		if err != nil {
			return "", err
		}
	} else {
		patterns, err = dockerignore.ReadIgnore(io.NopCloser(strings.NewReader(".git/")))
		if err != nil {
			return "", err
		}
	}

	fmt.Println("Archiving context", contextPath)
	contextArchive := path.Join(p.dir, hex.EncodeToString(md5.New().Sum([]byte(contextPath))), "context.tar.gz")

	err = os.MkdirAll(filepath.Dir(contextArchive), 0755)
	if err != nil {
		return "", err
	}

	out, err := os.Create(contextArchive)
	if err != nil {
		return "", err
	}
	defer out.Close()

	gw := gzip.NewWriter(out)
	defer gw.Close()
	tw := tar.NewWriter(gw)
	defer tw.Close()

	err = filepath.Walk(contextPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return err
		}

		skip, err := dockerignore.Matches(path, patterns)
		if err != nil || skip {
			return err
		}

		filename := filepath.Join(contextPath, path)

		// Open the file which will be written into the archive
		file, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer file.Close()

		// Create a tar Header from the FileInfo data
		header, err := tar.FileInfoHeader(info, info.Name())
		if err != nil {
			return err
		}

		// Use the full path as name (FileInfoHeader only takes the basename)
		// If we don't do this the directory structure would
		// not be preserved
		header.Name = filename

		// Write file header to the tar archive
		err = tw.WriteHeader(header)
		if err != nil {
			return err
		}

		// Copy file content to tar archive
		_, err = io.Copy(tw, file)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return "", err
	}

	if fi, err := os.Stat(contextArchive); err != nil {
		return "", err
	} else {
		fmt.Printf("Context size: %s\n", datasize.ByteSize(fi.Size()).HumanReadable())
	}
	fmt.Println(contextArchive)

	p.context[contextPath] = contextArchive

	return contextArchive, nil
}

func (p *ContextProvider) Close() error {
	return os.RemoveAll(p.dir)
}
