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

// deleteSSHKeyCmd represents the deleteSSHKey command
var deleteSSHKeyCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes an SSH key",
	Long: `Example:

packet ssh-key delete --id [ssh-key_UUID]

`,
	RunE: func(cmd *cobra.Command, args []string) error {
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

func deleteSSHKey(id string) error {
	_, err := apiClient.SSHKeys.Delete(id)
	if err != nil {
		return err
	}
	fmt.Println("SSH Key", sshKeyID, "successfully deleted.")
	return nil
}

func init() {
	deleteSSHKeyCmd.Flags().StringVarP(&sshKeyID, "id", "i", "", "UUID of the SSH key")
	_ = deleteSSHKeyCmd.MarkFlagRequired("id")

	deleteSSHKeyCmd.Flags().BoolVarP(&force, "force", "f", false, "Force removal of the SSH key")
}
