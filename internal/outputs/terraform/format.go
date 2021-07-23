package terraform

import (
	"bytes"
	"html/template"

	"github.com/packethost/packngo"
)

const deviceFormat = `
# terraform import metal_device.{{.Hostname}} {{.ID}}
resource "metal_device" "{{.Hostname}}" {
  plan = "{{.Plan.Slug}}"
  hostname = "{{.Hostname}}"
  billing_cycle = "{{.BillingCycle}}"
  metro = "{{.Metro.Code}}"
  operating_system = "{{.OS.Slug}}"
  project_id = "{{.Project.ID}}"

  tags = {{.Tags}}
}
`

func many(s string) string {
	return `{{range .}}` + s + `{{end}}`
}
func Marshal(i interface{}) ([]byte, error) {
	var f = ""
	switch i.(type) {
	case *packngo.Device:
		f = deviceFormat
	case []packngo.Device:
		f = many(deviceFormat)
	}
	tmpl, err := template.New("terraform").Parse(f)
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
