package virtualcircuit

import (
	metal "github.com/equinix/equinix-sdk-go/services/metalv1"
	"github.com/equinix/metal-cli/internal/outputs"
	"github.com/spf13/cobra"
)

type Client struct {
	Servicer Servicer
	Service  metal.InterconnectionsApiService
	Out      outputs.Outputer
}

func (c *Client) NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     `virtual-circuit`,
		Aliases: []string{"vc"},
		Short:   "virtual-circuit operations: create, get, update, delete",
		Long:    "For more information on https://deploy.equinix.com/developers/docs/metal/interconnections.",
		Args:    cobra.NoArgs,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if root := cmd.Root(); root != nil {
				if root.PersistentPreRun != nil {
					root.PersistentPreRun(cmd, args)
				}
			}
			c.Service = *c.Servicer.MetalAPI(cmd).InterconnectionsApi
		},
	}

	cmd.AddCommand(
		c.Create(),
		c.Retrieve(),
		c.Update(),
		c.Delete(),
	)

	return cmd
}

type Servicer interface {
	MetalAPI(*cobra.Command) *metal.APIClient
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
