package cmd

import (
	"github.com/spf13/cobra"

	"github.com/equinix/metal-cli/internal/capacity"
	root "github.com/equinix/metal-cli/internal/cli"
	"github.com/equinix/metal-cli/internal/completion"
	"github.com/equinix/metal-cli/internal/devices"
	"github.com/equinix/metal-cli/internal/docs"
	"github.com/equinix/metal-cli/internal/emdocs"
	"github.com/equinix/metal-cli/internal/env"
	"github.com/equinix/metal-cli/internal/events"
	"github.com/equinix/metal-cli/internal/facilities"
	"github.com/equinix/metal-cli/internal/gateway"
	"github.com/equinix/metal-cli/internal/hardware"
	initPkg "github.com/equinix/metal-cli/internal/init"
	"github.com/equinix/metal-cli/internal/ips"
	"github.com/equinix/metal-cli/internal/metros"
	"github.com/equinix/metal-cli/internal/organizations"
	"github.com/equinix/metal-cli/internal/os"
	outputPkg "github.com/equinix/metal-cli/internal/outputs"
	"github.com/equinix/metal-cli/internal/plans"
	"github.com/equinix/metal-cli/internal/ports"
	"github.com/equinix/metal-cli/internal/projects"
	"github.com/equinix/metal-cli/internal/ssh"
	"github.com/equinix/metal-cli/internal/twofa"
	"github.com/equinix/metal-cli/internal/users"
	"github.com/equinix/metal-cli/internal/vlan"
	"github.com/equinix/metal-cli/internal/vrf"
)

// Cli struct
type Cli struct {
	MainCmd  *cobra.Command
	Outputer outputPkg.Outputer
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

func (cli *Cli) RegisterCommands(client *root.Client) {
	cli.MainCmd.AddCommand(
		docs.NewCommand(),
		emdocs.NewCommand(),
		completion.NewCommand(),

		env.NewClient(client, apiTokenEnvVar).NewCommand(),
		initPkg.NewClient(client).NewCommand(),

		capacity.NewClient(client, cli.Outputer).NewCommand(),
		metros.NewClient(client, cli.Outputer).NewCommand(),
		facilities.NewClient(client, cli.Outputer).NewCommand(),
		os.NewClient(client, cli.Outputer).NewCommand(),
		plans.NewClient(client, cli.Outputer).NewCommand(),
		events.NewClient(client, cli.Outputer).NewCommand(),
		users.NewClient(client, cli.Outputer).NewCommand(),
		vlan.NewClient(client, cli.Outputer).NewCommand(),
		hardware.NewClient(client, cli.Outputer).NewCommand(),
		devices.NewClient(client, cli.Outputer).NewCommand(),
		organizations.NewClient(client, cli.Outputer).NewCommand(),
		projects.NewClient(client, cli.Outputer).NewCommand(),
		ips.NewClient(client, cli.Outputer).NewCommand(),
		ssh.NewClient(client, cli.Outputer).NewCommand(),
		twofa.NewClient(client, cli.Outputer).NewCommand(),
		gateway.NewClient(client, cli.Outputer).NewCommand(),
		ports.NewClient(client, cli.Outputer).NewCommand(),
		vrf.NewClient(client, cli.Outputer).NewCommand(),
	)
}
