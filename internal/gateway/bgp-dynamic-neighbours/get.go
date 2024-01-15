package bgp_dynamic_neighbours

import (
	"context"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func (c *Client) Get() *cobra.Command {
	var bgpNeighbourID string

	// getGwayBgpDynamicNeighbourCmd represents command to get Metal Gateway Dynamic Neighbour by ID.
	getGwayBgpDynamicNeighbourCmd := &cobra.Command{
		Use:     `get`,
		Short:   "",
		Long:    "",
		Example: ``,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true

			n, _, err := c.Service.
				BgpDynamicNeighborsIdGet(context.Background(), bgpNeighbourID).
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

	getGwayBgpDynamicNeighbourCmd.Flags().StringVar(&bgpNeighbourID, "bgp-dynamic-neighbour-id", "", "BGP Dynamic Neighbour ID. Ex: []")
	_ = getGwayBgpDynamicNeighbourCmd.MarkFlagRequired("bgp-dynamic-neighbour-id")
	return getGwayBgpDynamicNeighbourCmd
}
