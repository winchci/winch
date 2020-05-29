package templates

import (
	"bytes"
	"context"
	"github.com/mitchellh/colorstring"
	"io/ioutil"
	"net/url"
	"path"
	"path/filepath"
	"strings"
	"text/template"
)

var t *template.Template

var funcs = map[string]interface{}{
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
	}}
