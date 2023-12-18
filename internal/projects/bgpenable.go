package projects

import (
	"context"
	"fmt"
	"strconv"

	metal "github.com/equinix/equinix-sdk-go/services/metalv1"
	"github.com/spf13/cobra"
)

// BgpConfig holds the configuration for the BGP enable command
type BgpConfig struct {
	ProjectID      string
	UseCase        string
	MD5            string
	DeploymentType string
	ASN            int
}

func (c *Client) BGPEnable() *cobra.Command {
	config := BgpConfig{}

	bgpEnableProjectCmd := &cobra.Command{
		Use:     `bgp-enable --project-id <project_UUID> --deployment-type <deployment_type> [--asn <asn>] [--md5 <md5_secret>] [--use-case <use_case>]`,
		Short:   "Enables BGP on a project.",
		Long:    `Enables BGP on a project.`,
		Example: `  metal project bgp-enable --project-id 50693ba9-e4e4-4d8a-9eb2-4840b11e9375 --deployment-type local --asn 65000`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return enableBGP(c, &config)
		},
	}

	bgpEnableProjectCmd.Flags().StringVarP(&config.ProjectID, "project-id", "p", "", "Project ID (METAL_PROJECT_ID)")
	bgpEnableProjectCmd.Flags().StringVar(&config.UseCase, "use-case", "", "Use case for BGP")
	bgpEnableProjectCmd.Flags().IntVar(&config.ASN, "asn", 65000, "Local ASN")
	bgpEnableProjectCmd.Flags().StringVar(&config.DeploymentType, "deployment-type", "", "Deployment type (local, global)")
	bgpEnableProjectCmd.Flags().StringVar(&config.MD5, "md5", "", "BGP Password")

	_ = bgpEnableProjectCmd.MarkFlagRequired("project-id")
	_ = bgpEnableProjectCmd.MarkFlagRequired("asn")
	_ = bgpEnableProjectCmd.MarkFlagRequired("deployment-type")

	return bgpEnableProjectCmd
}

func enableBGP(c *Client, config *BgpConfig) error {
	req := metal.BgpConfigRequestInput{
		UseCase:        &config.UseCase,
		Asn:            int32(config.ASN),
		DeploymentType: metal.BgpConfigRequestInputDeploymentType(config.DeploymentType),
		Md5:            &config.MD5,
	}

	p, err := c.BGPConfigService.RequestBgpConfig(context.Background(), config.ProjectID).BgpConfigRequestInput(req).Execute()
	if err != nil {
		return fmt.Errorf("error enabling BGP for project %s: %w", config.ProjectID, err)
	}

	data := [][]string{{config.ProjectID, config.UseCase, strconv.Itoa(config.ASN), config.DeploymentType}}
	header := []string{"ID", "UseCase", "ASN", "DeploymentType"}
	return c.Out.Output(p, header, &data)
}
