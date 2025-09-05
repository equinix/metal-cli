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
	"context"
	"fmt"

	metal "github.com/equinix/equinix-sdk-go/services/metalv1"
	"github.com/spf13/cobra"
)

func (c *Client) Retrieve() *cobra.Command {
	var userID string

	// retriveUserCmd represents the retriveUser command
	retrieveUserCmd := &cobra.Command{
		Use:   `get [-i <user_UUID>]`,
		Short: "Retrieves information about the current user or a specified user.",
		Long:  "Returns either information about the current user or information about a specified user. Specified user information is only available if that user shares a project with the current user.",
		Example: `  # Retrieves the current user's information:
  metal user get
  
  # Returns information on user 3b0795ba-fd0b-4a9e-83a7-063e5e12409d:
  metal user get --i 3b0795ba-fd0b-4a9e-83a7-063e5e12409d`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			var err error
			var user *metal.User
			if userID == "" {
				user, _, err = c.Service.FindCurrentUser(context.Background()).Execute()
				if err != nil {
					return fmt.Errorf("Could not get current User: %w", err)
				}
			} else {
				user, _, err = c.Service.FindUserById(context.Background(), userID).Include(c.Servicer.Includes(nil)).Exclude(c.Servicer.Excludes(nil)).Execute()
				if err != nil {
					return fmt.Errorf("Could not get Users: %w", err)
				}
			}

			data := make([][]string, 1)

			data[0] = []string{user.GetId(), user.GetFullName(), user.GetEmail(), user.CreatedAt.String()}
			header := []string{"ID", "Full Name", "Email", "Created"}

			return c.Out.Output(user, header, &data)
		},
	}

	retrieveUserCmd.Flags().StringVarP(&userID, "id", "i", "", "UUID of the user.")

	return retrieveUserCmd
}
