package outputs

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"sigs.k8s.io/yaml"

	"github.com/equinix/metal-cli/internal/outputs/crossplane"
	"github.com/equinix/metal-cli/internal/outputs/terraform"
)

type Format string

const (
	FormatTable      Format = "table"
	FormatJSON       Format = "json"
	FormatYAML       Format = "yaml"
	FormatTerraform  Format = "tf"
	FormatCrossplane Format = "crossplane"
)

type Outputer interface {
	Output(interface{}, []string, *[][]string) error
	SetFormat(Format)
}

type Standard struct {
	Format Format
}

func outputCrossplane(in interface{}) error {
	output, err := crossplane.Marshal(in)
	if err != nil {
		return err
	}
	fmt.Println(string(output))
	return nil
}

func outputTerraform(in interface{}) error {
	output, err := terraform.Marshal(in)
	if err != nil {
		return err
	}
	fmt.Println(string(output))
	return nil
}

func outputJSON(in interface{}) error {
	output, err := json.MarshalIndent(in, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(output))
	return nil
}

func outputYAML(in interface{}) error {
	output, err := yaml.Marshal(in)
	if err != nil {
		return err
	}
	fmt.Println(string(output))
	return nil
}

func (o *Standard) Output(in interface{}, header []string, data *[][]string) error {
	if o.Format == FormatJSON {
		return outputJSON(in)
	} else if o.Format == FormatYAML {
		return outputYAML(in)
	} else if o.Format == FormatTerraform {
		return outputTerraform(in)
	} else if o.Format == FormatCrossplane {
		return outputCrossplane(in)
	} else {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetAutoWrapText(false)
		table.SetAlignment(tablewriter.ALIGN_LEFT)
		table.SetHeader(header)
		table.AppendBulk(*data)
		table.Render()
		return nil
	}
}

func (o *Standard) SetFormat(fmt Format) {
	o.Format = fmt
}
