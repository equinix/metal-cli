package terraform

import (
	"bytes"
	_ "embed"
	"fmt"
	"html"
	"html/template"
	"path"
	metal "github.com/equinix-labs/metal-go/metal/v1"
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

	switch v := i.(type) {
	case *metal.Device:
		fmt.Printf("single device")
		f = deviceFormat
	case []metal.Device:
		fmt.Printf("devices")
		f = many(deviceFormat)
	case *metal.Project:
		f = projectFormat
	case []metal.Project:
		f = many(projectFormat)
	default:
		return nil, fmt.Errorf("%v is not compatible with terraform output", v)
	}

	addQuotesToString(i)

	tmpl, err := template.New("terraform").Funcs(template.FuncMap{
		"hrefToID": func(href string) string {
			return fmt.Sprintf("\"%s", path.Base(href))
		},
		"nullIfNilOrEmpty": nullIfNilOrEmpty,
	}).Parse(f)
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, i)
	if err != nil {
		return nil, err
	}
	result := html.UnescapeString(buf.String())
	return []byte(result), nil
}
