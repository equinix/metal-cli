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

package bgp_dynamic_neighbours

import (
	"context"
	"fmt"
	"strconv"

	"github.com/pkg/errors"

	metal "github.com/equinix/equinix-sdk-go/services/metalv1"
	"github.com/spf13/cobra"
)

func (c *Client) Create() *cobra.Command {
	var gatewayId, bgpNeighbourRange string
	var asn int32

	// createGwBgpCmd represents the creation of gateway bgp dynamic neighbour command
	createGwBgpCmd := &cobra.Command{
		Use:     `create`,
		Short:   "",
		Long:    "",
		Example: ``,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true

			// "192.168.1.0/25", int32(12345)
			if bgpNeighbourRange == "" {
				fmt.Println("Please provide BGP neighbour IP range")
				return nil
			}
			if asn == 0 {
				fmt.Println("Please provide BGP neighbour ASN")
				return nil
			}

			bgpNeighbour, _, err := c.Service.
				CreateBgpDynamicNeighbor(context.Background(), gatewayId).
				BgpDynamicNeighborCreateInput(*metal.NewBgpDynamicNeighborCreateInput(bgpNeighbourRange, asn)).
				Include(c.Servicer.Includes([]string{"created_by"})).
				Exclude(c.Servicer.Excludes([]string{})).
				Execute()
			if err != nil {
				return errors.WithMessage(err, "Could not create BGP Dynamic Neighbour")
			}

			data := make([][]string, 1)
			data[0] = []string{bgpNeighbour.GetId(), bgpNeighbour.GetBgpNeighborRange(),
				strconv.Itoa(int(bgpNeighbour.GetBgpNeighborAsn())), string(bgpNeighbour.GetState()), bgpNeighbour.GetCreatedAt().String()}
			header := []string{"ID", "Range", "ASN", "State", "Created"}

			return c.Out.Output(bgpNeighbour, header, &data)
		},
	}

	createGwBgpCmd.Flags().StringVar(&gatewayId, "gateway-id", "", "")
	createGwBgpCmd.Flags().StringVar(&bgpNeighbourRange, "bgp-neighbour-range", "", "")
	createGwBgpCmd.Flags().Int32Var(&asn, "asn", 0, "")

	_ = createGwBgpCmd.MarkFlagRequired("gateway-id")
	return createGwBgpCmd
}
