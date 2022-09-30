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

	"github.com/packethost/packngo"
	"github.com/spf13/cobra"
)

func (c *Client) Update() *cobra.Command {
	var sshKeyID string
	// updateSSHKeyCmd represents the updateSSHKey command
	updateSSHKeyCmd := &cobra.Command{
		Use:   `update -i <SSH-key_UUID> [-k <public_key>] [-l <label>]`,
		Short: "Updates an SSH key.",
		Long:  "Updates an SSH key with either a new public key, a new label, or both.",
		Example: `  # Updates SSH key 5cb96463-88fd-4d68-94ba-2c9505ff265e with a new public key: 
  metal ssh-key update -i 5cb96463-88fd-4d68-94ba-2c9505ff265e -k AAAAB3N...user@domain.com
  
  # Updates SSH key 5cb96463-88fd-4d68-94ba-2c9505ff265e with a new label:
  metal ssh-key update -i 5cb96463-88fd-4d68-94ba-2c9505ff265e -l test-machine-2`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			req := &packngo.SSHKeyUpdateRequest{}
			if key != "" {
				req.Key = &key
			}

			if label != "" {
				req.Label = &label
			}
			sshKey, _, err := c.Service.Update(sshKeyID, req)
			if err != nil {
				return fmt.Errorf("Could not update SSH Key: %w", err)
			}

			data := make([][]string, 1)

			data[0] = []string{sshKey.ID, sshKey.Label, sshKey.Created}
			header := []string{"ID", "Label", "Created"}

			return c.Out.Output(sshKey, header, &data)
		},
	}

	updateSSHKeyCmd.Flags().StringVarP(&sshKeyID, "id", "i", "", "UUID of the SSH key")
	updateSSHKeyCmd.Flags().StringVarP(&key, "key", "k", "", "Public SSH key string")
	updateSSHKeyCmd.Flags().StringVarP(&label, "label", "l", "", "Name of the SSH key")

	_ = updateSSHKeyCmd.MarkFlagRequired("project-id")
	return updateSSHKeyCmd
}
