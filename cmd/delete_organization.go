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

	"github.com/manifoldco/promptui"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// deleteOrganizationCmd represents the deleteOrganization command
var deleteOrganizationCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes an organization",
	Long: `Example:
	
packet organization delete -i [organization_UUID]

	`,
	RunE: func(cmd *cobra.Command, args []string) error {
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

func deleteOrganization(id string) error {
	_, err := apiClient.Organizations.Delete(id)
	if err != nil {
		return err
	}

	fmt.Println("Organization", organizationID, "has been deleted successfully.")
	return nil
}

func init() {
	deleteOrganizationCmd.Flags().StringVarP(&organizationID, "organization-id", "i", "", "UUID of the organization")
	_ = deleteOrganizationCmd.MarkFlagRequired("organization-id")
	deleteOrganizationCmd.Flags().BoolVarP(&force, "force", "f", false, "Force removal of the organization")
}
