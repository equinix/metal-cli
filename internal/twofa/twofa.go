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
	"github.com/equinix/metal-cli/internal/outputs"

	"github.com/equinix/equinix-sdk-go/services/metalv1"
	"github.com/spf13/cobra"
)

type Client struct {
	Servicer     Servicer
	TwoFAService *metalv1.TwoFactorAuthApiService
	OtpService   *metalv1.OTPsApiService
	Out          outputs.Outputer
}

func (c *Client) NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     `2fa`,
		Aliases: []string{"tfa", "mfa", "totp"},
		Short:   "Two-factor Authentication operations: receive, enable, disable.",
		Long:    "Enable or disable two-factor authentication on your user account or receive an OTP token. More information is available at https://deploy.equinix.com/developers/docs/metal/identity-access-management/users/.",
		Args:    cobra.NoArgs,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if root := cmd.Root(); root != nil {
				if root.PersistentPreRun != nil {
					root.PersistentPreRun(cmd, args)
				}
			}
			c.TwoFAService = c.Servicer.MetalAPI(cmd).TwoFactorAuthApi
			c.OtpService = c.Servicer.MetalAPI(cmd).OTPsApi
		},
	}

	cmd.AddCommand(
		c.Receive(),
		c.Enable(),
		c.Disable(),
	)
	return cmd
}

type Servicer interface {
	MetalAPI(*cobra.Command) *metalv1.APIClient
	Filters() map[string]string
	Includes(defaultIncludes []string) (incl []string)
	Excludes(defaultExcludes []string) (excl []string)
}

func NewClient(s Servicer, out outputs.Outputer) *Client {
	return &Client{
		Servicer: s,
		Out:      out,
	}
}
