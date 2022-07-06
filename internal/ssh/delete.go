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

package ssh

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) Delete() *cobra.Command {
	var (
		sshKeyID string
		force    bool
	)
	deleteSSHKey := func(id string) error {
		_, err := c.Service.Delete(id)
		if err != nil {
			return err
		}
		fmt.Println("SSH Key", id, "successfully deleted.")
		return nil
	}
	// deleteSSHKeyCmd represents the deleteSSHKey command
	deleteSSHKeyCmd := &cobra.Command{
		Use:   `delete --id <SSH-key_UUID> [--force]`,
		Short: "Deletes an SSH key.",
		Long:  "Deletes an SSH key with a confirmation prompt. To skip the confirmation use --force. Does not remove the SSH key from existing servers.",
		Example: `  # Deletes an SSH key, with confirmation:
  metal ssh-key delete -i 5cb96463-88fd-4d68-94ba-2c9505ff265e
  >
  ✔ Are you sure you want to delete SSH Key 5cb96463-88fd-4d68-94ba-2c9505ff265e: y
  
  # Deletes an SSH key, skipping confirmation:
  metal ssh-key delete -i 5cb96463-88fd-4d68-94ba-2c9505ff265e -f`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			if !force {
				prompt := promptui.Prompt{
					Label:     fmt.Sprintf("Are you sure you want to delete SSH Key %s: ", sshKeyID),
					IsConfirm: true,
				}

				_, err := prompt.Run()
				if err != nil {
					return nil
				}
			}

			return errors.Wrap(deleteSSHKey(sshKeyID), "Could not delete SSH Key")
		},
	}

	deleteSSHKeyCmd.Flags().StringVarP(&sshKeyID, "id", "i", "", "The UUID of the SSH key.")
	_ = deleteSSHKeyCmd.MarkFlagRequired("id")

	deleteSSHKeyCmd.Flags().BoolVarP(&force, "force", "f", false, "Skips confirmation for the deletion of the SSH key.")
	return deleteSSHKeyCmd
}
