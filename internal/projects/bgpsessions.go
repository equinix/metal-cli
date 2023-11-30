package projects

import (
	"context"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

type BGPSessionsCommandArgs struct {
	ProjectID string
}

func (c *Client) BGPSessions() *cobra.Command {
	flags := BGPSessionsCommandArgs{}

	bgpSessionsProjectCmd := &cobra.Command{
		Use:     `bgp-sessions --project-id <project_UUID>`,
		Short:   "Gets BGP Sessions for a project.",
		Long:    `Gets BGP Sessions for a project.`,
		Example: `  metal project bgp-sessions --project-id 50693ba9-e4e4-4d8a-9eb2-4840b11e9375`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return retrieveBGPSessions(c, &flags)
		},
	}

	bgpSessionsProjectCmd.Flags().StringVarP(&flags.ProjectID, "project-id", "p", "", "Project ID (METAL_PROJECT_ID)")
	_ = bgpSessionsProjectCmd.MarkFlagRequired("project-id")

	return bgpSessionsProjectCmd
}

func retrieveBGPSessions(c *Client, args *BGPSessionsCommandArgs) error {
	p, _, err := c.BGPConfigService.FindProjectBgpSessions(context.Background(), args.ProjectID).Execute()
	if err != nil {
		return fmt.Errorf("error getting BGP Sessions for project %s: %w", args.ProjectID, err)
	}

	sessions := p.GetBgpSessions()
	if len(sessions) == 0 {
		fmt.Printf("No BGP Sessions found for project %s\n", args.ProjectID)
		return nil
	}

	data := make([][]string, len(sessions))
	for i, s := range sessions {
		defaultRoute := "false"
		if s.DefaultRoute != nil && *s.DefaultRoute {
			defaultRoute = "true"
		}

		data[i] = []string{
			s.GetId(),
			string(s.GetStatus()),
			strings.Join(s.LearnedRoutes, ","),
			defaultRoute,
		}
	}
	header := []string{"ID", "Status", "Learned Routes", "Default Route"}

	return c.Out.Output(p, header, &data)
}
