package outputs

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ghodss/yaml"
	"github.com/olekukonko/tablewriter"
)

type Format string

const (
	FormatText Format = "text"
	FormatJSON Format = "json"
	FormatYAML Format = "yaml"
)

type Outputer interface {
	Output(interface{}, []string, *[][]string) error
	SetFormat(Format)
}

type Standard struct {
	Format Format
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

type CellMerging struct {
	Format Format
}

func (o *CellMerging) Output(in interface{}, header []string, data *[][]string) error {
	if o.Format == FormatJSON {
		return outputJSON(in)
	} else if o.Format == FormatYAML {
		return outputYAML(in)
	} else {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetAutoMergeCells(true)
		table.SetRowLine(true)
		table.SetHeader(header)
		table.AppendBulk(*data)
		table.Render()
		return nil
	}
}

func (o *CellMerging) SetFormat(fmt Format) {
	o.Format = fmt
}

// FormatSwitch returns the Format chosen between JSON, Yaml, and Text based on
// the supplied boolean values
func FormatSwitch(json, yaml bool) Format {
	switch {
	case json:
		return FormatJSON
	case yaml:
		return FormatYAML
	default:
		return FormatText
	}
}
