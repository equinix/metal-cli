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
	"github.com/packethost/packngo"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	label string
	key   string
)

func (c *Client) Create() *cobra.Command {
	createSSHKeyCmd := &cobra.Command{
		Use:   `create --key <public_key> --label <label>`,
		Short: "Adds an SSH key for the current user's account.",
		Long:  "Adds an SSH key for the current user's account. The key will then be added to the user's servers at provision time.",
		Example: ` # Adds a key labled "example-key" to the current user account.
  metal ssh-key create --key ssh-rsa AAAAB3N...user@domain.com --label example-key`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			req := packngo.SSHKeyCreateRequest{
				Label: label,
				Key:   key,
			}

			s, _, err := c.Service.Create(&req)
			if err != nil {
				return errors.Wrap(err, "Could not create SSHKey")
			}

			data := make([][]string, 1)

			data[0] = []string{s.ID, s.Label, s.Created}
			header := []string{"ID", "Label", "Created"}
			return c.Out.Output(s, header, &data)
		},
	}

	createSSHKeyCmd.Flags().StringVarP(&label, "label", "l", "", "Name or other user-friendly description of the SSH key.")
	createSSHKeyCmd.Flags().StringVarP(&key, "key", "k", "", "User's full SSH public key string.")

	_ = createSSHKeyCmd.MarkFlagRequired("label")
	_ = createSSHKeyCmd.MarkFlagRequired("key")
	return createSSHKeyCmd
}
