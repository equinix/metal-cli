package virtualcircuit

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	metal "github.com/equinix/equinix-sdk-go/services/metalv1"
	"github.com/spf13/cobra"
)

type vcParamBuilder interface {
	SetName(string)
	SetDescription(string)
	SetTags([]string)
}

// compile-time assertion that parameter types implement vcParamBuilder
var (
	_ vcParamBuilder = (*metal.VlanVirtualCircuitCreateInput)(nil)
	_ vcParamBuilder = (*metal.VrfVirtualCircuitCreateInput)(nil)
)

func setParams(p vcParamBuilder, name, description string, tags []string) {
	if name != "" {
		p.SetName(name)
	}
	if description != "" {
		p.SetDescription(description)
	}

	p.SetTags(tags)

}

func validateParams(nnVlanFlag int, asnFlag int, subnetFlag string) error {
	if nnVlanFlag == 0 || asnFlag == 0 || subnetFlag == "" {
		return errors.New(" vlan ID, peer ASN and subnet of one of the VRF IP Blocks is required to create VRF virtual circuit")
	}
	return nil
}

func createVrfVirtualCircuit(vrfInput *metal.VrfVirtualCircuitCreateInput, projectID, customerIP, metalIP, md5 string, speed int) metal.VirtualCircuitCreateInput {

	vrfInput.SetCustomerIp(customerIP)
	vrfInput.SetMetalIp(metalIP)
	vrfInput.SetMd5(md5)
	if speed > 0 {
		vrfInput.SetSpeed(strconv.Itoa(speed))
	}

	return metal.VirtualCircuitCreateInput{VrfVirtualCircuitCreateInput: vrfInput}
}

func createVlanVirtualCircuit(vlanInput *metal.VlanVirtualCircuitCreateInput, vnid string, nnVlan int, speed int) metal.VirtualCircuitCreateInput {
	vlanInput.SetVnid(vnid)
	// As per the Spec It range is [ 2 .. 4094 ]
	if nnVlan > 2 {
		vlanInput.SetNniVlan(int32(nnVlan))
	}
	if speed > 0 {
		vlanInput.SetSpeed(strconv.Itoa(speed))
	}
	vlanInput.SetVnid(vnid)

	return metal.VirtualCircuitCreateInput{VlanVirtualCircuitCreateInput: vlanInput}
}

func (c *Client) Create() *cobra.Command {
	var (
		connectionID string
		portID       string
		name         string
		description  string
		projectID    string
		vnid         string
		vrf          string
		subnet       string
		customerIP   string
		metalIP      string
		md5          string
		peerAsn      int
		speed        int
		nnVlan       int
		tags         []string
	)

	createVirtualCircuitCmd := &cobra.Command{
		Use:   `create  [-c connection_id] [-p port_id] [-P <project_id> ] -n <name> [-d <description>] [--vnid <vnid> ] [-V <vlan> ] [-s <speed> ] [-t <tags> ]`,
		Short: "Creates an create-virtual-circuit for specific interconnection.",
		Long:  "Creates an create-virtual-circuit for specific interconnection",
		Example: `  # Creates a new virtual-circuit named "interconnection": 
  metal vc create [-c connection_id] [-p port_id] [-P <project_id> ] [-n <name>] [-d <description>] [--vnid <vnid> ] [-V <vlan> ] [-s <speed> ] [-t <tags> ]

  metal vc create -c 81c9cb9e-b02f-4c73-9e04-06702f1380a0 -p 9c8f0c71-591d-42fe-9519-2f632761e2da -P b4673e33-0f48-4948-961a-c31d6edf64f8 -n test-inter  -d test-interconnection -v 15315810-2fda-48b8-b8cd-441ebab684b5 -V 1010 -s 100
  
  metal vc create [-c connection_id] [-p port_id] [-P <project_id> ] [-n <name>] [-d <description>] [-v <vrf-id>] [-M <md5sum>] [-a <peer-asn>] [-S <subnet>] [-c <customer_ip>] [-m <metal_ip>]`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			var (
				createInput metal.VirtualCircuitCreateInput
				err         error
			)
			vrfFlag, _ := cmd.Flags().GetString("vrf")
			if vrfFlag != "" {
				nnVlanFlag, _ := cmd.Flags().GetInt("nnVlan")
				asnFlag, _ := cmd.Flags().GetInt("peerAsn")
				subnetFlag, _ := cmd.Flags().GetString("subnet")

				err = validateParams(nnVlanFlag, asnFlag, subnetFlag)
				if err != nil {
					return fmt.Errorf("%w", err)
				}
				vrfInput := metal.NewVrfVirtualCircuitCreateInput(int32(nnVlan), int64(peerAsn), projectID, subnet, vrf)
				setParams(vrfInput, name, description, tags)
				createInput = createVrfVirtualCircuit(vrfInput, projectID, customerIP, metalIP, md5, speed)
			} else {
				vlanInput := metal.NewVlanVirtualCircuitCreateInput(projectID)

				setParams(vlanInput, name, description, tags)
				createInput = createVlanVirtualCircuit(vlanInput, vnid, nnVlan, speed)
			}

			vc, _, err := c.Service.CreateInterconnectionPortVirtualCircuit(context.Background(), connectionID, portID).VirtualCircuitCreateInput(createInput).Execute()
			if err != nil {
				return fmt.Errorf("could not create Virtual Circuit: %w", err)
			}

			data := make([][]string, 1)

			if vrfFlag != "" {
				data[0] = []string{vc.VrfVirtualCircuit.GetId(), vc.VrfVirtualCircuit.GetName(), string(vc.VrfVirtualCircuit.GetType()), strconv.Itoa(int(vc.VrfVirtualCircuit.GetSpeed())), vc.VrfVirtualCircuit.GetCreatedAt().String()}
			} else {
				data[0] = []string{vc.VlanVirtualCircuit.GetId(), vc.VlanVirtualCircuit.GetName(), string(vc.VlanVirtualCircuit.GetType()), strconv.Itoa(int(vc.VlanVirtualCircuit.GetSpeed())), vc.VlanVirtualCircuit.GetCreatedAt().String()}
			}

			header := []string{"ID", "Name", "Type", "Speed", "Created"}

			return c.Out.Output(vc.VlanVirtualCircuit, header, &data)
		},
	}
	// Virtual Circuit Params
	createVirtualCircuitCmd.Flags().StringVarP(&connectionID, "connection-id", "c", "", "Specify the UUID of the interconnection.")
	createVirtualCircuitCmd.Flags().StringVarP(&portID, "port-id", "p", "", "Specify the UUID of the port.")

	// Common params for both VlanVirtualCircuit and VrfVirtualCircit
	createVirtualCircuitCmd.Flags().IntVarP(&nnVlan, "vlan", "V", 0, "Adds or updates vlan  Must be between 2 and 4094")
	createVirtualCircuitCmd.Flags().IntVarP(&speed, "speed", "s", 0, "bps speed or string (e.g. 52 - '52m' or '100g' or '4 gbps')")
	createVirtualCircuitCmd.Flags().StringSliceVarP(&tags, "tags", "t", []string{}, `Adds the tags for the virtual-circuit --tags "tag1,tag2" OR --tags "tag1" --tags "tag2"`)
	createVirtualCircuitCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the Virtual Circuit")
	createVirtualCircuitCmd.Flags().StringVarP(&description, "description", "d", "", "Description for a Virtual Circuit")
	createVirtualCircuitCmd.Flags().StringVarP(&projectID, "project-id", "P", "", "The project's UUID. This flag is required, unless specified in the config created by metal init or set as METAL_PROJECT_ID environment variable.")

	//VlanVirtualCircuit input params
	createVirtualCircuitCmd.Flags().StringVarP(&vnid, "vnid", "", "", "Specify the UUID  of the VLAN.")

	// VrfVirtualCircuit input params
	createVirtualCircuitCmd.Flags().StringVarP(&md5, "md5", "M", "", "The plaintext BGP peering password shared by neighbors as an MD5 checksum")
	createVirtualCircuitCmd.Flags().IntVarP(&peerAsn, "peer-asn", "a", 0, "The peer ASN that will be used with the VRF on the Virtual Circuit.")
	createVirtualCircuitCmd.Flags().StringVarP(&vrf, "vrf-id", "v", "", "The UUID of the VRF that will be associated with the Virtual Circuit.")
	createVirtualCircuitCmd.Flags().StringVarP(&subnet, "subnet", "S", "", "The /30 or /31 subnet of one of the VRF IP Blocks that will be used with the VRF for the Virtual Circuit. ")
	createVirtualCircuitCmd.Flags().StringVarP(&customerIP, "customer-ip", "", "", "An IP address from the subnet that will be used on the Customer side")
	createVirtualCircuitCmd.Flags().StringVarP(&metalIP, "metal-ip", "m", "", "An IP address from the subnet that will be used on the Metal side. ")

	_ = createVirtualCircuitCmd.MarkFlagRequired("connection-id")
	_ = createVirtualCircuitCmd.MarkFlagRequired("port-id")
	_ = createVirtualCircuitCmd.MarkFlagRequired("project-id")
	return createVirtualCircuitCmd
}
