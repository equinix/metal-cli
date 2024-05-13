// Copyright © 2024 Equinix Metal Developers <support@equinixmetal.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package gateway

import (
	"context"
	"fmt"
	"strconv"

	"github.com/pkg/errors"

	metal "github.com/equinix/equinix-sdk-go/services/metalv1"
	"github.com/spf13/cobra"
)

func (c *Client) CreateBgpNeighbors() *cobra.Command {
	var gatewayId, bgpNeighborRange string
	var asn int64

	// createGwBgpCmd represents the creation of gateway bgp dynamic neighbor command
	createGwBgpCmd := &cobra.Command{
		Use:   `create-bgp-dynamic-neighbors`,
		Short: "Creates a BGP Dynamic Neighbor",
		Long:  "Creates the BGP Dynamic Neighbor for the metal gateway with the specified IP Range and ASN",
		Example: `# Create a BGP Dynamic Neighbor using ip range and asn for the metal gateway id

	metal gateways create-bgp-dynamic-neighbor --id "9c56fa1d-ec05-470b-a938-0e5dd6a1540c" --bgp-neighbor-range "10.70.43.226/29" --asn 65000
`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true

			// "192.168.1.0/25", int32(12345)
			if bgpNeighborRange == "" {
				fmt.Println("Please provide BGP neighbor IP range")
				return nil
			}
			if asn == 0 {
				fmt.Println("Please provide BGP neighbor ASN")
				return nil
			}

			bgpNeighbor, _, err := c.VrfService.
				CreateBgpDynamicNeighbor(context.Background(), gatewayId).
				BgpDynamicNeighborCreateInput(*metal.NewBgpDynamicNeighborCreateInput(bgpNeighborRange, asn)).
				Include(c.Servicer.Includes([]string{"created_by"})).
				Exclude(c.Servicer.Excludes([]string{})).
				Execute()
			if err != nil {
				return errors.WithMessage(err, "Could not create BGP Dynamic Neighbor")
			}

			data := make([][]string, 1)
			data[0] = []string{bgpNeighbor.GetId(), bgpNeighbor.GetBgpNeighborRange(),
				strconv.Itoa(int(bgpNeighbor.GetBgpNeighborAsn())), string(bgpNeighbor.GetState()), bgpNeighbor.GetCreatedAt().String()}
			header := []string{"ID", "Range", "ASN", "State", "Created"}

			return c.Out.Output(bgpNeighbor, header, &data)
		},
	}

	createGwBgpCmd.Flags().StringVarP(&gatewayId, "id", "i", "", "Metal Gateway ID for which the BGP Dynamic Neighbor to be created.")
	createGwBgpCmd.Flags().StringVar(&bgpNeighborRange, "bgp-neighbor-range", "", "BGP Dynamic Neighbor IP Range from gateway.")
	createGwBgpCmd.Flags().Int64Var(&asn, "asn", 0, "ASN for the BGP Dynamic Neighbor IP range.")

	_ = createGwBgpCmd.MarkFlagRequired("id")
	return createGwBgpCmd
}
