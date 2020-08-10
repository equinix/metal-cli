package cmd

import (
	"fmt"
	"os"

	outputPkg "github.com/packethost/packet-cli/internal/output"
	"github.com/packethost/packet-cli/internal/volume"
	"github.com/packethost/packngo"
	"github.com/spf13/cobra"
)

// Cli struct
type Cli struct {
	Client   *packngo.Client
	MainCmd  *cobra.Command
	Outputer outputPkg.Outputer
}

// VERSION build
var (
	Version string = "devel"
)

// NewCli struct
func NewCli() *Cli {
	var err error
	cli := &Cli{}
	cli.Client, err = packngo.NewClientWithBaseURL("Packet CLI", os.Getenv("PACKET_TOKEN"), nil, "https://api.packet.net/")
	if err != nil {
		fmt.Println("Client error:", err)
		return nil
	}

	rootCmd.DisableSuggestions = false
	cli.MainCmd = rootCmd
	cli.Outputer = &outputPkg.Standard{}
	cli.RegisterCommands()

	err = rootCmd.Execute()
	if err != nil {
		return nil
	}

	return cli
}

type Registrar interface {
	Register(*cobra.Command, outputPkg.Outputer)
}

func (cli *Cli) RegisterCommands() {
	c := cli.Client
	for _, reggie := range []Registrar{
		&volume.VolumeClient{VolumeService: c.Volumes, VolumeAttachmentService: c.VolumeAttachments},
	} {
		reggie.Register(cli.MainCmd, cli.Outputer)
	}
}

func output(in interface{}, header []string, data *[][]string) error {
	format := outputPkg.FormatText

	// TODO(displague) remove isJSON and isYaml globals
	switch {
	case isJSON:
		format = outputPkg.FormatJSON
	case isYaml:
		format = outputPkg.FormatYAML
	}
	output := &outputPkg.Standard{Format: format}
	return output.Output(in, header, data)
}

func outputMergingCells(in interface{}, header []string, data *[][]string) error {
	format := outputPkg.FormatText

	// TODO(displague) remove isJSON and isYaml globals
	switch {
	case isJSON:
		format = outputPkg.FormatJSON
	case isYaml:
		format = outputPkg.FormatYAML
	}
	output := &outputPkg.CellMerging{Format: format}
	return output.Output(in, header, data)
}
