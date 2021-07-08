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
	"fmt"
	"strconv"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) Retrieve() *cobra.Command {
	var (
		assignmentID  string
		reservationID string
		projectID     string
	)

	// ipCmd represents the ip command
	var retrieveIPCmd = &cobra.Command{
		Use:     "get",
		Aliases: []string{"list"},
		Short:   "Retrieves information about IP addresses, IP reservations and IP assignments",
		Long: `Example:
	
To get all IP addresses under a project:

metal ip get --project-id [project_UUID] 

To get IP addresses by assignment id:

metal ip get --assignment-id [assignment_UUID]

To get IP addresses by reservation ID:

metal ip get --reservation-id [reservation_UUID]

	`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if projectID != "" && assignmentID == "" && reservationID == "" {
				ips, _, err := c.ProjectService.List(projectID, nil)
				if err != nil {
					return errors.Wrap(err, "Could not list Project IP addresses")
				}

				data := make([][]string, len(ips))

				for i, ip := range ips {
					code := ""
					if ip.Facility != nil {
						code = ip.Facility.Code
					}
					data[i] = []string{ip.ID, ip.Address, code, strconv.FormatBool(ip.Public), ip.Created}
				}
				header := []string{"ID", "Address", "Facility", "Public", "Created"}

				return c.Out.Output(ips, header, &data)
			} else if projectID == "" && reservationID == "" && assignmentID != "" {
				ip, _, err := c.ProjectService.Get(assignmentID, nil)
				if err != nil {
					return errors.Wrap(err, "Could not get Device IP address")
				}

				data := make([][]string, 1)

				data[0] = []string{ip.ID, ip.Address, strconv.FormatBool(ip.Public), ip.Created}
				header := []string{"ID", "Address", "Public", "Created"}

				return c.Out.Output(ip, header, &data)
			} else if projectID == "" && assignmentID == "" && reservationID != "" {
				ip, _, err := c.ProjectService.Get(reservationID, nil)
				if err != nil {
					return errors.Wrap(err, "Could not get Reservation IP address")
				}

				data := make([][]string, 1)

				data[0] = []string{ip.ID, ip.Address, ip.Facility.Code, strconv.FormatBool(ip.Public), ip.Created}
				header := []string{"ID", "Address", "Facility", "Public", "Created"}

				return c.Out.Output(ip, header, &data)
			} else if (projectID != "" && (assignmentID != "" || reservationID != "")) || (projectID == "" && assignmentID == "" && reservationID == "") {
				return fmt.Errorf("Either project-id or assignment-id or reservation-id can be passed as parameters.")
			}
			return nil
		},
	}

	retrieveIPCmd.Flags().StringVarP(&projectID, "project-id", "p", "", "Project ID (METAL_PROJECT_ID)")
	retrieveIPCmd.Flags().StringVarP(&assignmentID, "assignment-id", "a", "", "UUID of the assignment")
	retrieveIPCmd.Flags().StringVarP(&reservationID, "reservation-id", "r", "", "UUID of the reservation")
	return retrieveIPCmd
}
