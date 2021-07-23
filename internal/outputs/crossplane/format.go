package crossplane

import (
	"bytes"
	"html/template"

	"github.com/packethost/packngo"
)

const (
	deviceFormat = `
---
apiVersion: server.metal.equinix.com/v1alpha2
kind: Device
metadata:
  name: {{.Hostname}}
  annotations:
    crossplane.io/external-name: {{.ID}}
spec:
  ## Late-Initialization will provide the current values
  ## so we don't specify them here.

  forProvider:
    hostname: {{.Hostname}}
    plan: {{.Plan.Slug}}
    metro: {{.Metro.Code}}
    operatingSystem: {{.OS.Slug}}
    billingCycle: {{.BillingCycle}}
    locked: {{.Locked}}
    tags: {{.Tags}}

  ## The "default" provider will be used unless named here.
  # providerConfigRef:
  #  name: equinix-metal-provider

  ## EM devices do not persist passwords in the API long. This
  ## optional secret will not get the root pass for devices > 24h
  # writeConnectionSecretToRef:
  #  name: crossplane-example
  #  namespace: crossplane-system

  ## Do not delete devices that have been imported.
  # reclaimPolicy: Retain
`
)

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
	tmpl, err := template.New("crossplane").Parse(f)
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
