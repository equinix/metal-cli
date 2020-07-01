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

// updateProjectCmd represents the updateProject command
var updateSSHKeyCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates a project",
	Long: `Example:

packet ssh-key update --id [ssh-key_UUID] --key [new_key]

`,
	Run: func(cmd *cobra.Command, args []string) {
		req := &packngo.SSHKeyUpdateRequest{}
		if key != "" {
			req.Key = &key
		}

		if label != "" {
			req.Label = &label
		}
		sshKey, _, err := PacknGo.SSHKeys.Update(sshKeyID, req)
		if err != nil {
			fmt.Println("Client error:", err)
			return
		}

		data := make([][]string, 1)

		data[0] = []string{sshKey.ID, sshKey.Label, sshKey.Created}
		header := []string{"ID", "Label", "Created"}

		output(sshKey, header, &data)
	},
}

func init() {
	updateSSHKeyCmd.Flags().StringVarP(&sshKeyID, "id", "i", "", "UUID of the SSH key")
	updateSSHKeyCmd.Flags().StringVarP(&key, "key", "k", "", "Public SSH key string")
	updateSSHKeyCmd.Flags().StringVarP(&label, "label", "l", "", "Name of the SSH key")

	_ = updateSSHKeyCmd.MarkFlagRequired("project-id")
}
