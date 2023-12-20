package interconnections

import (
	"context"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func (c *Client) Delete() *cobra.Command {
	var connectionID, confirmation string
	var force bool

	deleteConnectionCmd := &cobra.Command{
		Use:   `delete -i <connection_id> `,
		Short: "Deletes a interconnection.",
		Long:  "Deletes the specified interconnection. Use --force to skip confirmation",
		Example: `  # Deletes the specified interconnection:
  metal interconnections delete -i 7ec86e23-8dcf-48ed-bd9b-c25c20958277
  >
  ✔ Are you sure you want to delete device 7ec86e23-8dcf-48ed-bd9b-c25c20958277 [Y/n]: Y
`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true

			if !force {
				fmt.Printf("Are you sure you want to delete device %s [Y/n]: ", connectionID)

				_, err := fmt.Scanln(&confirmation)
				if err != nil {
					fmt.Println("Error reading confirmation:", err)
					return nil
				}

				confirmation = strings.TrimSpace(strings.ToLower(confirmation))
				if confirmation != "yes" && confirmation != "y" {
					fmt.Println("Interconnection deletion cancelled.")
					return nil
				}
			}

			_, _, err := c.Service.DeleteInterconnection(context.Background(), connectionID).Execute()
			if err != nil {
				return err
			}
			fmt.Println("connection deletion initiated. Please check 'metal interconnections get -i", connectionID, "' for status")

			return nil
		},
	}

	deleteConnectionCmd.Flags().StringVarP(&connectionID, "id", "i", "", "The UUID of the interconnection.")
	deleteConnectionCmd.Flags().BoolVarP(&force, "force", "f", false, "Skips confirmation for the interconnection deletion.")
	_ = deleteConnectionCmd.MarkFlagRequired("id")

	return deleteConnectionCmd
}
