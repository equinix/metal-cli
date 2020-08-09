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
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	sshKeyID string
)

// retrieveSshKeysCmd represents the retrieveSshKeys command
var retrieveSSHKeysCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieves a list of available SSH keys or a single SSH key",
	Long: `Example:

Retrieve all SSH keys: 
packet ssh-key get
  
Retrieve a specific SSH key:
packet ssh-key get --id [ssh-key_UUID] 

`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if sshKeyID == "" {
			sshKeys, _, err := PacknGo.SSHKeys.List()
			if err != nil {
				return errors.Wrap(err, "Could not list SSH Keys")
			}

			data := make([][]string, len(sshKeys))

			for i, s := range sshKeys {
				data[i] = []string{s.ID, s.Label, s.Created}
			}
			header := []string{"ID", "Label", "Created"}

			return output(sshKeys, header, &data)
		} else {
			sshKey, _, err := PacknGo.SSHKeys.Get(sshKeyID, nil)
			if err != nil {
				return errors.Wrap(err, "Could not get SSH Key")
			}

			data := make([][]string, 1)

			data[0] = []string{sshKey.ID, sshKey.Label, sshKey.Created}
			header := []string{"ID", "Label", "Created"}

			return output(sshKey, header, &data)
		}
	},
}

func init() {
	retrieveSSHKeysCmd.Flags().StringVarP(&sshKeyID, "id", "i", "", "UUID of the SSH key")
	retrieveSSHKeysCmd.PersistentFlags().BoolVarP(&isJSON, "json", "j", false, "JSON output")
	retrieveSSHKeysCmd.PersistentFlags().BoolVarP(&isYaml, "yaml", "y", false, "YAML output")
}
