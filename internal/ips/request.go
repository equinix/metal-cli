package ips

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"

	metal "github.com/equinix/equinix-sdk-go/services/metalv1"
	"github.com/spf13/cobra"
)

func (c *Client) Request() *cobra.Command {
	var (
		ttype      string
		quantity   int
		comments   string
		facility   string
		metro      string
		projectID  string
		cidr       int
		network    string
		vrfID      string
		details    string
		tags       []string
		customdata string
	)

	// requestIPCmd represents the requestIp command
	requestIPCmd := &cobra.Command{
		Use:   `request -p <project-id> -t <ip_address_type> -q <quantity> (-m <metro> | -f <facility>) [-f <flags>] [-c <comments>]`,
		Short: "Request a block of IP addresses.",
		Long:  "Requests either a block of public IPv4 addresses or global IPv4 addresses for your project in a specific metro or facility.",
		Example: `  # Requests a block of 4 public IPv4 addresses in Dallas:
  metal ip request -p $METAL_PROJECT_ID -t public_ipv4 -q 4 -m da

  metal ip request -v df18fbd8-2919-4104-a042-5d42a05b8eed -t vrf --cidr 24 -n 172.89.1.0 --tags foo --tags bar --customdata '{"my":"goodness"}' --details "i don't think VRF users need this or will see it after submitting the request"`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			var (
				req                         *metal.IPReservationRequestInput
				vrfReq                      *metal.VrfIpReservationCreateInput
				requestIPReservationRequest *metal.RequestIPReservationRequest
			)
			cmd.SilenceUsage = true
			// It's a required flag in case of VRFIPReservations and we conduct thorough validation to ensure its inclusion.
			// By detecting its presence, we can identify whether it pertains to a standard IP Reservation Request or a VRF IP Reservation Request.

			if ttype != "vrf" {

				req = &metal.IPReservationRequestInput{
					Metro:    &metro,
					Tags:     tags,
					Quantity: int32(quantity),
					Type:     ttype,
					Facility: &facility,
				}

				requestIPReservationRequest = &metal.RequestIPReservationRequest{
					IPReservationRequestInput: req,
				}
			} else {
				// Below are required Flags in VRF IP Reservation Request.
				if cidr == 0 || network == "" || vrfID == "" {
					return errors.New(" cidr, network and ID of the VRF are required to create VFR IP Reservations")
				}
				// This is an optinal Flag in VRF IP Reservation Request.
				var data map[string]interface{}
				if customdata != "" {
					err := json.Unmarshal([]byte(customdata), &data)
					if err != nil {
						log.Fatalf("Error parsing custom data: %v", err)
					}
				}

				vrfReq = &metal.VrfIpReservationCreateInput{
					Type:       ttype,
					Cidr:       int32(cidr),
					Network:    network,
					VrfId:      vrfID,
					Details:    &details,
					Customdata: data,
					Tags:       tags,
				}
				requestIPReservationRequest = &metal.RequestIPReservationRequest{
					VrfIpReservationCreateInput: vrfReq,
				}
			}
			reservation, _, err := c.IPService.RequestIPReservation(context.Background(), projectID).RequestIPReservationRequest(*requestIPReservationRequest).Execute()
			if err != nil {
				return fmt.Errorf("could not request IP addresses: %w", err)
			}

			data := make([][]string, 1)
			if ttype != "vrf" {
				data[0] = []string{reservation.IPReservation.GetId(),
					string(reservation.IPReservation.GetType()),
					reservation.IPReservation.GetAddress(),
					strconv.FormatBool(reservation.IPReservation.GetPublic()),
					reservation.IPReservation.CreatedAt.String()}
			} else {
				data[0] = []string{reservation.VrfIpReservation.GetId(),
					string(reservation.VrfIpReservation.GetType()),
					reservation.VrfIpReservation.GetAddress(),
					strconv.FormatBool(reservation.VrfIpReservation.GetPublic()),
					reservation.VrfIpReservation.CreatedAt.String()}
			}
			header := []string{"ID", "Type", "Address", "Public", "Created"}

			return c.Out.Output(reservation, header, &data)
		},
	}

	requestIPCmd.Flags().StringVarP(&projectID, "project-id", "p", "", "The project's UUID. This flag is required, unless specified in the config created by metal init or set as METAL_PROJECT_ID environment variable.")
	requestIPCmd.Flags().StringVarP(&ttype, "type", "t", "", "The type of IP Address, either public_ipv4 or global_ipv4.")
	requestIPCmd.Flags().StringVarP(&facility, "facility", "f", "", "Code of the facility where the IP Reservation will be created")
	requestIPCmd.Flags().StringVarP(&metro, "metro", "m", "", "Code of the metro where the IP Reservation will be created")
	requestIPCmd.Flags().IntVarP(&quantity, "quantity", "q", 0, "Number of IP addresses to reserve.")
	requestIPCmd.Flags().StringSliceVarP(&tags, "tags", "", []string{}, `Adds the tags for the IP Reservations --tags "tag1,tag2" OR --tags "tag1" --tags "tag2"`)
	requestIPCmd.Flags().IntVar(&cidr, "cidr", 0, "The size of the desired subnet in bits.")
	requestIPCmd.Flags().StringVarP(&network, "network", "n", "", "The starting address for this VRF IP Reservation's subnet")
	requestIPCmd.Flags().StringVarP(&vrfID, "vrf-id", "v", "", "Specify the VRF UUID.")
	requestIPCmd.Flags().StringVarP(&details, "details", "", "", "VRF IP Reservation's details")
	requestIPCmd.Flags().StringVarP(&customdata, "customdata", "", "", "customdata is to add to the reservation, in a comma-separated list.")

	_ = requestIPCmd.MarkFlagRequired("project-id")
	_ = requestIPCmd.MarkFlagRequired("type")

	requestIPCmd.Flags().StringVarP(&comments, "comments", "c", "", "General comments or description.")
	return requestIPCmd
}
