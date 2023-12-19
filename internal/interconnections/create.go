package interconnections

import (
	"context"
	"errors"
	"fmt"

	metal "github.com/equinix/equinix-sdk-go/services/metalv1"
	"github.com/spf13/cobra"
)

func (c *Client) Create() *cobra.Command {
	var name, metro, redundancy, connType, projectID, organizationID, svcTokenType string
	var vrfs []string
	var vlans []int32
	var speed int32

	createInterconnectionsCmd := &cobra.Command{
		Use:   `create -n <name> [-m <metro>] [-r <redundancy> ] [-t <type> ] [-p <project_id> ] | [-O <organization_id> ]`,
		Short: "Creates an interconnection.",
		Long:  "Creates a new interconnection as per the organization ID or project ID ",
		Example: `  # Creates a new interconnection named "it-interconnection":
  metal interconnections create -n <name> [-m <metro>] [-r <redundancy>] [-t "dedicated" ] [-p <project_id>] | [-O <organization_id>]

  metal interconnections create -n <name> [-m <metro>] [-r <redundancy>] [-t "shared" ] [-p <project_id>] | [-O <organization_id>] -T <service_token_type>

  metal interconnections create -n <name> [-m <metro>] [-r <redundancy>] [-t "shared" ] [-p <project_id>] | [-O <organization_id>] -T <service_token_type> -v <vrfs>`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true

			var interconn *metal.Interconnection
			var err error

			if err := validInputArgs(projectID, organizationID, vlans, vrfs, svcTokenType); err != nil {
				return err
			}

			createOrganizationInterconnectionRequest := metal.CreateOrganizationInterconnectionRequest{}

			switch {
			case vlanFabricVcCreate(vlans):
				in := metal.NewVlanFabricVcCreateInput(
					metro, name, redundancy, metal.VlanFabricVcCreateInputServiceTokenType(svcTokenType),
					metal.VlanFabricVcCreateInputType(connType),
				)
				in.Vlans = vlans
				// default speed
				in.SetSpeed(speed)

				createOrganizationInterconnectionRequest.
					VlanFabricVcCreateInput = in
			case vrfsFabricVcCreate(vrfs):
				createOrganizationInterconnectionRequest.
					VrfFabricVcCreateInput = metal.NewVrfFabricVcCreateInput(
					metro, name, redundancy, metal.VlanFabricVcCreateInputServiceTokenType(svcTokenType),
					metal.VlanFabricVcCreateInputType(connType), vrfs,
				)
			default:
				createOrganizationInterconnectionRequest.
					DedicatedPortCreateInput = metal.NewDedicatedPortCreateInput(
					metro, name, redundancy, metal.DedicatedPortCreateInputType(connType),
				)
			}

			interconn, err = c.handleCreate(
				projectID,
				organizationID,
				createOrganizationInterconnectionRequest)
			if err != nil {
				return fmt.Errorf("could not create interconnections: %w", err)
			}

			data := make([][]string, 1)

			data[0] = []string{interconn.GetId(), interconn.GetName(), string(interconn.GetType()), interconn.GetCreatedAt().String()}
			header := []string{"ID", "Name", "Type", "Created"}

			return c.Out.Output(interconn, header, &data)
		},
	}

	createInterconnectionsCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the interconnection")
	createInterconnectionsCmd.Flags().StringVarP(&metro, "metro", "m", "", "metro in the interconnection")
	createInterconnectionsCmd.Flags().StringVarP(&redundancy, "redundancy", "r", "", "Website URL of the organization.")
	createInterconnectionsCmd.Flags().StringVarP(&connType, "type", "t", "", "type of of interconnection.")
	createInterconnectionsCmd.Flags().StringSliceVar(&vrfs, "vrfs", []string{}, "Array of strings VRF <uuid>.")
	createInterconnectionsCmd.Flags().StringVarP(&projectID, "projectID", "p", "", "project ID")
	createInterconnectionsCmd.Flags().StringVarP(&organizationID, "organizationID", "O", "", "Org ID")
	createInterconnectionsCmd.Flags().Int32SliceVar(&vlans, "vlans", []int32{}, "Array of int vLANs")
	createInterconnectionsCmd.Flags().StringVarP(&svcTokenType, "service-token-type", "T", "", "Type of service token for shared connection. Enum: 'a_side', 'z_side'")
	createInterconnectionsCmd.Flags().Int32Var(&speed, "speed", int32(1000000000), "the maximum speed of the interconnections")

	_ = createInterconnectionsCmd.MarkFlagRequired("name")
	_ = createInterconnectionsCmd.MarkFlagRequired("metro")
	_ = createInterconnectionsCmd.MarkFlagRequired("redundancy")
	_ = createInterconnectionsCmd.MarkFlagRequired("type")
	return createInterconnectionsCmd
}

func vlanFabricVcCreate(vlans []int32) bool {
	return len(vlans) > 0
}

func vrfsFabricVcCreate(vrfs []string) bool {
	return len(vrfs) > 0
}

func (c *Client) handleCreate(projectID, organizationID string,
	req metal.CreateOrganizationInterconnectionRequest) (*metal.Interconnection, error) {

	if projectID != "" {
		interconn, _, err := c.Service.
			CreateProjectInterconnection(context.Background(), projectID).
			CreateOrganizationInterconnectionRequest(req).
			Execute()
		return interconn, err
	}

	interconn, _, err := c.Service.
		CreateOrganizationInterconnection(context.Background(), organizationID).
		CreateOrganizationInterconnectionRequest(req).
		Execute()
	return interconn, err
}

func validInputArgs(projectID, organizationID string, vlans []int32, vrfs []string, svcTokenType string) error {
	if projectID == "" && organizationID == "" {
		return errors.New("could you provide at least either of projectID OR organizationID")
	}

	if (vlanFabricVcCreate(vlans) || vrfsFabricVcCreate(vrfs)) && svcTokenType == "" {
		return errors.New("flag 'service-token-type' is required for vlan or vrfs fabric VC create")
	}

	if vlanFabricVcCreate(vlans) && vrfsFabricVcCreate(vrfs) {
		return errors.New("vLans and vrfs both are provided. Please provide any one type of interconnection")
	}

	return nil
}
