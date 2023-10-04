package main // or your appropriate package name

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

func (c *Client) Delete() *cobra.Command {
	var (
		vrfID string
		force  bool
	)

	deleteVrf := func(id string) error {
		_, err := c.Service.Delete(id)
		if err != nil {
			return err
		}
		fmt.Println("VRF", id, "successfully deleted.")
		return nil // No need to return 'err' here; it's always nil.
	}

	deleteMetalVrfCmd := &cobra.Command{
		Use:   "delete vrf -i <metal_vrf_UUID> [-f]",
		Short: "Deletes a VRF.",
		Long:  "Deletes the specified VRF with a confirmation prompt. To skip the confirmation, use --force.",
		Example: `# Deletes a VRF, with confirmation.
  metal delete vrf -i 77e6d57a-d7a4-4816-b451-cf9b043444e2

  # Deletes a VRF, skipping confirmation.
  metal delete vrf -f -i 77e6d57a-d7a4-4816-b451-cf9b043444e2`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true

			if !force {
				prompt := promptui.Prompt{
					Label:     fmt.Sprintf("Are you sure you want to delete VRF %s: ", vrfID),
					IsConfirm: true,
				}

				result, err := prompt.Run()
				if err != nil || result != "y" {
					fmt.Println("VRF deletion aborted.")
					return nil
				}
			}

			if err := deleteVrf(vrfID); err != nil {
				return fmt.Errorf("Could not delete VRF: %w", err)
			}

			return nil
		},
	}

	deleteMetalVrfCmd.Flags().StringVarP(&vrfID, "id", "i", "", "UUID of the VRF.")
	_ = deleteMetalVrfCmd.MarkFlagRequired("id")
	deleteMetalVrfCmd.Flags().BoolVarP(&force, "force", "f", false, "Skips confirmation for the removal of the VRF.")

	return deleteMetalVrfCmd
}
