// Copyright © 2022 Equinix Metal Developers <support@equinixmetal.com>
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

package projects

import (
	"context"
	"fmt"
	"strconv"

	metal "github.com/equinix-labs/metal-go/metal/v1"
	"github.com/spf13/cobra"
)

func (c *Client) BGPEnable() *cobra.Command {
	var (
		projectID, useCase, md5, deploymentType string
		asn                                     int
	)
	// bgpEnableProjectCmd represents the updateProject command
	bgpEnableProjectCmd := &cobra.Command{
		Use:   `bgp-enable --project-id <project_UUID> --deployment-type <deployment_type> [--asn <asn>] [--md5 <md5_secret>] [--use-case <use_case>]`,
		Short: "Enables BGP on a project.",
		Long:  `Enables BGP on a project.`,
		Example: `  # Enable BGP on project 50693ba9-e4e4-4d8a-9eb2-4840b11e9375:
  metal project bgp-enable --project-id 50693ba9-e4e4-4d8a-9eb2-4840b11e9375 --deployment-type local --asn 65000`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			req := metal.NewBgpConfigRequestInput(int32(asn), deploymentType)
			req.UseCase = &useCase
			req.Md5 = &md5

			p, err := c.BGPConfigService.RequestBgpConfig(context.Background(), projectID).BgpConfigRequestInput(*req).Execute()
			if err != nil {
				return fmt.Errorf("Could not update Project: %w", err)
			}

			data := make([][]string, 1)

			data[0] = []string{projectID, useCase, strconv.Itoa(asn), deploymentType}
			header := []string{"ID", "UseCase", "ASN", "DeploymentType"}
			return c.Out.Output(p, header, &data)
		},
	}

	bgpEnableProjectCmd.Flags().StringVarP(&projectID, "project-id", "p", "", "Project ID (METAL_PROJECT_ID)")
	bgpEnableProjectCmd.Flags().StringVar(&useCase, "use-case", "", "Use case for BGP")
	bgpEnableProjectCmd.Flags().IntVar(&asn, "asn", 65000, "Local ASN")
	bgpEnableProjectCmd.Flags().StringVar(&deploymentType, "deployment-type", "", "Deployment type (local, global)")
	bgpEnableProjectCmd.Flags().StringVar(&md5, "md5", "", "BGP Password")

	_ = bgpEnableProjectCmd.MarkFlagRequired("project-id")
	_ = bgpEnableProjectCmd.MarkFlagRequired("asn")
	_ = bgpEnableProjectCmd.MarkFlagRequired("deployment-type")
	return bgpEnableProjectCmd
}
