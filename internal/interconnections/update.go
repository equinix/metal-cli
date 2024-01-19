package interconnections

import (
	"context"
	"fmt"

	metal "github.com/equinix/equinix-sdk-go/services/metalv1"
	"github.com/spf13/cobra"
)

func (c *Client) Update() *cobra.Command {
	var (
		description  string
		mode         string
		name         string
		contactEmail string
		redundancy   string
		connectionID string
		tags         []string
	)

	// updateConnectionCmd represents the updateConnectionCmd command
	updateConnectionCmd := &cobra.Command{
		Use:   `update -i <connection_id>`,
		Short: "Updates a connection.",
		Long:  "Updates a specified connection.",
		Example: `  # Updates a specified connection.:
  metal interconnections update --id 30c15082-a06e-4c43-bfc3-252616b46eba -n [<name>] -d [<description>] -r [<'redundant'|'primary'>]-m [<standard|tunnel>] -e [<E-mail>] --tags="tag1,tag2"`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			intInput := metal.NewInterconnectionUpdateInput()

			if description != "" {
				intInput.SetDescription(description)
			}

			if name != "" {
				intInput.SetName(name)
			}

			if mode != "" {
				mod := metal.InterconnectionMode(mode)
				intInput.SetMode(mod)
			}

			if contactEmail != "" {
				intInput.SetContactEmail(contactEmail)
			}

			if len(tags) > 0 {
				intInput.SetTags(tags)
			}

			interconn, _, err := c.Service.UpdateInterconnection(context.Background(), connectionID).InterconnectionUpdateInput(*intInput).Execute()
			if err != nil {
				return fmt.Errorf("could not update interconnection: %w", err)
			}

			data := make([][]string, 1)

			data[0] = []string{interconn.GetId(), interconn.GetName(), string(interconn.GetType()), interconn.GetCreatedAt().String()}
			header := []string{"ID", "Name", "Type", "Created"}

			return c.Out.Output(interconn, header, &data)
		},
	}

	updateConnectionCmd.Flags().StringVarP(&connectionID, "id", "i", "", "The UUID of the interconnection.")
	updateConnectionCmd.Flags().StringVarP(&name, "name", "n", "", "The new name of the interconnection.")
	updateConnectionCmd.Flags().StringVarP(&description, "description", "d", "", "Adds or updates the description for the interconnection.")
	updateConnectionCmd.Flags().StringVarP(&mode, "mode", "m", "", "Adds or updates the mode for the interconnection.")
	updateConnectionCmd.Flags().StringVarP(&contactEmail, "contactEmail", "e", "", "adds or updates the Email")
	updateConnectionCmd.Flags().StringSliceVarP(&tags, "tags", "t", []string{}, `Adds or updates the tags for the connection --tags="tag1,tag2".`)
	updateConnectionCmd.Flags().StringVarP(&redundancy, "redundancy", "r", "", "Updating from 'redundant' to 'primary' will remove a secondary port, while updating from 'primary' to 'redundant' will add one.")

	_ = updateConnectionCmd.MarkFlagRequired("id")
	return updateConnectionCmd
}
