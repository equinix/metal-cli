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

func (c *Client) Update() *cobra.Command {
	var projectID, name, paymentMethodID string
	// updateProjectCmd represents the updateProject command
	updateProjectCmd := &cobra.Command{
		Use:   `update -i <project_UUID> [-n <name>] [-m <payment_method_UUID>]`,
		Short: "Updates a project.",
		Long:  "Updates the specified project with a new name, a new payment method, or both.",
		Example: `  # Updates the specified project with a new name:
  metal project update -i $METAL_PROJECT_ID -n new-prod-cluster05
  
  # Updates the specified project with a new payment method:
  metal project update -i $METAL_PROJECT_ID -m e2fcdf91-b6dc-4d6a-97ad-b26a14b66839`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			req := &packngo.ProjectUpdateRequest{}
			if name != "" {
				req.Name = &name
			}

			if paymentMethodID != "" {
				req.PaymentMethodID = &paymentMethodID
			}
			p, _, err := c.ProjectService.Update(projectID, req)
			if err != nil {
				return errors.Wrap(err, "Could not update Project")
			}

			data := make([][]string, 1)

			data[0] = []string{p.ID, p.Name, p.Created}
			header := []string{"ID", "Name", "Created"}
			return c.Out.Output(p, header, &data)
		},
	}

	updateProjectCmd.Flags().StringVarP(&projectID, "id", "i", "", "The project's UUID. This flag is required, unless specified in the config created by metal init or set as METAL_PROJECT_ID environment variable.")
	updateProjectCmd.Flags().StringVarP(&name, "name", "n", "", "The new name for the project.")
	updateProjectCmd.Flags().StringVarP(&paymentMethodID, "payment-method-id", "m", "", "The UUID of the new payment method.")

	_ = updateProjectCmd.MarkFlagRequired("id")
	return updateProjectCmd
}
