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

func (c *Client) Receive() *cobra.Command {

	var app, sms bool

	// receive2faCmd represents the receive2fa command
	var receive2faCmd = &cobra.Command{
		Use:   `receive (-s | -a)`,
		Short: "Generates a two-factor authentication token for use in enabling two-factor authentication on the current user's account.",
		Long:  "Generates a two-factor authentication token for use in enabling two-factor authentication on the current user's account. In order to use SMS, a phone number must be associated with the account to receive the code. If you are using an app, a URI for the application is returned.",
		Example: `  # Issue the token via SMS:
  metal 2fa receive -s 

  # Issue the token via app:
  metal 2fa receive -a`,

		RunE: func(cmd *cobra.Command, args []string) error {
			if sms == app {
				return fmt.Errorf("either sms or app should be set")
			}

			cmd.SilenceUsage = true
			if sms {
				_, err := c.Service.ReceiveSms()
				if err != nil {
					return errors.Wrap(err, "Could not issue token via SMS")
				}

				fmt.Println("SMS token sent to your phone")
				return nil
			}

			otpURI, _, err := c.Service.SeedApp()
			if err != nil {
				return errors.Wrap(err, "Could not get the OTP Seed URI")
			}

			data := make([][]string, 1)

			data[0] = []string{otpURI}
			header := []string{"OTP URI"}
			return c.Out.Output(otpURI, header, &data)
		},
	}

	receive2faCmd.Flags().BoolVarP(&sms, "sms", "s", false, "Issues SMS OTP token to the phone number associated with the current user account.")
	receive2faCmd.Flags().BoolVarP(&app, "app", "a", false, "Issues an OTP URI for an authentication application.")
	return receive2faCmd
}
