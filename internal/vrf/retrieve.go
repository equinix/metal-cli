package vrf

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	metal "github.com/equinix/equinix-sdk-go/services/metalv1"
	"github.com/spf13/cobra"
)

func (c *Client) Retrieve() *cobra.Command {
	var (
		projectID string
		metro     string
		vrfID     string
	)

	// retrieveVrfsCmd represents the retrieveMetalGateways command
	retrieveVrfsCmd := &cobra.Command{
		Use:     `get -p <project_Id> `,
		Aliases: []string{"list"},
		Short:   "Lists VRFs.",
		Long:    "Retrieves a list of all VRFs for the specified project or the details of the specified VRF ID. Either a project ID or a VRF ID is required.",
		Example: ` # Gets the details of the specified device
  metal vrf get -i 3b0795ba-ec9a-4a9e-83a7-043e7e11407c

  # Lists VRFs for project 3b0795ba-ec9a-4a9e-83a7-043e7e11407c:
  metal vrf list -p 3b0795ba-ec9a-4a9e-83a7-043e7e11407c`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			inc := []string{}
			exc := []string{}
			var (
				vrfsList *metal.VrfList
				vrf      *metal.Vrf
				vrfs     []metal.Vrf
				err      error
			)
			vrfIdFlag, _ := cmd.Flags().GetString("vrfID")
			if vrfIdFlag != "" {
				vrf, _, err = c.Service.FindVrfById(context.Background(), vrfID).Include(c.Servicer.Includes(inc)).Exclude(c.Servicer.Excludes(exc)).Execute()
				if err != nil {
					return fmt.Errorf("error when calling `VRFsApi.FindVrfById``: %w", err)
				}
				vrfs[0] = *vrf
			} else {
				vrfsList, _, err = c.Service.FindVrfs(context.Background(), projectID).Metro(metro).Include(c.Servicer.Includes(inc)).Exclude(c.Servicer.Excludes(exc)).Execute()
				if err != nil {
					return fmt.Errorf("error when calling `VRFsApi.FindVrfs``: %w", err)
				}
				vrfs = vrfsList.GetVrfs()
			}

			data := make([][]string, len(vrfs))

			for i, vrf := range vrfs {

				data[i] = []string{vrf.GetId(), vrf.GetName(), vrf.GetDescription(), strings.Join(vrf.GetIpRanges(), ","), strconv.FormatInt(int64(vrf.GetLocalAsn()), 10), vrf.CreatedAt.String()}
			}
			header := []string{"ID", "Name", "Description", "IPRanges", "LocalASN", "Created"}

			return c.Out.Output(vrfs, header, &data)
		},
	}
	retrieveVrfsCmd.Flags().StringVarP(&projectID, "project-id", "p", "", "The project's UUID. This flag is required, unless specified in the config created by metal init or set as METAL_PROJECT_ID environment variable.")
	retrieveVrfsCmd.Flags().StringVarP(&metro, "metro", "m", "", "Filter by Metro ID (uuid) or Metro Code")
	retrieveVrfsCmd.Flags().StringVarP(&vrfID, "vrfID", "v", "", "VRF UUID")

	_ = retrieveVrfsCmd.MarkFlagRequired("project-id")

	return retrieveVrfsCmd
}
