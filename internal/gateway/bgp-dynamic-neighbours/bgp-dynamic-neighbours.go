package bgp_dynamic_neighbours

import (
	metal "github.com/equinix/equinix-sdk-go/services/metalv1"
	"github.com/equinix/metal-cli/internal/outputs"

	"github.com/spf13/cobra"
)

type Client struct {
	Servicer Servicer
	Service  *metal.VRFsApiService
	Out      outputs.Outputer
}

func (c *Client) NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     `bgp-dynamic-neighbour`,
		Aliases: []string{"bgp-neighbour", "neighbours"},
		Short:   "Metal Gateway BGP Dynamic Neighbour operations: create, delete, and get, list",
		Long:    "",

		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if root := cmd.Root(); root != nil {
				if root.PersistentPreRun != nil {
					root.PersistentPreRun(cmd, args)
				}

				c.Service = c.Servicer.MetalAPI(cmd).VRFsApi
			}
		},
	}

	cmd.AddCommand(
		c.Create(),
		c.Get(),
		c.List(),
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
