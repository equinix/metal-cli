package terraform

import (
	"bytes"
	_ "embed"
	"html/template"
	"path"

	"github.com/packethost/packngo"
)

var (
	//go:embed device.tf.gotmpl
	deviceFormat string

	//go:embed project.tf.gotmpl
	projectFormat string
)

func many(s string) string {
	return `{{range .}}` + s + `{{end}}`
}

func Marshal(i interface{}) ([]byte, error) {
	f := ""
	switch i.(type) {
	case *packngo.Device:
		f = deviceFormat
	case []packngo.Device:
		f = many(deviceFormat)
	case *packngo.Project:
		f = projectFormat
	case []packngo.Project:
		f = many(projectFormat)
	}

	tmpl, err := template.New("terraform").Funcs(template.FuncMap{
		"hrefToID": func(href string) string {
			return path.Base(href)
		},
	}).Parse(f)
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, i)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
