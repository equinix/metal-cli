package facilities

import (
	metal "github.com/equinix/equinix-sdk-go/services/metalv1"
	"github.com/equinix/metal-cli/internal/outputs"
	"github.com/spf13/cobra"
)

type Client struct {
	Servicer Servicer
	Service  *metal.FacilitiesApiService
	Out      outputs.Outputer
}

func (c *Client) NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     `facilities`,
		Aliases: []string{"facility"},
		Short:   "Facility operations: get.",
		Long:    "Information about specific facilities. Facility-level operations have mostly been replaced by Metros, but remains for backwards-compatibility. Documentation about facilities is available at https://deploy.equinix.com/developers/docs/metal/locations/facilities/.",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if root := cmd.Root(); root != nil {
				if root.PersistentPreRun != nil {
					root.PersistentPreRun(cmd, args)
				}
			}
			c.Service = c.Servicer.MetalAPI(cmd).FacilitiesApi
		},
	}

	cmd.AddCommand(
		c.Retrieve(),
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
