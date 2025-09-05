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

package gateway

import (
	"context"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func (c *Client) Delete() *cobra.Command {
	var (
		gwayID       string
		force        bool
		confirmation string
	)
	includes := []string{"virtual_network", "ip_reservation"}

	deleteGway := func(id string) error {
		_, _, err := c.Service.
			DeleteMetalGateway(context.Background(), id).
			Include(c.Servicer.Includes(includes)).
			Exclude(c.Servicer.Excludes(nil)).
			Execute()
		if err != nil {
			return err
		}
		fmt.Println("Metal Gateway", id, "successfully deleted.")
		return err
	}

	// deleteMetalGatewayCmd represents the deleteMetalGateway command
	deleteMetalGatewayCmd := &cobra.Command{
		Use:   `delete -i <metal_gateway_UUID> [-f]`,
		Short: "Deletes a Metal Gateway.",
		Long:  "Deletes the specified Gateway with a confirmation prompt. To skip the confirmation use --force.",
		Example: `  # Deletes a Gateway, with confirmation.
  metal gateway delete -i 77e6d57a-d7a4-4816-b451-cf9b043444e2
  >
  ✔ Are you sure you want to delete Metal Gateway 77e6d57a-d7a4-4816-b451-cf9b043444e2 [Y/N]: y

  # Deletes a Gateway, skipping confirmation.
  metal gateway delete -f -i 77e6d57a-d7a4-4816-b451-cf9b043444e2`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true

			if !force {
				fmt.Printf("Are you sure you want to delete Metal Gateway %s [Y/N]: ", gwayID)

				_, err := fmt.Scanln(&confirmation)
				if err != nil {
					fmt.Println("Error reading confirmation:", err)
					return nil
				}

				confirmation = strings.TrimSpace(strings.ToLower(confirmation))
				if confirmation != "yes" && confirmation != "y" {
					fmt.Println("Metal Gateway deletion cancelled.")
					return nil
				}
			}
			if err := deleteGway(gwayID); err != nil {
				return fmt.Errorf("could not delete Metal Gateway: %w", err)
			}
			return nil
		},
	}

	deleteMetalGatewayCmd.Flags().StringVarP(&gwayID, "id", "i", "", "UUID of the Gateway.")
	_ = deleteMetalGatewayCmd.MarkFlagRequired("id")
	deleteMetalGatewayCmd.Flags().BoolVarP(&force, "force", "f", false, "Skips confirmation for the removal of the Metal Gateway.")

	return deleteMetalGatewayCmd
}
