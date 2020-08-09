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

	"github.com/packethost/packngo"
	"github.com/spf13/cobra"
)

var (
	name            string
	paymentMethodID string
)

// projectCreateCmd represents the projectCreate command
var createProjectCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a project",
	Long: `Example:

packet project create --name [project_name]
  
  `,
	Run: func(cmd *cobra.Command, args []string) {
		req := packngo.ProjectCreateRequest{
			Name: name,
		}

		if organizationID != "" {
			req.OrganizationID = organizationID
		}

		if paymentMethodID != "" {
			req.PaymentMethodID = paymentMethodID
		}

		p, _, err := PacknGo.Projects.Create(&req)
		if err != nil {
			fmt.Println("Client error:", err)
			return
		}

		data := make([][]string, 1)

		data[0] = []string{p.ID, p.Name, p.Created}
		header := []string{"ID", "Name", "Created"}
		output(p, header, &data)
	},
}

func init() {
	createProjectCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the project")
	createProjectCmd.Flags().StringVarP(&organizationID, "organization-id", "o", "", "UUID of the organization")
	createProjectCmd.Flags().StringVarP(&paymentMethodID, "payment-method-id", "m", "", "UUID of the payment method")

	_ = createProjectCmd.MarkFlagRequired("name")
	createProjectCmd.PersistentFlags().BoolVarP(&isJSON, "json", "j", false, "JSON output")
	createProjectCmd.PersistentFlags().BoolVarP(&isYaml, "yaml", "y", false, "YAML output")
}
