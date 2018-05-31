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

	"github.com/spf13/cobra"
)

var (
	assignmentID string
)

// ipCmd represents the ip command
var ipCmd = &cobra.Command{
	Use:   "ip",
	Short: "Retrieves IP addresses",
	Long: `Example:
	
	To get all IP addresses under a project:

	packet get ip --project-id [project_uuid] 
	
	To get an assigned IP address:

	packet get ip --assignement-id [assignement_uuid]
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if projectID != "" && assignmentID == "" {
			ips, _, err := PacknGo.ProjectIPs.List(projectID)
			if err != nil {
				fmt.Println("Client error:", err)
				return
			}

			data := make([][]string, len(ips))

			for i, ip := range ips {
				data[i] = []string{ip.ID, ip.Address, ip.Facility.Code, strconv.FormatBool(ip.Public), ip.Created}
			}
			header := []string{"ID", "Address", "Facility", "Public", "Created"}

			output(ips, header, &data)
		} else if projectID == "" && assignmentID != "" {
			ip, _, err := PacknGo.ProjectIPs.Get(assignmentID)
			if err != nil {
				fmt.Println("Client error:", err)
				return
			}

			data := make([][]string, 1)

			data[0] = []string{ip.ID, ip.Address, ip.Facility.Code, strconv.FormatBool(ip.Public), ip.Created}
			header := []string{"ID", "Address", "Facility", "Public", "Created"}

			output(ip, header, &data)
		} else if projectID != "" && assignmentID != "" {
			fmt.Println("Either project-id or assignement-id can be passed as parameters.")
		}
	},
}

func init() {
	getCmd.AddCommand(ipCmd)
	ipCmd.Flags().StringVarP(&projectID, "project-id", "p", "", "--project-id or -p [project_UUID]")
	ipCmd.Flags().StringVarP(&assignmentID, "assignment-id", "a", "", "--assignment-id or -a [assignment_UUID]")
}
