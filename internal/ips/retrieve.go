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

func (c *Client) Retrieve() *cobra.Command {
	var (
		assignmentID  string
		reservationID string
		projectID     string
	)

	// ipCmd represents the ip command
	retrieveIPCmd := &cobra.Command{
		Use:     `get -p <project_UUID> | -a <assignment_UUID> | -r <reservation_UUID>`,
		Aliases: []string{"list"},
		Short:   "Retrieves information about IP addresses, IP address reservations, and IP address assignments.",
		Long:    "Retrieves information about the IP addresses in a project, the IP addresses that are in a specified assignment, or the IP addresses that are in a specified reservation.",
		Example: `  # Lists all IP addresses in a project:
  metal ip get -p bb73aa19-c216-4ce2-a613-e5ca93732722 

  # Gets information about the IP addresses from an assignment ID:
  metal ip get -a bb526d47-8536-483c-b436-116a5fb72235

  # Gets the IP addresses from a reservation ID:
  metal ip get -r da1bb048-ea6e-4911-8ab9-b95635ca127a`,

		RunE: func(cmd *cobra.Command, args []string) error {
			if assignmentID != "" && reservationID != "" {
				return fmt.Errorf("either assignment-id or reservation-id can be set")
			}

			if assignmentID == "" && reservationID == "" && projectID == "" {
				return fmt.Errorf("either project-id or assignment-id or reservation-id must be set")
			}

			cmd.SilenceUsage = true
			inc := []string{}
			exc := []string{}
			types := []metal.FindIPReservationsTypesParameterInner{}

			if assignmentID != "" {
				ip, _, err := c.IPService.FindIPAddressById(context.Background(), assignmentID).Include(inc).Exclude(exc).Execute()
				if err != nil {
					return fmt.Errorf("Could not get Device IP address: %w", err)
				}

				data := make([][]string, 1)

				data[0] = []string{ip.IPAssignment.GetId(), ip.IPAssignment.GetAddress(), strconv.FormatBool(ip.IPAssignment.GetPublic()), ip.IPAssignment.CreatedAt.String()}
				header := []string{"ID", "Address", "Public", "Created"}

				return c.Out.Output(ip, header, &data)
			} else if reservationID != "" {
				ip, _, err := c.IPService.FindIPAddressById(context.Background(), reservationID).Include(inc).Exclude(exc).Execute()
				if err != nil {
					return fmt.Errorf("Could not get Reservation IP address: %w", err)
				}

				data := make([][]string, 1)
				code := ""
				metro := ""
				if ip.IPReservation.Facility != nil {
					code = ip.IPReservation.Facility.GetCode()
				}

				if ip.IPReservation.Metro != nil {
					metro = ip.IPReservation.Metro.GetCode()
				}

				data[0] = []string{ip.IPReservation.GetId(), ip.IPReservation.GetAddress(), metro, code, strconv.FormatBool(ip.IPReservation.GetPublic()), ip.IPReservation.CreatedAt.String()}
				header := []string{"ID", "Address", "Metro", "Facility", "Public", "Created"}

				return c.Out.Output(ip, header, &data)
			}

			// Don't have a paginator for this one
			resp, _, err := c.IPService.FindIPReservations(context.Background(), projectID).Include(inc).Exclude(exc).Types(types).Execute()
			if err != nil {
				return fmt.Errorf("Could not list Project IP addresses: %w", err)
			}
			ips := resp.IpAddresses
			data := make([][]string, len(ips))

			for i, ip := range ips {
				code := ""
				metro := ""
				if ip.IPReservation.Facility != nil {
					code = ip.IPReservation.Facility.GetCode()
				}
				if ip.IPReservation.Metro != nil {
					metro = ip.IPReservation.Metro.GetCode()
				}
				data[i] = []string{ip.IPReservation.GetId(), ip.IPReservation.GetAddress(), metro, code, strconv.FormatBool(ip.IPReservation.GetPublic()), ip.IPReservation.CreatedAt.String()}
			}
			header := []string{"ID", "Address", "Metro", "Facility", "Public", "Created"}

			return c.Out.Output(ips, header, &data)
		},
	}

	retrieveIPCmd.Flags().StringVarP(&projectID, "project-id", "p", "", "A Project UUID (METAL_PROJECT_ID).")
	retrieveIPCmd.Flags().StringVarP(&assignmentID, "assignment-id", "a", "", "UUID of an IP address assignment. When you assign an IP address to a server, it gets an assignment UUID.")
	retrieveIPCmd.Flags().StringVarP(&reservationID, "reservation-id", "r", "", "UUID of an IP address reservation.")
	return retrieveIPCmd
}
