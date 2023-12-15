package interconnections

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
		Use:     `interconnections`,
		Aliases: []string{"conn"},
		Short:   "interconnections operations: create, get, update, delete",
		Long:    "Get information on Metro locations. For more information on https://deploy.equinix.com/developers/docs/metal/interconnections.",
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
		c.Retrieve(),
		c.Create(),
		c.Delete(),
		c.Update(),
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
