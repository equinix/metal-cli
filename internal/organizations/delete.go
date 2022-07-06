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

package organizations

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) Delete() *cobra.Command {
	var (
		organizationID string
		force          bool
	)
	deleteOrganization := func(id string) error {
		_, err := c.Service.Delete(id)
		if err != nil {
			return err
		}

		fmt.Println("Organization", id, "has been deleted successfully.")
		return nil
	}

	// deleteOrganizationCmd represents the deleteOrganization command
	deleteOrganizationCmd := &cobra.Command{
		Use:   `delete -i <organization_UUID>`,
		Short: "Deletes an organization.",
		Long:  "Deletes an organization. You can not delete an organization that contains projects or has outstanding charges. Only organization owners can delete an organization.",
		Example: `  # Deletes an organization, with confirmation: 
  metal organization delete -i 3bd5bf07-6094-48ad-bd03-d94e8712fdc8
  >
  ✔ Are you sure you want to delete organization 3bd5bf07-6094-48ad-bd03-d94e8712fdc8: y
  
  # Deletes an organization, skipping confirmation:
  metal organization delete -i 3bd5bf07-6094-48ad-bd03-d94e8712fdc8 -f`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			if !force {
				prompt := promptui.Prompt{
					Label:     fmt.Sprintf("Are you sure you want to delete organization %s: ", organizationID),
					IsConfirm: true,
				}

				_, err := prompt.Run()
				if err != nil {
					return nil
				}
			}

			return errors.Wrap(deleteOrganization(organizationID), "Could not delete Organization")
		},
	}

	deleteOrganizationCmd.Flags().StringVarP(&organizationID, "organization-id", "i", "", "The UUID of the organization.")
	_ = deleteOrganizationCmd.MarkFlagRequired("organization-id")
	deleteOrganizationCmd.Flags().BoolVarP(&force, "force", "f", false, "Skips confirmation for the removal of the organization.")
	return deleteOrganizationCmd
}
