package gateway

import (
	"context"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func (c *Client) GetBgpNeighbors() *cobra.Command {
	var bgpNeighborID string

	// getGwayBgpDynamicNeighborCmd represents command to get Metal Gateway Dynamic Neighbor by ID.
	getGwayBgpDynamicNeighborCmd := &cobra.Command{
		Use:   `get-bgp-dynamic-neighbors`,
		Short: "Gets a BGP Dynamic Neighbor",
		Long:  "Gets the BGP Dynamic Neighbor with the specified ID",
		Example: `# Gets a BGP Dynamic Neighbor using the bgp dynamic neighbor ID

	$ metal gateways get-bgp-dynamic-neighbor --id "9c56fa1d-ec05-470b-a938-0e5dd6a1540c"
`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true

			n, _, err := c.VrfService.
				BgpDynamicNeighborsIdGet(context.Background(), bgpNeighborID).
				Include(c.Servicer.Includes([]string{"created_by"})).
				Exclude(c.Servicer.Excludes([]string{})).
				Execute()
			if err != nil {
				return fmt.Errorf("Could not create Metal Gateway: %w", err)
			}

			data := make([][]string, 1)

			data[0] = []string{n.GetId(), n.GetBgpNeighborRange(),
				strconv.Itoa(int(n.GetBgpNeighborAsn())), string(n.GetState()), n.GetCreatedAt().String()}
			header := []string{"ID", "Range", "ASN", "State", "Created"}

			return c.Out.Output(n, header, &data)
		},
	}

	getGwayBgpDynamicNeighborCmd.Flags().StringVar(&bgpNeighborID, "bgp-neighbor-id", "", "UUID of BGP Dynamic Neighbor ID.")
	_ = getGwayBgpDynamicNeighborCmd.MarkFlagRequired("id")
	return getGwayBgpDynamicNeighborCmd
}
