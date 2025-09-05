package hardware

import (
	"context"
	"fmt"

	metal "github.com/equinix/equinix-sdk-go/services/metalv1"
	"github.com/spf13/cobra"
)

func (c *Client) Move() *cobra.Command {
	var projectID, hardwareReservationID string

	moveHardwareReservationCmd := &cobra.Command{
		Use:   `move -i <hardware_reservation_id> -p <project_id>`,
		Short: "Moves a hardware reservation.",
		Long:  "Moves a hardware reservation to a specified project. Both the hardware reservation ID and the Project ID for the destination project are required.",
		Example: `  # Moves a hardware reservation to the specified Project:
  metal hardware-reservation move -i 8404b73c-d18f-4190-8c49-20bb17501f88 -p 278bca90-f6b2-4659-b1a4-1bdffa0d80b7`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			header := []string{"ID", "Facility", "Plan", "Created"}
			moveHardReserveRequest := metal.NewMoveHardwareReservationRequest()
			moveHardReserveRequest.ProjectId = &projectID
			r, _, err := c.Service.MoveHardwareReservation(context.Background(), hardwareReservationID).MoveHardwareReservationRequest(*moveHardReserveRequest).Execute()
			if err != nil {
				return fmt.Errorf("could not move Hardware Reservation: %w", err)
			}

			data := make([][]string, 1)

			data[0] = []string{r.GetId(), r.Facility.GetCode(), r.Plan.GetName(), r.CreatedAt.String()}

			return c.Out.Output(r, header, &data)
		},
	}

	moveHardwareReservationCmd.Flags().StringVarP(&hardwareReservationID, "id", "i", "", "The UUID of the hardware reservation.")
	moveHardwareReservationCmd.Flags().StringVarP(&projectID, "project-id", "p", "", "The Project ID of the Project you are moving the hardware reservation to.")
	_ = moveHardwareReservationCmd.MarkFlagRequired("project-id")
	_ = moveHardwareReservationCmd.MarkFlagRequired("id")

	return moveHardwareReservationCmd
}
