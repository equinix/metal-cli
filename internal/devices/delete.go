package devices

import (
	"context"
	"fmt"

	"github.com/rivo/tview"
	"github.com/spf13/cobra"
)

func (c *Client) Delete() *cobra.Command {
	var deviceID string
	var force bool

	deleteDevice := func(id string) error {
		_, err := c.Service.DeleteDevice(context.Background(), id).ForceDelete(force).Execute()
		if err != nil {
			return err
		}
		fmt.Println("Device deletion initiated. Please check 'metal device get -i", deviceID, "' for status")
		return nil
	}

	deleteDeviceCmd := &cobra.Command{
		Use:   `delete -i <device_id> [-f]`,
		Short: "Deletes a device.",
		Long:  "Deletes the specified device with a confirmation prompt. To skip the confirmation use --force.",
		Example: `  # Deletes the specified device:
  metal device delete -i 7ec86e23-8dcf-48ed-bd9b-c25c20958277
  >
  ✔ Are you sure you want to delete device 7ec86e23-8dcf-48ed-bd9b-c25c20958277: y
		
  # Deletes a VLAN, skipping confirmation:
  metal device delete -f -i 7ec86e23-8dcf-48ed-bd9b-c25c20958277`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true

			if !force {
				app := tview.NewApplication()

				modal := tview.NewModal().
					SetText(fmt.Sprintf("Are you sure you want to delete Device %s ?", deviceID)).
					AddButtons([]string{"Yes", "No"}).
					SetDoneFunc(func(buttonIndex int, buttonLabel string) {
						if buttonLabel == "Yes" {
							if err := deleteDevice(deviceID); err != nil {
								fmt.Printf("could not delete Device: %v\n", err)
							}
						}
						app.Stop()
					})

				if err := app.SetRoot(modal, false).Run(); err != nil {
					return fmt.Errorf("prompt failed: %w", err)
				}
			} else {
				if err := deleteDevice(deviceID); err != nil {
					return fmt.Errorf("could not delete Device: %w", err)
				}
			}
			return nil
		},
	}

	deleteDeviceCmd.Flags().StringVarP(&deviceID, "id", "i", "", "The UUID of the device.")
	deleteDeviceCmd.Flags().BoolVarP(&force, "force", "f", false, "Skips confirmation for the device deletion.")
	_ = deleteDeviceCmd.MarkFlagRequired("id")

	return deleteDeviceCmd
}
