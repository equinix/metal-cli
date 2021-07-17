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

package twofa

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) Enable() *cobra.Command {

	var token string
	var sms, app bool

	// enable2faCmd represents the enable2fa command
	var enable2faCmd = &cobra.Command{
		Use:   "enable",
		Short: "Enables two factor authentication",
		Long: `Example:

Enable two factor authentication via SMS
metal 2fa enable -s -c [code]

Enable two factor authentication via APP
metal 2fa enable -a -c [code]
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if sms == app {
				return fmt.Errorf("Either sms or app should be set")
			} else if sms {
				_, err := c.Service.EnableSms(token)
				if err != nil {
					return errors.Wrap(err, "Could not enable Two-Factor Authentication")
				}
			} else if app {
				_, err := c.Service.EnableApp(token)
				if err != nil {
					return errors.Wrap(err, "Could not enable Two-Factor Authentication")
				}
			}
			fmt.Println("Two factor authentication successfully enabled.")
			return nil
		},
	}

	enable2faCmd.Flags().BoolVarP(&sms, "sms", "s", false, "Issues SMS otp token to user's phone")
	enable2faCmd.Flags().BoolVarP(&app, "app", "a", false, "Issues otp uri for auth application")
	enable2faCmd.Flags().StringVarP(&token, "code", "c", "", "Two factor authentication code")
	_ = enable2faCmd.MarkFlagRequired("code")
	return enable2faCmd
}
