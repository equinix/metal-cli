package interconnections

import (
	"context"
	"fmt"

	metal "github.com/equinix-labs/metal-go/metal/v1"
	"github.com/spf13/cobra"
)

func (c *Client) Create() *cobra.Command {
	var name, metro, redundancy, connType, projectID, organizationID string
	var vrfs []string

	createInterconnectionsCmd := &cobra.Command{
		Use:   `create -n <name> [-m <metro>] [-r <redundancy> ] [-t <type> ] [-p <project_id> ] | [-O <organization_id> ]`,
		Short: "Creates an interconnection.",
		Long:  "Creates a new interconnection as per the organization ID or project ID ",
		Example: `  # Creates a new interconnection named "it-interconnection":
  metal interconnections create -n <name> [-m <metro>] [-r <redundancy>] [-t "dedicated" ] [-p <project_id>] | [-O <organization_id>]

  metal interconnections create -n <name> [-m <metro>] [-r <redundancy>] [-t "shared" ] [-p <project_id>] | [-O <organization_id>] -T <service_token_type>

  metal interconnections create -n <name> [-m <metro>] [-r <redundancy>] [-t "shared" ] [-p <project_id>] | [-O <organization_id>] -T <service_token_type> -v <vrfs>`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			var interconn *metal.Interconnection
			var err error

			createOrganizationInterconnectionRequest := metal.CreateOrganizationInterconnectionRequest{DedicatedPortCreateInput: metal.NewDedicatedPortCreateInput(metro, name, redundancy, metal.DedicatedPortCreateInputType(connType))}
			if projectID != "" {

				interconn, _, err = c.Service.CreateProjectInterconnection(context.Background(), projectID).CreateOrganizationInterconnectionRequest(createOrganizationInterconnectionRequest).Execute()
				if err != nil {
					return fmt.Errorf("could not create interconnections: %w", err)

				}
			} else if organizationID != "" {
				interconn, _, err = c.Service.CreateOrganizationInterconnection(context.Background(), organizationID).CreateOrganizationInterconnectionRequest(createOrganizationInterconnectionRequest).Execute()
				if err != nil {
					return fmt.Errorf("could not create interconnections: %w", err)
				}
			} else {
				return fmt.Errorf("Could you provide at least either of projectID OR organizationID")
			}

			data := make([][]string, 1)

			data[0] = []string{interconn.GetId(), interconn.GetName(), string(interconn.GetType()), interconn.GetCreatedAt().String()}
			header := []string{"ID", "Name", "Type", "Created"}

			return c.Out.Output(interconn, header, &data)
		},
	}

	createInterconnectionsCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the interconnection")
	createInterconnectionsCmd.Flags().StringVarP(&metro, "metro", "m", "", "metro in the interconnection")
	createInterconnectionsCmd.Flags().StringVarP(&redundancy, "redundancy", "r", "", "Website URL of the organization.")
	createInterconnectionsCmd.Flags().StringVarP(&connType, "type", "t", "", "type of of interconnection.")
	// createInterconnectionsCmd.Flags().StringVarP(&connType, "serviceTokentype", "T", "", "service token type for interconnection either fabric OR Metal builds")
	createInterconnectionsCmd.Flags().StringSliceVarP(&vrfs, "vrfs", "v", []string{}, "Return only the specified vrfs.")
	createInterconnectionsCmd.Flags().StringVarP(&projectID, "projectID", "p", "", "project ID")
	createInterconnectionsCmd.Flags().StringVarP(&organizationID, "organizationID", "O", "", "Org ID")

	_ = createInterconnectionsCmd.MarkFlagRequired("name")
	_ = createInterconnectionsCmd.MarkFlagRequired("metro")
	_ = createInterconnectionsCmd.MarkFlagRequired("redundancy")
	_ = createInterconnectionsCmd.MarkFlagRequired("type")
	return createInterconnectionsCmd
}
