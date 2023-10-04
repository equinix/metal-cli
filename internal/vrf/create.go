package vrf

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	metal "github.com/equinix-labs/metal-go/metal/v1"
	"github.com/spf13/cobra"
)

func (c *Client) Create() *cobra.Command {
	var (
		projectID   string
		metro       string
		name        string
		description string
		ipRanges    []string
		localASN    int32
		tags        []string
	)

	// createVRFCmd represents the creatVRF command
	createVRFCmd := &cobra.Command{
		Use:     `create vrf <my_vrf> [-m <metro_code>] [-AS int] [-I str] [-d <description>]`,
		Short:   "Creates a virtual network.",
		Long:    "Creates a VRF",
		Example: `# Creates a VRF, metal vrf create `,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true

			req := metal.VrfCreateInput{
				Metro:       metro,
				Name:        name,
				IpRanges:    ipRanges,
				LocalAsn:    &localASN,
				Tags:        tags,
				Description: &description,
			}
			// Retreived these params from packetngo based on the VRF function there, assuming vrfs need these to work
			// From Packetngo >>>>
			// IPRanges is a list of all IPv4 and IPv6 Ranges that will be available to
			// BGP Peers. IPv4 addresses must be /8 or smaller with a minimum size of
			// /29. IPv6 must be /56 or smaller with a minimum size of /64. Ranges must
			// not overlap other ranges within the VRF.
			// >>>>>>>>>>>>>>>>>>>
			// Not quite sure if a CIDR can be specified here using a "/" or how it works exactly in tandem with VRFs
			// may also need some logic for processing errors on local asn, also not sure about this
			request := c.Service.CreateVrf(context.Background(), projectID).VrfCreateInput(req).Exclude(nil).Include(nil)
			n, _, err := request.Execute()
			if err != nil {
				return fmt.Errorf("Could not create VRF: %w", err)
			}

			data := make([][]string, 1)

			// This output block below is probably incorrect but leaving it for now for testing later.
			data[0] = []string{n.GetName(), *n.GetMetro().Code, n.GetDescription(), strconv.Itoa(int(n.GetLocalAsn())), strings.Join(n.GetIpRanges(), ","), n.GetCreatedAt().String()}
			header := []string{"Name", "Metro", "Description", "LocalASN", "IPrange", "Created"}

			return c.Out.Output(n, header, &data)
		},
	}

	createVRFCmd.Flags().StringVarP(&projectID, "project-id", "p", "", "The project's UUID. This flag is required, unless specified in the config created by metal init or set as METAL_PROJECT_ID environment variable.")
	createVRFCmd.Flags().StringSliceVarP(&tags, "tags", "t", []string{}, `Tag or list of tags for the VRF: --tags="tag1,tag2".`)
	createVRFCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the VRF")
	createVRFCmd.Flags().StringVarP(&description, "description", "d", "", "Description of the VRF.")

	createVRFCmd.Flags().StringVarP(&metro, "metro", "m", "", "Code of the metro where the VRF will be created")
	createVRFCmd.Flags().Int32VarP(&localASN, "local-asn", "a", 0, "Local ASN for the VRF")
	createVRFCmd.Flags().StringSliceVarP(&ipRanges, "ip-ranges", "r", []string{}, "IP range or list of IP ranges for the VRF")

	// making them all required here
	_ = createVRFCmd.MarkFlagRequired("name")
	_ = createVRFCmd.MarkFlagRequired("metro")
	_ = createVRFCmd.MarkFlagRequired("local-asn")
	_ = createVRFCmd.MarkFlagRequired("IPrange")

	return createVRFCmd
}