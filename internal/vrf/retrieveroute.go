package vrf

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

func (c *Client) RetrieveRoute() *cobra.Command {
	var (
		vrfID string
	)

	// GetVrfRouteCmd represents the GetVrfRouteCmd command
	GetRouteCmd := &cobra.Command{
		Use:   "get-route [-i <VrfRoute-Id>]",
		Short: "Retrieve all routes in the VRF",
		Long:  "Retrieve all routes in the VRF",
		Example: ` #Retrieve all routes in the VRF
	# Retrieve all routes in the VRF
	metal vrf get-route -i bb526d47-8536-483c-b436-116a5fb72235`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			inc := []string{}
			exc := []string{}

			vrfRouteList, _, err := c.Service.GetVrfRoutes(context.Background(), vrfID).Include(c.Servicer.Includes(inc)).Exclude(c.Servicer.Excludes(exc)).Execute()
			if err != nil {
				return fmt.Errorf("could not Get all routes in the VRF: %w", err)
			}
			vrfRoute := vrfRouteList.GetRoutes()
			data := make([][]string, len(vrfRoute))
			header := []string{"ID", "Type", "Prefix", "NextHop", "Created"}
			for i, route := range vrfRoute {
				data[i] = []string{route.GetId(), string(route.GetType()), route.GetPrefix(), route.GetNextHop(), route.CreatedAt.String()}
			}
			return c.Out.Output(vrfRoute, header, &data)
		},
	}

	GetRouteCmd.Flags().StringVarP(&vrfID, "id", "i", "", "Specify the VRF UUID to list its associated routes configurations")

	// making them all required here
	_ = GetRouteCmd.MarkFlagRequired("id")

	return GetRouteCmd
}
