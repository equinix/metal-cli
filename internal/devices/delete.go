package devices

import (
	"context"
	"fmt"

	"github.com/c-bata/go-prompt"
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
				fmt.Printf("Are you sure you want to delete device %s (y/N): ", deviceID)
				userInput := prompt.Input(">", func(d prompt.Document) []prompt.Suggest {
					return []prompt.Suggest{
						{Text: "y", Description: "Yes"},
						{Text: "n", Description: "No"},
					}
				})
				if userInput != "y" && userInput != "Y" {
					return nil
				}
			}
			if err := deleteDevice(deviceID); err != nil {
				return fmt.Errorf("could not delete Device: %w", err)
			}
			return nil
		},
	}

	deleteDeviceCmd.Flags().StringVarP(&deviceID, "id", "i", "", "The UUID of the device.")
	deleteDeviceCmd.Flags().BoolVarP(&force, "force", "f", false, "Skips confirmation for the device deletion.")
	_ = deleteDeviceCmd.MarkFlagRequired("id")

	return deleteDeviceCmd
}
