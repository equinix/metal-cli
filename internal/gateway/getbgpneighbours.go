package gateway

import (
	"context"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func (c *Client) GetBgpNeighbours() *cobra.Command {
	var bgpNeighbourID string

	// getGwayBgpDynamicNeighbourCmd represents command to get Metal Gateway Dynamic Neighbour by ID.
	getGwayBgpDynamicNeighbourCmd := &cobra.Command{
		Use:   `get-bgp-dynamic-neighbours`,
		Short: "Gets a BGP Dynamic Neighbour",
		Long:  "Gets the BGP Dynamic Neighbour for the metal gateway with the specified ID",
		Example: `# Gets a BGP Dynamic Neighbour using the bgp dynamic neighbour ID

	$ metal gateways get-bgp-dynamic-neighbour --id "9c56fa1d-ec05-470b-a938-0e5dd6a1540c"
`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true

			n, _, err := c.VrfService.
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

	getGwayBgpDynamicNeighbourCmd.Flags().StringVarP(&bgpNeighbourID, "id", "i", "", "BGP Dynamic Neighbour ID. Ex: []")
	_ = getGwayBgpDynamicNeighbourCmd.MarkFlagRequired("id")
	return getGwayBgpDynamicNeighbourCmd
}
