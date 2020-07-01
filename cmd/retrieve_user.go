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

var (
	userID string
)

// retriveUserCmd represents the retriveUser command
var retriveUserCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieves information about the current user or a specified user",
	Long: `Example:

Retrieve the current user:
packet user get
  
Retrieve a specific user:
packet user get --id [user_UUID]

  `,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		var user *packngo.User
		if userID == "" {
			user, _, err = PacknGo.Users.Current()
			if err != nil {
				fmt.Println("Client error:", err)
				return
			}
		} else {
			user, _, err = PacknGo.Users.Get(userID, nil)
			if err != nil {
				fmt.Println("Client error:", err)
				return
			}
		}

		data := make([][]string, 1)

		data[0] = []string{user.ID, user.FullName, user.Email, user.Created}
		header := []string{"ID", "Full Name", "Email", "Created"}

		output(user, header, &data)
	},
}

func init() {
	retriveUserCmd.Flags().StringVarP(&userID, "id", "i", "", "UUID of the user")
	retriveUserCmd.PersistentFlags().BoolVarP(&isJSON, "json", "j", false, "JSON output")
	retriveUserCmd.PersistentFlags().BoolVarP(&isYaml, "yaml", "y", false, "YAML output")
}
