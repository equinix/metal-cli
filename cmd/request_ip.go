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

package cmd

import (
	"fmt"
	"strconv"

	"github.com/packethost/packngo"
	"github.com/spf13/cobra"
)

var (
	ttype    string
	quantity int
	comments string
)

// requestIPCmd represents the requestIp command
var requestIPCmd = &cobra.Command{
	Use:   "request",
	Short: "Request an IP block",
	Long: `Example:

packet ip request --quantity [quantity] --facility [facility_code] --type [address_type]

	`,
	Run: func(cmd *cobra.Command, args []string) {
		req := &packngo.IPReservationRequest{
			Type:     ttype,
			Quantity: quantity,
			Facility: facility,
		}

		if comments != "" {
			req.Comments = comments
		}
		reservation, _, err := PacknGo.ProjectIPs.Request(projectID, req)
		if err != nil {
			fmt.Println("Client error:", err)
			return
		}

		data := make([][]string, 1)

		data[0] = []string{reservation.ID, reservation.Address, strconv.FormatBool(reservation.Public), reservation.Created}
		header := []string{"ID", "Address", "Public", "Created"}

		output(reservation, header, &data)
	},
}

func init() {
	requestIPCmd.Flags().StringVarP(&projectID, "project-id", "p", "", "UUID of the project")
	requestIPCmd.Flags().StringVarP(&ttype, "type", "t", "", "Address type public_ipv4 or global_ipv6")
	requestIPCmd.Flags().StringVarP(&facility, "facility", "f", "", "Code of the facility")
	requestIPCmd.Flags().IntVarP(&quantity, "quantity", "q", 0, "Number of IP addresses to reserve")

	requestIPCmd.MarkFlagRequired("project-id")
	requestIPCmd.MarkFlagRequired("type")
	requestIPCmd.MarkFlagRequired("quantity")
	requestIPCmd.MarkFlagRequired("facility")

	requestIPCmd.Flags().StringVarP(&comments, "comments", "c", "", "General comments")
}
