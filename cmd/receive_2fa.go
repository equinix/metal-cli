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

var app bool
var sms bool

// receive2faCmd represents the receive2fa command
var receive2faCmd = &cobra.Command{
	Use:   "receive",
	Short: "Receive two factor authentication token",
	Long: `Example:
Issue the token via SMS:
packet 2fa receive -s 

Issue the token via app:
packet 2fa receive -a

`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if sms == app {
			return fmt.Errorf("Either sms or app should be set")
		}

		if sms {
			_, err := PacknGo.TwoFactorAuth.ReceiveSms()
			if err != nil {
				return errors.Wrap(err, "Could not issue token via SMS")
			}

			fmt.Println("SMS token sent to your phone")
			return nil
		}

		otpURI, _, err := PacknGo.TwoFactorAuth.SeedApp()
		if err != nil {
			return errors.Wrap(err, "Could not get the OTP Seed URI")
		}

		data := make([][]string, 1)

		data[0] = []string{otpURI}
		header := []string{"OTP URI"}
		return output(otpURI, header, &data)
	},
}

func init() {
	twofaCmd.AddCommand(receive2faCmd)
	receive2faCmd.Flags().BoolVarP(&sms, "sms", "s", false, "Issues SMS otp token to user's phone")
	receive2faCmd.Flags().BoolVarP(&app, "app", "a", false, "Issues otp uri for auth application")
}
