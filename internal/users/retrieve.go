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

package users

import (
	"github.com/packethost/packngo"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) Retrieve() *cobra.Command {
	var (
		userID string
	)

	// retriveUserCmd represents the retriveUser command
	retrieveUserCmd := &cobra.Command{
		Use:   "get",
		Short: "Retrieves information about the current user or a specified user",
		Long: `Example:

Retrieve the current user:
metal user get
  
Retrieve a specific user:
metal user get --id [user_UUID]

  `,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			var err error
			var user *packngo.User
			if userID == "" {
				user, _, err = c.Service.Current()
				if err != nil {
					return errors.Wrap(err, "Could not get current User")
				}
			} else {
				user, _, err = c.Service.Get(userID, c.Servicer.ListOptions(nil, nil))
				if err != nil {
					return errors.Wrap(err, "Could not get Users")
				}
			}

			data := make([][]string, 1)

			data[0] = []string{user.ID, user.FullName, user.Email, user.Created}
			header := []string{"ID", "Full Name", "Email", "Created"}

			return c.Out.Output(user, header, &data)
		},
	}

	retrieveUserCmd.Flags().StringVarP(&userID, "id", "i", "", "UUID of the user")

	return retrieveUserCmd
}
