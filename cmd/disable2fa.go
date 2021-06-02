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

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// disable2faCmd represents the disable2fa command
var disable2faCmd = &cobra.Command{
	Use:   "disable",
	Short: "Disables two factor authentication",
	Long: `Example:

Disable two factor authentication via SMS
metal 2fa disable -s -t [token]

Disable two factor authentication via APP
metal 2fa disable -a -t [token]
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if sms == app {
			return fmt.Errorf("Either sms or app should be set")
		} else if sms {
			_, err := apiClient.TwoFactorAuth.DisableSms(token)
			if err != nil {
				return errors.Wrap(err, "Could not disable Two-Factor Authentication via SMS")
			}
		} else if app {
			_, err := apiClient.TwoFactorAuth.DisableApp(token)
			if err != nil {
				return errors.Wrap(err, "Could not disable Two-Factor Authentication via App")
			}
		}
		fmt.Println("Two factor authentication successfully disabled.")
		return nil
	},
}

func init() {
	twofaCmd.AddCommand(disable2faCmd)
	disable2faCmd.Flags().BoolVarP(&sms, "sms", "s", false, "Issues SMS otp token to user's phone")
	disable2faCmd.Flags().BoolVarP(&app, "app", "a", false, "Issues otp uri for auth application")
	disable2faCmd.Flags().StringVarP(&token, "token", "t", "", "Two factor authentication token")
	_ = disable2faCmd.MarkFlagRequired("token")
}
