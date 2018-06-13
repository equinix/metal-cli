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

	"github.com/spf13/cobra"
)

// retrieveOrganizationCmd represents the retrieveOrganization command
var retrieveOrganizationCmd = &cobra.Command{
	Use:   "organization",
	Short: "Retrieves an organization or list of organizations",
	Long: `Example:
	
	To retireve list of all available organizations:
	packet get organization

	To retrieve a single organization:
	packet get organization -i [organization-id]
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if organizationID == "" {
			orgs, _, err := PacknGo.Organizations.List()
			if err != nil {
				fmt.Println("Client error:", err)
				return
			}

			data := make([][]string, len(orgs))

			for i, p := range orgs {
				data[i] = []string{p.ID, p.Name, p.Created}
			}
			header := []string{"ID", "Name", "Created"}

			output(orgs, header, &data)
		} else {
			org, _, err := PacknGo.Organizations.Get(organizationID)
			if err != nil {
				fmt.Println("Client error:", err)
				return
			}

			data := make([][]string, 1)

			data[0] = []string{org.ID, org.Name, org.Created}
			header := []string{"ID", "Name", "Created"}

			output(org, header, &data)
		}
	},
}

func init() {
	getCmd.AddCommand(retrieveOrganizationCmd)
	retrieveOrganizationCmd.Flags().StringVarP(&organizationID, "organization-id", "i", "", "--organization-id or -i")
}
