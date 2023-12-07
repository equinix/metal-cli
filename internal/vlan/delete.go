package vlan

import (
	"context"
	"fmt"

	"github.com/rivo/tview"
	"github.com/spf13/cobra"
)

func (c *Client) Delete() *cobra.Command {
	var (
		vnetID string
		force  bool
	)

	deleteVnet := func(id string) error {
		_, _, err := c.Service.DeleteVirtualNetwork(context.Background(), vnetID).Execute()
		if err != nil {
			return err
		}
		fmt.Println("Virtual Network", id, "successfully deleted.")
		return err
	}

	// deleteVirtualNetworkCmd represents the deleteVirtualNetwork command
	deleteVirtualNetworkCmd := &cobra.Command{
		Use:   `delete -i <virtual_network_UUID> [-f]`,
		Short: "Deletes a virtual network.",
		Long:  "Deletes the specified VLAN with a confirmation prompt. To skip the confirmation use --force. You are not able to delete a VLAN that is attached to any ports.",
		Example: `  # Deletes a VLAN, with confirmation.
  metal virtual-network delete -i 77e6d57a-d7a4-4816-b451-cf9b043444e2
  >
  ✔ Are you sure you want to delete virtual network 77e6d57a-d7a4-4816-b451-cf9b043444e2: y

  # Deletes a VLAN, skipping confirmation.
  metal virtual-network delete -f -i 77e6d57a-d7a4-4816-b451-cf9b043444e2`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true

			if !force {
				app := tview.NewApplication()

				modal := tview.NewModal().
					SetText(fmt.Sprintf("Are you sure you want to delete virtual network %s ?", vnetID)).
					AddButtons([]string{"Yes", "No"}).
					SetDoneFunc(func(buttonIndex int, buttonLabel string) {
						if buttonLabel == "Yes" {
							if err := deleteVnet(vnetID); err != nil {
								fmt.Printf("could not delete Virtual Network: %v\n", err)
							}
						}
						app.Stop()
					})

				if err := app.SetRoot(modal, false).Run(); err != nil {
					return fmt.Errorf("prompt failed: %w", err)
				}
			} else {
				if err := deleteVnet(vnetID); err != nil {
					return fmt.Errorf("could not delete Virtual Network: %w", err)
				}
			}
			return nil
		},
	}

	deleteVirtualNetworkCmd.Flags().StringVarP(&vnetID, "id", "i", "", "UUID of the VLAN.")
	_ = deleteVirtualNetworkCmd.MarkFlagRequired("id")
	deleteVirtualNetworkCmd.Flags().BoolVarP(&force, "force", "f", false, "Skips confirmation for the removal of the virtual network.")

	return deleteVirtualNetworkCmd
}
