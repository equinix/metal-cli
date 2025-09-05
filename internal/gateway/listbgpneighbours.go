package gateway

import (
	"context"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func (c *Client) ListBgpNeighbors() *cobra.Command {
	var gatewayId string

	// createMetalGatewayCmd represents the createMetalGateway command
	createMetalGatewayCmd := &cobra.Command{
		Use:   `list-bgp-dynamic-neighbors`,
		Short: "Lists BGP Dynamic Neighbors for Metal Gateway",
		Long:  "Lists the BGP Dynamic Neighbor for the metal gateway with the specified gateway ID",
		Example: `# Lists BGP Dynamic Neighbor for the specified metal gateway ID

	$ metal gateways list-bgp-dynamic-neighbor --id "9c56fa1d-ec05-470b-a938-0e5dd6a1540c"
`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true

			n, _, err := c.VrfService.
				GetBgpDynamicNeighbors(context.Background(), gatewayId).
				Include(c.Servicer.Includes([]string{"created_by"})).
				Exclude(c.Servicer.Excludes([]string{})).
				Execute()
			if err != nil {
				return fmt.Errorf("Could not list BGP Dynamic Neighbors for Metal Gateway [%s]: %w", gatewayId, err)
			}

			data := make([][]string, len(n.GetBgpDynamicNeighbors()))
			for _, neighbor := range n.GetBgpDynamicNeighbors() {
				data[0] = []string{neighbor.GetId(), neighbor.GetBgpNeighborRange(),
					strconv.Itoa(int(neighbor.GetBgpNeighborAsn())), string(neighbor.GetState()), neighbor.GetCreatedAt().String()}
			}

			header := []string{"ID", "Range", "ASN", "State", "Created"}
			return c.Out.Output(n, header, &data)
		},
	}

	createMetalGatewayCmd.Flags().StringVarP(&gatewayId, "id", "i", "", "UUID of Metal Gateway.")

	_ = createMetalGatewayCmd.MarkFlagRequired("id")
	return createMetalGatewayCmd
}
