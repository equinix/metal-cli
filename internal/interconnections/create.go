package interconnections

import (
	"context"
	"errors"
	"fmt"
	"strconv"

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

			if err := validInputArgs(projectID, organizationID, connType, vlans, vrfs, svcTokenType); err != nil {
				return err
			}

			createOrganizationInterconnectionRequest := metal.CreateOrganizationInterconnectionRequest{}

			switch {
			case vlanFabricVcCreate(connType, vlans):
				in := metal.NewVlanFabricVcCreateInput(
					metro, name, redundancy, metal.VlanFabricVcCreateInputServiceTokenType(svcTokenType),
					metal.VlanFabricVcCreateInputType(connType),
				)
				in.Vlans = vlans
				// default speed
				in.SetSpeed(strconv.Itoa(int(speed)))

				createOrganizationInterconnectionRequest.
					VlanFabricVcCreateInput = in
			case vrfsFabricVcCreate(connType, vrfs):
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

	createInterconnectionsCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the interconnection.")
	createInterconnectionsCmd.Flags().StringVarP(&metro, "metro", "m", "", "Metro Id or Metro Code from where the interconnection will be originated.")
	createInterconnectionsCmd.Flags().StringVarP(&redundancy, "redundancy", "r", "", "Types of redundancy for the interconnection. Either 'primary' or 'redundant'.")
	createInterconnectionsCmd.Flags().StringVarP(&connType, "type", "t", "", "Type of of interconnection. Either 'dedicated' or 'shared' when requesting for a Fabric VC.")
	createInterconnectionsCmd.Flags().StringSliceVar(&vrfs, "vrf", []string{}, "A list of VRFs to attach to the Interconnection. Ex: --vrfs uuid1, uuid2 .")
	createInterconnectionsCmd.Flags().StringVarP(&projectID, "project-id", "p", "", "The project's UUID. Either one of this flag or --organization-id is required.")
	createInterconnectionsCmd.Flags().StringVar(&organizationID, "organization-id", "", "The Organization's UUID to be used for creating org level interconnection request. Either one of this flag or --project-id is required.")
	createInterconnectionsCmd.Flags().Int32SliceVar(&vlans, "vlan", []int32{}, "A list of VLANs to attach to the Interconnection. Ex: --vlans 1000, 1001 .")
	createInterconnectionsCmd.Flags().StringVarP(&svcTokenType, "service-token-type", "T", "", "Type of service token for shared connection. Enum: 'a_side', 'z_side'.")
	createInterconnectionsCmd.Flags().Int32VarP(&speed, "speed", "s", int32(1000000000), "The maximum speed of the interconnections.")

	_ = createInterconnectionsCmd.MarkFlagRequired("name")
	_ = createInterconnectionsCmd.MarkFlagRequired("metro")
	_ = createInterconnectionsCmd.MarkFlagRequired("redundancy")
	_ = createInterconnectionsCmd.MarkFlagRequired("type")
	return createInterconnectionsCmd
}

func vlanFabricVcCreate(connType string, vlans []int32) bool {
	return connType == "shared" && len(vlans) > 0
}

func vrfsFabricVcCreate(connType string, vrfs []string) bool {
	return connType == "shared" && len(vrfs) > 0
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

func validInputArgs(projectID, organizationID, connType string, vlans []int32, vrfs []string, svcTokenType string) error {
	if projectID == "" && organizationID == "" {
		return errors.New("could you provide at least either of projectID OR organizationID")
	}

	if (vlanFabricVcCreate(connType, vlans) || vrfsFabricVcCreate(connType, vrfs)) && svcTokenType == "" {
		return errors.New("flag 'service-token-type' is required for vlan or vrfs fabric VC create")
	}

	if vlanFabricVcCreate(connType, vlans) && vrfsFabricVcCreate(connType, vrfs) {
		return errors.New("vlans and vrfs both are provided. Please provide any one type of interconnection")
	}

	return nil
}
