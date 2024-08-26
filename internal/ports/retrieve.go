// Copyright © 2022 Equinix Metal Developers <support@equinixmetal.com>
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

package ports

import (
	"context"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func (c *Client) Retrieve() *cobra.Command {
	var portID string
	// retrievePortCmd represents the retrievePort command
	retrievePortCmd := &cobra.Command{
		Use:     `get -i <port_UUID>`,
		Aliases: []string{},
		Short:   "Retrieves the details of the specified port.",
		Long:    "Retrieves the details of the specified port. Details of an port are only available to its members.",
		Example: `  # Retrieves details of a port:
  metal port get -i 3bd5bf07-6094-48ad-bd03-d94e8712fdc8`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			port, _, err := c.PortService.FindPortById(context.Background(), portID).
				Include(c.Servicer.Includes(nil)).
				Execute()
			if err != nil {
				return fmt.Errorf("Could not get Port: %w", err)
			}

			data := make([][]string, 1)

			data[0] = []string{port.GetId(), port.GetName(), string(port.GetType()), string(port.GetNetworkType()), port.Data.GetMac(), strconv.FormatBool(port.Data.GetBonded())}
			header := []string{"ID", "Name", "Type", "Network Type", "MAC", "Bonded"}

			return c.Out.Output(port, header, &data)
		},
	}

	retrievePortCmd.Flags().StringVarP(&portID, "port-id", "i", "", "The UUID of a port.")
	_ = retrievePortCmd.MarkFlagRequired("port-id")

	return retrievePortCmd
}
