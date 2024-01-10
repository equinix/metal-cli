package interconnections

import (
	"context"
	"fmt"

	metal "github.com/equinix/equinix-sdk-go/services/metalv1"
	"github.com/spf13/cobra"
)

func (c *Client) Retrieve() *cobra.Command {
	// metrosCmd represents the metros command
	var projectID, organizationID, connID string
	retrieveInterconnectionsCmd := &cobra.Command{
		Use:     "get",
		Aliases: []string{"list"},
		Short:   "Retrieves interconnections for the current user, an organization, a project or the details of a specific interconnection.",
		Long:    "Retrieves interconnections for the current user, an organization, a project or the details of a specific interconnection.",
		Example: `  # Retrieve all interconnections of a current user::
  
  # Retrieve the details of a specific interconnection:
  metal interconnections get -i e9a969b3-8911-4667-9d99-57cd3dd4ef6f

  # Retrieve all the interconnection of an organization:
  metal interconnections get -O c079178c-9557-48f2-9ce7-cfb927b81928

  # Retrieve all interconnection of a project:
  metal interconnections get -p 1867ee8f-6a11-470a-9505-952d6a324040 `,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			var interConns []metal.Interconnection
			inc := []string{}
			exc := []string{}
			header := []string{"ID", "Name", "Type", "Created"}

			if connID != "" && projectID != "" && organizationID != "" {
				return fmt.Errorf("projectID, connID, and organizationID parameters are mutually exclusive")
			} else if organizationID != "" {
				interConnectionsList, _, err := c.Service.OrganizationListInterconnections(context.Background(), organizationID).Include(inc).Exclude(exc).Execute()
				if err != nil {
					return fmt.Errorf("could not list Organization Interconnections: %w", err)
				}
				interConns = interConnectionsList.GetInterconnections()
			} else if projectID != "" {
				resp, err := c.Service.ProjectListInterconnections(context.Background(), projectID).Include(inc).Exclude(exc).ExecuteWithPagination()
				if err != nil {
					return fmt.Errorf("could not list Project Interconnections: %w", err)
				}
				interConns = resp.Interconnections
			} else if connID != "" {
				interConnection, _, err := c.Service.GetInterconnection(context.Background(), connID).Include(inc).Exclude(exc).Execute()
				if err != nil {
					return fmt.Errorf("could not list Organization Interconnections: %w", err)
				}
				data := make([][]string, 1)
				data[0] = []string{interConnection.GetId(), interConnection.GetName(), string(interConnection.GetType()), interConnection.GetCreatedAt().String()}
				return c.Out.Output(interConns, header, &data)
			}
			data := make([][]string, len(interConns))

			for i, interConn := range interConns {
				data[i] = []string{interConn.GetId(), interConn.GetName(), string(interConn.GetType()), interConn.GetCreatedAt().String()}
			}
			return c.Out.Output(interConns, header, &data)
		},
	}
	retrieveInterconnectionsCmd.Flags().StringVarP(&projectID, "projectID", "p", "", "Project ID (METAL_PROJECT_ID)")
	retrieveInterconnectionsCmd.Flags().StringVarP(&connID, "connID", "i", "", "UUID of the interconnection")
	retrieveInterconnectionsCmd.Flags().StringVarP(&organizationID, "organizationID", "O", "", "UUID of the organization")
	return retrieveInterconnectionsCmd
}
