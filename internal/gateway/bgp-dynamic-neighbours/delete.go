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

	"github.com/pkg/errors"

	"github.com/spf13/cobra"
)

func (c *Client) Delete() *cobra.Command {
	var bgpNeighbourId string

	// deleteGwBgpCmd represents the delete gateway bgp dynamic neighbour command
	deleteGwBgpCmd := &cobra.Command{
		Use:     `delete`,
		Short:   "",
		Long:    "",
		Example: ``,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true

			_, _, err := c.Service.
				DeleteBgpDynamicNeighborById(context.Background(), bgpNeighbourId).
				Include(c.Servicer.Includes([]string{"created_by"})).
				Exclude(c.Servicer.Excludes([]string{})).
				Execute()
			if err != nil {
				return errors.WithMessage(err, "Could not create BGP Dynamic Neighbour")
			}

			fmt.Println("BGP Dynamic Neighbour deletion initiated. Please check 'metal gateway bgp-dynamic-neighbour get -i", bgpNeighbourId, "' for status")
			return nil
		},
	}

	deleteGwBgpCmd.Flags().StringVar(&bgpNeighbourId, "bgp-dynamic-neighbour-id", "", "")

	_ = deleteGwBgpCmd.MarkFlagRequired("bgp-neighbour-id")
	return deleteGwBgpCmd
}
