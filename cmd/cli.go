package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ghodss/yaml"
	"github.com/olekukonko/tablewriter"
	"github.com/packethost/packngo"
	"github.com/spf13/cobra"
)

// Cli struct
type Cli struct {
	Client  *packngo.Client
	MainCmd *cobra.Command

	cfgFile     string
	isJSON      bool
	isYaml      bool
	packetToken string

	includes *[]string // nolint:unused
	excludes *[]string // nolint:unused
}

// VERSION build
var (
	Version string = "devel"
)

// NewCli struct
func NewCli() *Cli {
	cli := &Cli{}

	cli.MainCmd = rootCmd
	return cli
}

func output(in interface{}, header []string, data *[][]string) error {
	if !isJSON && !isYaml {

		table := tablewriter.NewWriter(os.Stdout)
		table.SetAutoWrapText(false)
		table.SetAlignment(tablewriter.ALIGN_LEFT)
		table.SetHeader(header)
		table.AppendBulk(*data)
		table.Render()
	} else if isJSON {
		output, err := json.MarshalIndent(in, "", "  ")
		if err != nil {
			return err
		}
		fmt.Println(string(output))
	} else if isYaml {
		output, err := yaml.Marshal(in)
		if err != nil {
			return err
		}
		fmt.Println(string(output))
	}
	return nil
}

func outputMergingCells(in interface{}, header []string, data *[][]string) error {
	if !isJSON && !isYaml {

		table := tablewriter.NewWriter(os.Stdout)
		table.SetAutoMergeCells(true)
		table.SetRowLine(true)
		table.SetHeader(header)
		table.AppendBulk(*data)
		table.Render()
	} else if isJSON {
		output, err := json.MarshalIndent(in, "", "  ")
		if err != nil {
			return err
		}
		fmt.Println(string(output))
	} else if isYaml {
		output, err := yaml.Marshal(in)
		if err != nil {
			return err
		}
		fmt.Println(string(output))
	}
	return nil
}
