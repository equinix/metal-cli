package vrf

import (
	metal "github.com/equinix-labs/metal-go/metal/v1"
	"github.com/equinix/metal-cli/internal/outputs"
	"github.com/spf13/cobra"
)

type Client struct {
	Servicer Servicer
	Service  metal.VRFsApiService
	Out      outputs.Outputer
}

func (c *Client) NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     `vrf`,
		Aliases: []string{"vrf"},
		Short:   "VRF operations : create, TODO: make other commands.",
		Long:    "Experimental VRF function",

		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if root := cmd.Root(); root != nil {
				if root.PersistentPreRun != nil {
					root.PersistentPreRun(cmd, args)
				}
			}
			c.Service = *c.Servicer.MetalAPI(cmd).VRFsApi
		},
	}

	cmd.AddCommand(
		c.Create(),
	)
	return cmd
}

type Servicer interface {
	MetalAPI(*cobra.Command) *metal.APIClient
}

func NewClient(s Servicer, out outputs.Outputer) *Client {
	return &Client{
		Servicer: s,
		Out:      out,
	}
}