/*
winch - Universal Build and Release Tool
Copyright (C) 2020 Switchbit, Inc.

This program is free software: you can redistribute it and/or modify it under the terms of the GNU
General Public License as published by the Free Software Foundation, either version 3 of the License,
or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even
the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public
License for more details.

You should have received a copy of the GNU General Public License along with this program. If not,
see <https://www.gnu.org/licenses/>.
*/

package templates

import (
	"bytes"
	"context"
	"github.com/iancoleman/strcase"
	"io/ioutil"
	"net/url"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/mitchellh/colorstring"
)

var t *template.Template

var funcs = map[string]interface{}{
	"snake": strcase.ToSnake,
	"screaming_snake": strcase.ToScreamingSnake,
	"kebab": strcase.ToKebab,
	"screaming_kebab": strcase.ToScreamingKebab,
	"camel": strcase.ToCamel,
	"lower_camel": strcase.ToLowerCamel,
	"upper": strings.ToUpper,
	"lower": strings.ToLower,
	"withoutscheme": func(str string) string {
		u, err := url.Parse(str)
		if err != nil {
			return str
		}

		u.Scheme = ""
		return strings.TrimPrefix(u.String(), "//")
	},
}

// Execute the named template using the data as the initial context
func Execute(_ context.Context, name string, data interface{}) (string, error) {
	b := new(bytes.Buffer)
	err := t.ExecuteTemplate(b, name, data)
	if err != nil {
		return "", err
	}

	return colorstring.Color(b.String()), nil
}

func Load(dir string, path string) *template.Template {
	if strings.HasPrefix(path, "!") {
		return t.Lookup(path[1:])
	} else {
		return template.Must(template.ParseFiles(filepath.Join(dir, path)))
	}
}

func init() {
	t = template.New("cli").Funcs(funcs)

	d, err := Assets.Open("/")
	if err != nil {
		panic(err)
	}

	defer d.Close()

	tree, err := d.Readdir(-1)
	if err != nil {
		panic(err)
	}

	for _, f := range tree {
		if path.Ext(f.Name()) == ".tmpl" {
			fh, err := Assets.Open(f.Name())
			if err != nil {
				panic(err)
			}

			data, err := ioutil.ReadAll(fh)
			if err != nil {
				panic(err)
			}

			template.Must(t.New(f.Name()).Parse(string(data)))
		}
	}
}
