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
	"strconv"

	"github.com/packethost/packngo"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) BGPEnable() *cobra.Command {
	var (
		projectID, useCase, md5, deploymentType string
		asn                                     int
	)
	// updateProjectCmd represents the updateProject command
	updateProjectCmd := &cobra.Command{
		Use:   "bgp-enable",
		Short: "Enables BGP on a project",
		Long: `Example:

metal project bgp-enable --id [project_UUID] --asn [asn] --md5 [md5_secret] --use-case [use_case] --deployment-type [deployment_type]

`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			req := packngo.CreateBGPConfigRequest{
				UseCase:        useCase,
				Asn:            asn,
				DeploymentType: deploymentType,
				Md5:            md5,
			}

			p, err := c.BGPConfigService.Create(projectID, req)
			if err != nil {
				return errors.Wrap(err, "Could not update Project")
			}

			data := make([][]string, 1)

			data[0] = []string{projectID, useCase, strconv.Itoa(asn), deploymentType}
			header := []string{"ID", "UseCase", "ASN", "DeploymentType"}
			return c.Out.Output(p, header, &data)
		},
	}

	updateProjectCmd.Flags().StringVarP(&projectID, "id", "i", "", "Project ID (METAL_PROJECT_ID)")
	updateProjectCmd.Flags().StringVar(&useCase, "useCase", "", "Use case for BGP")
	updateProjectCmd.Flags().IntVar(&asn, "asn", 65000, "Local ASN")
	updateProjectCmd.Flags().StringVar(&deploymentType, "deploymentType", "", "Deployment type (local, global)")
	updateProjectCmd.Flags().StringVar(&md5, "md5", "", "BGP Password")

	_ = updateProjectCmd.MarkFlagRequired("id")
	_ = updateProjectCmd.MarkFlagRequired("asn")
	_ = updateProjectCmd.MarkFlagRequired("deployment_type")
	return updateProjectCmd
}
