package vrf

import (
	"context"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func (c *Client) DeleteRoute() *cobra.Command {
	var (
		vrfID        string
		force        bool
		confirmation string
	)
	deleteRoute := func(id string) error {
		_, _, err := c.Service.DeleteVrfRouteById(context.Background(), id).Execute()
		if err != nil {
			return err
		}
		fmt.Println("VRF Route", id, "successfully deleted.")
		fmt.Println("VRF Route deletion initiated. Please check 'metal vrf get-route -i", vrfID, "' for status")
		return nil // No need to return 'err' here; it's always nil.
	}

	// GetVrfRouteCmd represents the GetVrfRouteCmd command
	DeleteRouteCmd := &cobra.Command{
		Use:   "delete-route [-i <VrfRoute-Id>]",
		Short: "Delete a VRF Route",
		Long:  "Delete a VRF Route",
		Example: ` #Delete a VRF Route
	metal vrf delete-route -i 77e6d57a-d7a4-4816-b451-cf9b043444e2
	>
	✔ Are you sure you want to delete device 7ec86e23-8dcf-48ed-bd9b-c25c20958277 [Y/N]: y
	
	# Deletes a VRF, skipping confirmation.
	metal vrf delete-route -f -i 77e6d57a-d7a4-4816-b451-cf9b043444e2`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true

			if !force {
				fmt.Printf("Are you sure you want to delete VRF %s [Y/N]: ", vrfID)

				_, err := fmt.Scanln(&confirmation)
				if err != nil {
					fmt.Println("Error reading confirmation:", err)
					return nil
				}

				confirmation = strings.TrimSpace(strings.ToLower(confirmation))
				if confirmation != "yes" && confirmation != "y" {
					fmt.Println("VRF deletion cancelled.")
					return nil
				}
			}

			if err := deleteRoute(vrfID); err != nil {
				return fmt.Errorf("could not delete VRF: %w", err)
			}

			return nil
		},
	}

	DeleteRouteCmd.Flags().StringVarP(&vrfID, "id", "i", "", "Specify the VRF UUID to delete the associated route configurations.")
	DeleteRouteCmd.Flags().BoolVarP(&force, "force", "f", false, "Skips confirmation for the removal of the VRF routes.")
	// making them all required here
	_ = DeleteRouteCmd.MarkFlagRequired("id")

	return DeleteRouteCmd
}
