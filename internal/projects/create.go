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

package projects

import (
	"github.com/packethost/packngo"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) Create() *cobra.Command {
	var (
		name            string
		paymentMethodID string
		organizationID  string
	)

	// projectCreateCmd represents the projectCreate command
	createProjectCmd := &cobra.Command{
		Use:   `create -n <project_name> [-O <organization_UUID>] [-m <payment_method_UUID>]`,
		Short: "Creates a project.",
		Long:  "Creates a project with the specified name. If no organization is specified, the project is created in the current user's default organization. If no payment method is specified the organization's default payment method is used.",
		Example: `  # Creates a new project named dev-cluster02: 
  metal project create --name dev-cluster02
  
  # Creates a new project named dev-cluster03 in the specified organization with a payment method:
  metal project create -n dev-cluster03 -O 814b09ca-0d0c-4656-9de0-4ce65c6faf70 -m ab1fbdaa-8b25-4c3e-8360-e283852e3747`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			req := packngo.ProjectCreateRequest{
				Name: name,
			}

			if organizationID != "" {
				req.OrganizationID = organizationID
			}

			if paymentMethodID != "" {
				req.PaymentMethodID = paymentMethodID
			}

			p, _, err := c.ProjectService.Create(&req)
			if err != nil {
				return errors.Wrap(err, "Could not create Project")
			}

			data := make([][]string, 1)

			data[0] = []string{p.ID, p.Name, p.Created}
			header := []string{"ID", "Name", "Created"}
			return c.Out.Output(p, header, &data)
		},
	}

	createProjectCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the project")
	createProjectCmd.Flags().StringVarP(&organizationID, "organization-id", "O", "", "The UUID of the organization.")
	createProjectCmd.Flags().StringVarP(&paymentMethodID, "payment-method-id", "m", "", "The UUID of the payment method.")

	_ = createProjectCmd.MarkFlagRequired("name")
	return createProjectCmd
}
