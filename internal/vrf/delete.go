package vrf

import (
	"context"
	"fmt"

	"github.com/c-bata/go-prompt"
	"github.com/spf13/cobra"
)

func (c *Client) Delete() *cobra.Command {
	var (
		vrfID string
		force bool
	)

	deleteVrf := func(id string) error {
		_, err := c.Service.DeleteVrf(context.Background(), id).Execute()
		if err != nil {
			return err
		}
		fmt.Println("VRF", id, "successfully deleted.")
		fmt.Println("VRF deletion initiated. Please check 'metal vrf get -i", vrfID, "' for status")
		return nil // No need to return 'err' here; it's always nil.
	}

	deleteVrfCmd := &cobra.Command{
		Use:   "delete vrf -i <metal_vrf_UUID> [-f]",
		Short: "Deletes a VRF.",
		Long:  "Deletes the specified VRF with a confirmation prompt. To skip the confirmation, use --force.",
		Example: `# Deletes a VRF, with confirmation.
  metal delete vrf -i 77e6d57a-d7a4-4816-b451-cf9b043444e2
  >
  ✔ Are you sure you want to delete device 7ec86e23-8dcf-48ed-bd9b-c25c20958277: y

  # Deletes a VRF, skipping confirmation.
  metal delete vrf -f -i 77e6d57a-d7a4-4816-b451-cf9b043444e2`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true

			if !force {
				fmt.Printf("Are you sure you want to delete device %s (y/N): ", vrfID)
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

			if err := deleteVrf(vrfID); err != nil {
				return fmt.Errorf("could not delete VRF: %w", err)
			}

			return nil
		},
	}

	deleteVrfCmd.Flags().StringVarP(&vrfID, "id", "i", "", "UUID of the VRF.")
	deleteVrfCmd.Flags().BoolVarP(&force, "force", "f", false, "Skips confirmation for the removal of the VRF.")

	_ = deleteVrfCmd.MarkFlagRequired("id")

	return deleteVrfCmd
}
