package vrf

import (
	"context"
	"fmt"

	"github.com/equinix/equinix-sdk-go/services/metalv1"
	"github.com/spf13/cobra"
)

func (c *Client) UpdateRoute() *cobra.Command {
	var (
		vrfID   string
		prefix  string
		nextHop string
		tags    []string
	)

	// CreateVrfRouteCmd represents the CreateVrfRouteCmd command
	UpdateRouteCmd := &cobra.Command{
		Use:   "update-route [-i <VrfRoute-Id>] [-p <Prefix>] [-n NextHop] [-t <tags> ]",
		Short: "Requests a VRF Route be redeployed/update across the network.",
		Long:  "Requests a VRF Route be redeployed/update across the network.",
		Example: ` #Requests a VRF Route be redeployed/update across the network.
	
	metal vrf update-route [-i <VrfID>] [-p <prefix>] [-n nextHop] [-t <tags> ]`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			inc := []string{}
			exc := []string{}

			vrfRouteUpdateInput := metalv1.VrfRouteUpdateInput{}

			if prefix != "" {
				vrfRouteUpdateInput.Prefix = &prefix
			}
			if nextHop != "" {
				vrfRouteUpdateInput.NextHop = &nextHop
			}
			if cmd.Flag("tags").Changed {
				vrfRouteUpdateInput.Tags = tags
			}

			vrfRoute, _, err := c.Service.UpdateVrfRouteById(context.Background(), vrfID).VrfRouteUpdateInput(vrfRouteUpdateInput).Include(c.Servicer.Includes(inc)).Exclude(c.Servicer.Excludes(exc)).Execute()
			if err != nil {
				return fmt.Errorf("could not update VrfRoute: %w", err)
			}

			data := make([][]string, 1)

			data[0] = []string{vrfRoute.GetId(), string(vrfRoute.GetType()), vrfRoute.GetPrefix(), vrfRoute.GetNextHop(), vrfRoute.CreatedAt.String()}
			header := []string{"ID", "Type", "Prefix", "NextHop", "Created"}

			return c.Out.Output(vrfRoute, header, &data)
		},
	}

	UpdateRouteCmd.Flags().StringVarP(&vrfID, "id", "i", "", "Specify the VRF UUID to update the associated route configurations.")
	UpdateRouteCmd.Flags().StringVarP(&prefix, "prefix", "p", "", "The IPv4 prefix for the route, in CIDR-style notation. For a static default route, this will always be '0.0.0.0/0'")
	UpdateRouteCmd.Flags().StringVarP(&nextHop, "nextHop", "n", "", "Name of the Virtual Routing and Forwarding")
	UpdateRouteCmd.Flags().StringSliceVarP(&tags, "tags", "t", []string{}, `updates the tags for the Virtual Routing and Forwarding --tags "tag1,tag2" OR --tags "tag1" --tags "tag2" (NOTE: --tags "" will remove all tags from the Virtual Routing and Forwarding`)

	// making them all required here
	_ = UpdateRouteCmd.MarkFlagRequired("id")

	return UpdateRouteCmd
}
