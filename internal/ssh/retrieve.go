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
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) Retrieve() *cobra.Command {

	var (
		sshKeyID string
	)

	// retrieveSshKeysCmd represents the retrieveSshKeys command
	var retrieveSSHKeysCmd = &cobra.Command{
		Use: `get [-i <SSH-key_UUID>]`,
		Aliases: []string{"list"},
		Short: "Retrieves a list of SSH keys or a specified SSH key.",
		Long: "Retrieves a list of SSH keys associated with the current user's account or the details of single SSH key.",
		Example: `  # Retrieves the SSH keys of the current user: 
  metal ssh-key get
  
  # Returns the details of SSH key 5cb96463-88fd-4d68-94ba-2c9505ff265e:
  metal ssh-key get --id 5cb96463-88fd-4d68-94ba-2c9505ff265e`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			if sshKeyID == "" {
				sshKeys, _, err := c.Service.List()
				if err != nil {
					return errors.Wrap(err, "Could not list SSH Keys")
				}

				data := make([][]string, len(sshKeys))

				for i, s := range sshKeys {
					data[i] = []string{s.ID, s.Label, s.Created}
				}
				header := []string{"ID", "Label", "Created"}

				return c.Out.Output(sshKeys, header, &data)
			} else {
				sshKey, _, err := c.Service.Get(sshKeyID, nil)
				if err != nil {
					return errors.Wrap(err, "Could not get SSH Key")
				}

				data := make([][]string, 1)

				data[0] = []string{sshKey.ID, sshKey.Label, sshKey.Created}
				header := []string{"ID", "Label", "Created"}

				return c.Out.Output(sshKey, header, &data)
			}
		},
	}

	retrieveSSHKeysCmd.Flags().StringVarP(&sshKeyID, "id", "i", "", "The UUID of an SSH key.")
	return retrieveSSHKeysCmd
}
