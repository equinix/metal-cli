package projects

import (
	"context"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// BGPConfigFlags holds the flags for the BGPConfig command
type BGPConfigFlags struct {
	projectID string
}

func (c *Client) BGPConfig() *cobra.Command {
	flags := BGPConfigFlags{}

	bgpConfigProjectCmd := &cobra.Command{
		Use:     `bgp-config --project-id <project_UUID>`,
		Short:   "Gets BGP Config for a project.",
		Long:    `Gets BGP Config for a project.`,
		Example: `  metal project bgp-config --project-id 50693ba9-e4e4-4d8a-9eb2-4840b11e9375 -d `,
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return getBGPConfig(c, &flags)
		},
	}

	bgpConfigProjectCmd.Flags().StringVarP(&flags.projectID, "project-id", "p", "", "Project ID (METAL_PROJECT_ID)")

	_ = bgpConfigProjectCmd.MarkFlagRequired("project-id")

	return bgpConfigProjectCmd
}

func getBGPConfig(c *Client, flags *BGPConfigFlags) error {
	p, _, err := c.BGPConfigService.FindBgpConfigByProject(context.Background(), flags.projectID).Execute()
	if err != nil {
		return fmt.Errorf("error getting BGP Config for project %s: %w", flags.projectID, err)
	}

	data := [][]string{
		{
			flags.projectID,
			string(p.GetStatus()),
			strconv.Itoa(int(p.GetAsn())),
			string(p.GetDeploymentType()),
			strconv.Itoa(int(p.GetMaxPrefix())),
		},
	}
	header := []string{"ID", "Status", "ASN", "DeploymentType", "MaxPrefix"}

	return c.Out.Output(p, header, &data)
}
