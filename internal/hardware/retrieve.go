package hardware

import (
	"context"
	"fmt"

	metal "github.com/equinix/equinix-sdk-go/services/metalv1"
	"github.com/equinix/metal-cli/internal/outputs"
	"github.com/spf13/cobra"
)

func (c *Client) Retrieve() *cobra.Command {
	retrieveHardwareReservationsCmd := &cobra.Command{
		Use:     `get [-p <project_id>] | [-i <hardware_reservation_id>]`,
		Aliases: []string{"list"},
		Short:   "Lists a Project's hardware reservations or the details of a specified hardware reservation.",
		Long:    "Lists a Project's hardware reservations or the details of a specified hardware reservation. When using --json or --yaml flags, the --include=project,facility,device flag is implied.",
		Example: `  # Retrieve all hardware reservations of a project:
  metal hardware-reservations get -p $METAL_PROJECT_ID
  
  # Retrieve the details of a specific hardware reservation:
  metal hardware-reservations get -i 8404b73c-d18f-4190-8c49-20bb17501f88`,

		RunE: func(cmd *cobra.Command, args []string) error {
			projectID, _ := cmd.Flags().GetString("project-id")
			hardwareReservationID, _ := cmd.Flags().GetString("id")

			header := []string{"ID", "Facility", "Metro", "Plan", "Created"}

			inc := []string{}
			exc := []string{}

			// only fetch extra details when rendered
			switch c.Servicer.Format() {
			case outputs.FormatJSON, outputs.FormatYAML:
				inc = append(inc, "project", "facility", "device")
			default:
				inc = []string{"facility.metro"}
			}

			if hardwareReservationID == "" && projectID == "" {
				return fmt.Errorf("either id or project-id should be set")
			}

			cmd.SilenceUsage = true
			if hardwareReservationID != "" {
				r, _, err := c.Service.FindHardwareReservationById(context.Background(), hardwareReservationID).Include(c.Servicer.Includes(inc)).Exclude(c.Servicer.Excludes(exc)).Execute()
				if err != nil {
					return fmt.Errorf("could not get Hardware Reservation: %w", err)
				}

				data := make([][]string, 1)
				metro := ""
				if r.Facility.Metro != nil {
					metro = *r.Facility.Metro.Code
				}

				data[0] = []string{r.GetId(), r.Facility.GetCode(), metro, r.Plan.GetName(), r.CreatedAt.String()}

				return c.Out.Output(r, header, &data)
			}
			request := c.Service.FindProjectHardwareReservations(context.Background(), projectID).Include(c.Servicer.Includes(inc)).Exclude(c.Servicer.Excludes(exc))
			filters := c.Servicer.Filters()

			if filters["query"] != "" {
				request = request.Query(filters["query"])
			}

			if filters["state"] != "" {
				state, _ := metal.NewFindProjectHardwareReservationsStateParameterFromValue(filters["state"])
				request = request.State(*state)
			}

			if filters["provisionable"] != "" {
				provisionable, _ := metal.NewFindProjectHardwareReservationsProvisionableParameterFromValue(filters["provisionable"])
				request = request.Provisionable(*provisionable)
			}

			reservationsList, err := request.ExecuteWithPagination()
			if err != nil {
				return fmt.Errorf("could not list Hardware Reservations: %w", err)
			}
			reservations := reservationsList.GetHardwareReservations()
			data := make([][]string, len(reservations))

			for i, r := range reservations {
				metro := ""
				if r.Facility.Metro != nil {
					metro = r.Facility.Metro.GetCode()
				}
				data[i] = []string{r.GetId(), r.Facility.GetCode(), metro, r.Plan.GetName(), r.CreatedAt.String()}
			}

			return c.Out.Output(reservations, header, &data)
		},
	}

	retrieveHardwareReservationsCmd.Flags().StringP("project-id", "p", "", "A project's UUID.")
	retrieveHardwareReservationsCmd.Flags().StringP("id", "i", "", "The UUID of a hardware reservation.")

	return retrieveHardwareReservationsCmd
}
