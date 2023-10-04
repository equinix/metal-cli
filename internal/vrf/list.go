package vrf

import (
	metal "github.com/equinix-labs/metal-go/metal/v1"
	"github.com/equinix/metal-cli/internal/outputs"
	"github.com/spf13/cobra"
	"fmt"
)

type Client struct {
	Servicer Servicer
	Service  metal.VRFsApiService
	Out      outputs.Outputer
}


func (c *Client) Retrieve() *cobra.Command {
	var projectID string

	// retrieveVrfsCmd represents the retrieveMetalGateways command
	retrieveVrfsCmd := &cobra.Command{
		Use:     `get -p <project_UUID>`,
		Aliases: []string{"list"},
		Short:   "Lists VRFs.",
		Long:    "Retrieves a list of all VRFs for the specified project.",
		Example: `
  # Lists VRFs for project 3b0795ba-ec9a-4a9e-83a7-043e7e11407c:
  metal vrf list -p 3b0795ba-ec9a-4a9e-83a7-043e7e11407c`,

  RunE: func(cmd *cobra.Command, args []string) error {
	cmd.SilenceUsage = true
	vrfs, _, err := c.Service.Get(ProjectID, c.Servicer.GetOptions(nil, nil))  // <- Not sure about the correct service to call Im assuming its this? https://github.com/packethost/packngo/blob/master/vrf.go
	if err != nil {																
		return fmt.Errorf("Could not list VRFs: %w", err)
	}

	data := make([][]string, len(vrfs.VirtualRoutingFunctions))

		for i, n := range vnets.VirtualRoutingFunctions {
			data[i] = []string{n.ID, n.Description, n.IPranges, n.LocalASN n.Metro, n.CreatedAt}  //Data should contain all info about VRF not sure where the formatting comes from, going with this.
		}
		header := []string{"ID", "Description", "VXLAN", "Facility", "Created"}

		return c.Out.Output(vrfs, header, &data)
	},
} 	
retrieveVrfsCmd.Flags().StringVarP(&projectID, "project-id", "p", "", "The project's UUID. This flag is required, unless specified in the config created by metal init or set as METAL_PROJECT_ID environment variable.")
_ = retrieveVrfsCmd.MarkFlagRequired("project-id")

return retrieveVrfsCmd
}

