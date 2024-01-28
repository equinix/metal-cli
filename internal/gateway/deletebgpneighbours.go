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

	"github.com/pkg/errors"

	"github.com/spf13/cobra"
)

func (c *Client) DeleteBgpNeighbors() *cobra.Command {
	var bgpNeighborId string

	// deleteGwBgpCmd represents the delete gateway bgp dynamic neighbor command
	deleteGwBgpCmd := &cobra.Command{
		Use:   `delete-bgp-dynamic-neighbors`,
		Short: "Deletes a BGP Dynamic Neighbor",
		Long:  "Deletes the BGP Dynamic Neighbor for the metal gateway with the specified ID",
		Example: `# Deletes a BGP Dynamic Neighbor using the bgp dynamic neighbor ID

	$ metal gateways delete-bgp-dynamic-neighbor --id "9c56fa1d-ec05-470b-a938-0e5dd6a1540c"

	BGP Dynamic Neighbor deletion initiated. Please check 'metal gateway get-bgp-dynamic-neighbor -i 9c56fa1d-ec05-470b-a938-0e5dd6a1540c for status
`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true

			_, _, err := c.VrfService.
				DeleteBgpDynamicNeighborById(context.Background(), bgpNeighborId).
				Include(c.Servicer.Includes([]string{"created_by"})).
				Exclude(c.Servicer.Excludes([]string{})).
				Execute()
			if err != nil {
				return errors.WithMessage(err, "Could not create BGP Dynamic Neighbor")
			}

			fmt.Println("BGP Dynamic Neighbor deletion initiated. Please check 'metal gateway get-bgp-dynamic-neighbor -i", bgpNeighborId, "' for status")
			return nil
		},
	}

	deleteGwBgpCmd.Flags().StringVar(&bgpNeighborId, "bgp-neighbor-id", "", "UUID of BGP Dynamic Neighbor.")

	_ = deleteGwBgpCmd.MarkFlagRequired("id")
	return deleteGwBgpCmd
}
