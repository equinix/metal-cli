package vrf

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	metal "github.com/equinix/equinix-sdk-go/services/metalv1"
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
		Use:   "create [-p <project_id] [-d <description>] [-m <metro>] [-n <name>] [-a <localASN>] [-r <IPranges>] [-t <tags> ]",
		Short: "Creates a Virtual Routing and Forwarding(VRF) for a specified project.",
		Long:  "Creates a Creates a Virtual Routing and Forwarding(VRF) for a specified project.",
		Example: ` # Creates an Creates a Virtual Routing and Forwarding(VRF) for a specified project.
	
	metal vrf create [-p <project_id] [-d <description>] [-m <metro>] [-n <name>] [-a <localASN>] [-r <ipranges>] [-t <tags> ]`,
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

			inc := []string{}
			exc := []string{}
			vrfRequest, _, err := c.Service.CreateVrf(context.Background(), projectID).VrfCreateInput(req).Exclude(c.Servicer.Excludes(exc)).Include(c.Servicer.Includes(inc)).Execute()
			if err != nil {
				return fmt.Errorf("could not create VRF: %w", err)
			}

			data := make([][]string, 1)

			data[0] = []string{vrfRequest.GetId(), vrfRequest.GetName(), vrfRequest.GetDescription(), strings.Join(vrfRequest.GetIpRanges(), ","), strconv.Itoa(int(vrfRequest.GetLocalAsn())), vrfRequest.GetCreatedAt().String()}
			header := []string{"ID", "Name", "Description", "IPranges", "LocalASN", "Created"}

			return c.Out.Output(vrfRequest, header, &data)
		},
	}

	createVRFCmd.Flags().StringVarP(&projectID, "project-id", "p", "", "The project's UUID. This flag is required, unless specified in the config created by metal init or set as METAL_PROJECT_ID environment variable.")
	createVRFCmd.Flags().StringSliceVarP(&tags, "tags", "t", []string{}, `Adds the tags for the virtual-circuit --tags "tag1,tag2" OR --tags "tag1" --tags "tag2"`)
	createVRFCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the Virtual Routing and Forwarding")
	createVRFCmd.Flags().StringVarP(&description, "description", "d", "", "Description of the Virtual Routing and Forwarding.")

	createVRFCmd.Flags().StringVarP(&metro, "metro", "m", "", "The UUID (or metro code) for the Metro in which to create this Virtual Routing and Forwarding")
	createVRFCmd.Flags().Int32VarP(&localASN, "local-asn", "a", 0, "Local ASN for the VRF")
	createVRFCmd.Flags().StringSliceVarP(&ipRanges, "ipranges", "r", []string{}, "A list of CIDR network addresses. Like [10.0.0.0/16, 2001:d78::/56]. IPv4 blocks must be between /8 and /29 in size. IPv6 blocks must be between /56 and /64.")

	// making them all required here
	_ = createVRFCmd.MarkFlagRequired("name")
	_ = createVRFCmd.MarkFlagRequired("metro")
	_ = createVRFCmd.MarkFlagRequired("local-asn")
	_ = createVRFCmd.MarkFlagRequired("ipranges")

	return createVRFCmd
}
