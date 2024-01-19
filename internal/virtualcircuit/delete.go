package virtualcircuit

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

func (c *Client) Delete() *cobra.Command {
	var vcID string

	deleteVirtualcircuitCmd := &cobra.Command{
		Use:   `delete -i <virtual-circuit_id> `,
		Short: "Deletes a virtual-circuit.",
		Long:  "Deletes the specified virtual-circuit.",
		Example: `  # Deletes the specified virtual-circuit:
  metal vc delete -i 7ec86e23-8dcf-48ed-bd9b-c25c20958277`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			inc := []string{}
			exc := []string{}
			_, _, err := c.Service.DeleteVirtualCircuit(context.Background(), vcID).Include(inc).Exclude(exc).Execute()
			if err != nil {
				return fmt.Errorf("could not delete Metal Virtual circuit: %w", err)
			}
			fmt.Println("virtual-circuit deletion initiated. Please check 'metal virtual-circuit get -i", vcID, "' for status")

			return nil
		},
	}

	deleteVirtualcircuitCmd.Flags().StringVarP(&vcID, "id", "i", "", "Specify the UUID of the virtual-circuit.")
	_ = deleteVirtualcircuitCmd.MarkFlagRequired("id")

	return deleteVirtualcircuitCmd
}
