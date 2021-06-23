package cmd

import (
	"github.com/packethost/packngo"
	"github.com/spf13/cobra"

	"github.com/equinix/metal-cli/internal/capacity"
	"github.com/equinix/metal-cli/internal/cli"
	root "github.com/equinix/metal-cli/internal/cli"
	"github.com/equinix/metal-cli/internal/completion"
	"github.com/equinix/metal-cli/internal/docs"
	"github.com/equinix/metal-cli/internal/env"
	outputPkg "github.com/equinix/metal-cli/internal/outputs"
)

// Cli struct
type Cli struct {
	//	Client     *packngo.Client
	MainCmd    *cobra.Command
	Outputer   outputPkg.Outputer
	rootClient *cli.Client
}

// VERSION build
var (
	Version string = "devel"
)

const (
	consumerToken  = "Equinix Metal CLI"
	apiTokenEnvVar = "METAL_AUTH_TOKEN"
	apiURL         = "https://api.equinix.com/metal/v1/"
)

// NewCli struct
func NewCli() *Cli {
	cli := &Cli{
		Outputer: &outputPkg.Standard{},
	}

	rootClient := root.NewClient(consumerToken, apiURL, Version)
	rootClient.Init()
	rootCmd := rootClient.NewCommand()
	rootCmd.DisableSuggestions = false
	cli.MainCmd = rootCmd

	cli.RegisterCommands(rootClient)

	cobra.OnInitialize(
		func() {
			cli.Outputer.SetFormat(rootClient.Format())
		},
	)
	return cli
}

func (cli *Cli) API() *packngo.Client {
	return cli.rootClient.API()
}

type Registrar interface {
	NewCommand() *cobra.Command
}

func (cli *Cli) RegisterCommands(client *root.Client) {
	cli.MainCmd.AddCommand(
		docs.NewCommand(),
		completion.NewCommand(),

		env.NewClient(client, apiTokenEnvVar).NewCommand(),

		capacity.NewClient(client, cli.Outputer).NewCommand(),
		/*
			devices.NewClient(c, cli.Outputer).NewCommand(),
			events.NewClient(c, cli.Outputer).NewCommand(),
			facilities.NewClient(c, cli.Outputer).NewCommand(),
			hwReservations.NewClient(c, cli.Outputer).NewCommand(),
			ips.NewClient(c, cli.Outputer).NewCommand(),
			metros.NewClient(c, cli.Outputer).NewCommand(),
			operatingSystems.NewClient(c, cli.Outputer).NewCommand(),
			organizations.NewClient(c, cli.Outputer).NewCommand(),
			plans.NewClient(c, cli.Outputer).NewCommand(),
			projects.NewClient(c, cli.Outputer).NewCommand(),
			sshKeys.NewClient(c, cli.Outputer).NewCommand(),
			twofa.NewClient(c, cli.Outputer).NewCommand(),
			users.NewClient(c, cli.Outputer).NewCommand(),
			virtualNetworks.NewClient(c, cli.Outputer).NewCommand(),
		*/
	)
}
