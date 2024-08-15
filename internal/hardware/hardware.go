package hardware

import (
	metal "github.com/equinix/equinix-sdk-go/services/metalv1"
	"github.com/equinix/metal-cli/internal/outputs"
	"github.com/spf13/cobra"
)

type Client struct {
	Servicer Servicer
	Service  metal.HardwareReservationsApiService
	Out      outputs.Outputer
}

func (c *Client) NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     `hardware-reservation`,
		Aliases: []string{"hardware-reservations", "hardware"},
		Short:   "Hardware reservation operations: get, move.",
		Long:    "Information and operations on Hardware Reservations. Provisioning specific devices from a reservation can be performed with the `metal device` command. Documentation is available on https://deploy.equinix.com/developers/docs/metal/deploy/reserved/.",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if root := cmd.Root(); root != nil {
				if root.PersistentPreRun != nil {
					root.PersistentPreRun(cmd, args)
				}
			}
			c.Service = *c.Servicer.MetalAPI(cmd).HardwareReservationsApi
		},
	}

	cmd.AddCommand(
		c.Retrieve(),
		c.Move(),
	)
	return cmd
}

type Servicer interface {
	MetalAPI(*cobra.Command) *metal.APIClient
	Format() outputs.Format
	Filters() map[string]string
	Includes(defaultIncludes []string) (incl []string)
	Excludes(defaultExcludes []string) (excl []string)
}

func NewClient(s Servicer, out outputs.Outputer) *Client {
	return &Client{
		Servicer: s,
		Out:      out,
	}
}
