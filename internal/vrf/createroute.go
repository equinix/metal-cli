package vrf

import (
	"context"
	"fmt"

	"github.com/equinix/equinix-sdk-go/services/metalv1"
	"github.com/spf13/cobra"
)

func (c *Client) CreateRoute() *cobra.Command {
	var (
		vrfID   string
		prefix  string
		nextHop string
		tags    []string
	)

	// CreateVrfRouteCmd represents the CreateVrfRouteCmd command
	createRouteCmd := &cobra.Command{
		Use:     "create-route [-i <VrfId>] [-p <Prefix>] [-n NextHop] [-t <tags> ]",
		Aliases: []string{"route"},
		Short:   "Create a route in a VRF. Currently only static default routes are supported.",
		Long:    "Create a route in a VRF. Currently only static default routes are supported.",
		Example: ` # Create a route in a VRF. Currently only static default routes are supported..
	
	metal vrf create-route [-i <VrfID>] [-p <prefix>] [-n nextHop] [-t <tags> ]`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			inc := []string{}
			exc := []string{}

			vrfRouteCreateInput := metalv1.VrfRouteCreateInput{
				Prefix:  prefix,
				NextHop: nextHop,
				Tags:    tags,
			}

			vrfRoute, _, err := c.Service.CreateVrfRoute(context.Background(), vrfID).VrfRouteCreateInput(vrfRouteCreateInput).Include(c.Servicer.Includes(inc)).Exclude(c.Servicer.Excludes(exc)).Execute()
			if err != nil {
				return fmt.Errorf("could not create VrfRoute: %w", err)
			}

			data := make([][]string, 1)

			// This output block below is probably incorrect but leaving it for now for testing later.
			data[0] = []string{vrfRoute.GetId(), string(vrfRoute.GetType()), vrfRoute.GetPrefix(), vrfRoute.GetNextHop(), vrfRoute.CreatedAt.String()}
			header := []string{"ID", "Type", "Prefix", "NextHop", "Created"}

			return c.Out.Output(vrfRoute, header, &data)
		},
	}

	createRouteCmd.Flags().StringVarP(&vrfID, "id", "i", "", "Specify the VRF UUID activate route configurations")
	createRouteCmd.Flags().StringVarP(&prefix, "prefix", "p", "", "The IPv4 prefix for the route, in CIDR-style notation. For a static default route, this will always be '0.0.0.0/0'")
	createRouteCmd.Flags().StringVarP(&nextHop, "nextHop", "n", "", "The IPv4 address within the VRF of the host that will handle this route")
	createRouteCmd.Flags().StringSliceVarP(&tags, "tags", "t", []string{}, `Adds the tags for the connection --tags="tag1,tag2".`)

	// making them all required here
	_ = createRouteCmd.MarkFlagRequired("id")
	_ = createRouteCmd.MarkFlagRequired("prefix")
	_ = createRouteCmd.MarkFlagRequired("nextHop")

	return createRouteCmd
}
