// Copyright © 2018 Jasmin Gacic <jasmin@stackpointcloud.com>
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

package ips

import (
	"context"
	"fmt"
	"strconv"

	metal "github.com/equinix/equinix-sdk-go/services/metalv1"
	"github.com/spf13/cobra"
)

func (c *Client) Request() *cobra.Command {
	var (
		ttype     string
		quantity  int
		comments  string
		facility  string
		metro     string
		projectID string
		tags      []string
	)

	// requestIPCmd represents the requestIp command
	requestIPCmd := &cobra.Command{
		Use:   `request -p <project_id> -t <ip_address_type> -q <quantity> (-m <metro> | -f <facility>) [-f <flags>] [-c <comments>]`,
		Short: "Request a block of IP addresses.",
		Long:  "Requests either a block of public IPv4 addresses or global IPv4 addresses for your project in a specific metro or facility.",
		Example: `  # Requests a block of 4 public IPv4 addresses in Dallas:
  metal ip request -p $METAL_PROJECT_ID -t public_ipv4 -q 4 -m da`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true

			req := &metal.IPReservationRequestInput{
				Metro:    &metro,
				Tags:     tags,
				Quantity: int32(quantity),
				Type:     ttype,
				Facility: &facility,
			}

			requestIPReservationRequest := &metal.RequestIPReservationRequest{
				IPReservationRequestInput: req,
			}

			reservation, _, err := c.IPService.RequestIPReservation(context.Background(), projectID).RequestIPReservationRequest(*requestIPReservationRequest).Execute()
			if err != nil {
				return fmt.Errorf("Could not request IP addresses: %w", err)
			}

			data := make([][]string, 1)

			data[0] = []string{reservation.IPReservation.GetId(),
				reservation.IPReservation.GetAddress(),
				strconv.FormatBool(reservation.IPReservation.GetPublic()),
				reservation.IPReservation.CreatedAt.String()}
			header := []string{"ID", "Address", "Public", "Created"}

			return c.Out.Output(reservation, header, &data)
		},
	}

	requestIPCmd.Flags().StringVarP(&projectID, "project-id", "p", "", "The project's UUID. This flag is required, unless specified in the config created by metal init or set as METAL_PROJECT_ID environment variable.")
	requestIPCmd.Flags().StringVarP(&ttype, "type", "t", "", "The type of IP Address, either public_ipv4 or global_ipv4.")
	requestIPCmd.Flags().StringVarP(&facility, "facility", "f", "", "Code of the facility where the IP Reservation will be created")
	requestIPCmd.Flags().StringVarP(&metro, "metro", "m", "", "Code of the metro where the IP Reservation will be created")
	requestIPCmd.Flags().IntVarP(&quantity, "quantity", "q", 0, "Number of IP addresses to reserve.")
	requestIPCmd.Flags().StringSliceVar(&tags, "tags", nil, "Tag or Tags to add to the reservation, in a comma-separated list.")

	_ = requestIPCmd.MarkFlagRequired("project-id")
	_ = requestIPCmd.MarkFlagRequired("type")
	_ = requestIPCmd.MarkFlagRequired("quantity")

	requestIPCmd.Flags().StringVarP(&comments, "comments", "c", "", "General comments or description.")
	return requestIPCmd
}
