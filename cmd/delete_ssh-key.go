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
	"github.com/spf13/cobra"
)

// deleteProjectCmd represents the deleteProject command
var deleteSSHKeyCmd = &cobra.Command{
	Use:   "ssh-key",
	Short: "Deletes an SSH key",
	Long: `Example:
	packet delete ssh-key --id [ssh-key_UUID]`,
	Run: func(cmd *cobra.Command, args []string) {
		if !force {
			prompt := promptui.Prompt{
				Label:     fmt.Sprintf("Are you sure you want to delete project %s: ", projectID),
				IsConfirm: true,
			}

			_, err := prompt.Run()
			if err != nil {
				return
			}

			err = deleteSSHKey(sshKeyID)
			if err != nil {
				fmt.Println("Client error:", err)
				return
			}
		} else {
			err := deleteSSHKey(sshKeyID)
			if err != nil {
				fmt.Println("Client error:", err)
				return
			}
		}
	},
}

func deleteSSHKey(id string) error {
	_, err := PacknGo.SSHKeys.Delete(id)
	if err != nil {
		return err
	}
	fmt.Println("SSH Key", sshKeyID, "successfully deleted.")
	return nil
}

func init() {
	deleteCmd.AddCommand(deleteSSHKeyCmd)
	deleteSSHKeyCmd.Flags().StringVarP(&sshKeyID, "id", "i", "", "--id or -i [ssh-key_UUID]")
	deleteSSHKeyCmd.MarkFlagRequired("id")

	deleteSSHKeyCmd.Flags().BoolVarP(&force, "force", "f", false, "--force or -f")
}
