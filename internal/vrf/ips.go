package vrf

import (
	"context"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func (c *Client) Ips() *cobra.Command {
	var (
		vrfID string
		ipID  string
	)

	// createVRFCmd represents the creatVRF command
	ipsVRFCmd := &cobra.Command{
		Use:   "ips [-v <vrf-id] [-i <ip-id>]",
		Short: "Retrieves the list of VRF IP Reservations for the VRF.",
		Long:  "Retrieves the list of VRF IP Reservations for the VRF.",
		Example: ` # Retrieves the list of VRF IP Reservations for the VRF.
	
	metal vrf ips [-v <vrf-id] 

	# Retrieve a specific IP Reservation for a VRF
	metal vrf ips [-v <vrf-id] [-i <ip-id>]`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true

			inc := []string{}
			exc := []string{}

			var data [][]string
			var header = []string{"ID", "Type", "ADDRESS", "PUBLIC", "METRO", "Created"}

			if ipID == "" {
				vrfIpReservationList, _, err := c.Service.FindVrfIpReservations(context.Background(), vrfID).Include(c.Servicer.Includes(inc)).Exclude(c.Servicer.Excludes(exc)).Execute()
				if err != nil {
					return fmt.Errorf("could not find VrfIpReservations: %w", err)
				}
				ipReservations := vrfIpReservationList.GetIpAddresses()
				data = make([][]string, len(ipReservations))
				for i, vrf := range ipReservations {
					data[i] = []string{vrf.GetId(), string(vrf.GetType()), vrf.GetAddress(), strconv.FormatBool(vrf.GetPublic()), vrf.Metro.GetCode(), vrf.CreatedAt.String()}
				}
			} else {
				ipReservation, _, err := c.Service.FindVrfIpReservation(context.Background(), vrfID, ipID).Include(c.Servicer.Includes(inc)).Exclude(c.Servicer.Excludes(exc)).Execute()
				if err != nil {
					return fmt.Errorf("could not find VrfIpReservation: %w", err)
				}
				data = [][]string{{ipReservation.GetId(), string(ipReservation.GetType()), ipReservation.GetAddress(), strconv.FormatBool(ipReservation.GetPublic()), ipReservation.Metro.GetCode(), ipReservation.CreatedAt.String()}}
			}

			return c.Out.Output(nil, header, &data)
		},
	}

	ipsVRFCmd.Flags().StringVarP(&vrfID, "vrf-id", "v", "", "Specify the VRF UUID to list its associated IP reservations.")
	ipsVRFCmd.Flags().StringVarP(&ipID, "id", "i", "", "Specify the IP UUID to retrieve the details of a VRF IP reservation.")

	// making them all required here
	_ = ipsVRFCmd.MarkFlagRequired("vrf-id")

	return ipsVRFCmd
}
