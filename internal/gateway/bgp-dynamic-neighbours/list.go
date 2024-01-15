package bgp_dynamic_neighbours

import (
	"context"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func (c *Client) List() *cobra.Command {
	var gatewayId string

	// createMetalGatewayCmd represents the createMetalGateway command
	createMetalGatewayCmd := &cobra.Command{
		Use:     `list`,
		Short:   "",
		Long:    "",
		Example: ``,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true

			n, _, err := c.Service.
				GetBgpDynamicNeighbors(context.Background(), gatewayId).
				Include(c.Servicer.Includes([]string{"created_by"})).
				Exclude(c.Servicer.Excludes([]string{})).
				Execute()
			if err != nil {
				return fmt.Errorf("Could not list BGP Dynamic Neighbours for Metal Gateway [%s]: %w", gatewayId, err)
			}

			data := make([][]string, len(n.GetBgpDynamicNeighbors()))
			for _, neighbour := range n.GetBgpDynamicNeighbors() {
				data[0] = []string{neighbour.GetId(), neighbour.GetBgpNeighborRange(),
					strconv.Itoa(int(neighbour.GetBgpNeighborAsn())), string(neighbour.GetState()), neighbour.GetCreatedAt().String()}
			}

			header := []string{"ID", "Range", "ASN", "State", "Created"}
			return c.Out.Output(n, header, &data)
		},
	}

	createMetalGatewayCmd.Flags().StringVar(&gatewayId, "gateway-id", "", "")

	_ = createMetalGatewayCmd.MarkFlagRequired("gateway-id")
	return createMetalGatewayCmd
}
