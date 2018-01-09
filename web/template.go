package web

import (
	"bytes"
	"errors"
	"html/template"
	"io"
	"path/filepath"
	"strings"

	"github.com/gobuffalo/packr"
	"github.com/labstack/echo"
)

type Templates struct {
	templates *template.Template
}

var (
	views     = packr.NewBox("../views")
	templates *Templates
	checksums map[string]interface{}
)

func init() {
	// Compiling templates from go-bindata
	filenames := views.List()
	var t *template.Template
	if len(filenames) == 0 {
		panic(errors.New("template: no files views folder nor go generate not called"))
	}
	for _, filename := range filenames {
		s := views.String(filename)
		name := filepath.Base(filename)

		var tmpl *template.Template
		if t == nil {
			// First template
			t = template.New(name)
		}
		if name == t.Name() {
			tmpl = t
		} else {
			tmpl = t.New(name)
		}

		if _, err := tmpl.Parse(s); err != nil {
			panic(err)
		}
	}

	templates = &Templates{
		templates: t,
	}

	// Generating asset checksums
	checksums = map[string]interface{}{}
	assets.Walk(func(filename string, f packr.File) error {
		info, err := f.FileInfo()
		if err != nil {
			panic(err)
		}

		// /public/assets/javascripts/app.js => app_js
		name := strings.Replace(filepath.Base(info.Name()), ".", "_", -1)
		checksums[name] = info.ModTime().Unix()
		return nil
	})
}

// Render implements an interface.
func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	var buf bytes.Buffer
	if err := t.templates.ExecuteTemplate(&buf, name, data); err != nil {
		return err
	}

	return t.templates.ExecuteTemplate(w, "layout.tmpl", echo.Map{
		"yield":     template.HTML(buf.String()),
		"title":     "Risuto",
		"checksums": checksums,
	})
}
