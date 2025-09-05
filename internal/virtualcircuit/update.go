package virtualcircuit

import (
	"context"
	"fmt"
	"strconv"

	metal "github.com/equinix/equinix-sdk-go/services/metalv1"
	"github.com/spf13/cobra"
)

type vcUpdateParamBuilder interface {
	SetName(string)
	SetDescription(string)
	SetTags([]string)
}

// compile-time assertion that parameter types implement vcParamBuilder
var (
	_ vcParamBuilder = (*metal.VlanVirtualCircuitUpdateInput)(nil)
	_ vcParamBuilder = (*metal.VrfVirtualCircuitUpdateInput)(nil)
)

func updateParams(p vcUpdateParamBuilder, name, description string, tags []string) {
	if name > "" {
		p.SetName(name)
	}
	if description > "" {
		p.SetDescription(description)
	}

	p.SetTags(tags)

}

func updateVlanVirtualCircuit(vlanUpdateInput *metal.VlanVirtualCircuitUpdateInput, vnid, speed string) metal.VirtualCircuitUpdateInput {
	vlanUpdateInput.SetVnid(vnid)

	// As per the Spec It range is [ 2 .. 4094 ]
	if speed != "" {
		vlanUpdateInput.SetSpeed(speed)
	}

	return metal.VirtualCircuitUpdateInput{VlanVirtualCircuitUpdateInput: vlanUpdateInput}
}

func updateVrfVirtualCircuit(vrfUpdateInput *metal.VrfVirtualCircuitUpdateInput, customerIP, metalIP, subnet, md5, speed string, peerAsn int) metal.VirtualCircuitUpdateInput {

	vrfUpdateInput.SetSpeed(speed)
	vrfUpdateInput.SetSubnet(subnet)
	vrfUpdateInput.SetCustomerIp(customerIP)
	vrfUpdateInput.SetMetalIp(metalIP)
	vrfUpdateInput.SetMd5(md5)
	vrfUpdateInput.SetCustomerIp(customerIP)

	if peerAsn > 0 {
		vrfUpdateInput.SetPeerAsn(int64(peerAsn))
	}

	return metal.VirtualCircuitUpdateInput{VrfVirtualCircuitUpdateInput: vrfUpdateInput}
}

func (c *Client) Update() *cobra.Command {
	var (
		description string
		speed       string
		name        string
		vnid        string
		vcID        string
		subnet      string
		customerIP  string
		metalIP     string
		md5         string
		tags        []string
		peerAsn     int
	)
	// updateConnectionCmd represents the updateConnectionCmd command
	updateVirtualCircuitCmd := &cobra.Command{
		Use:   `update -i <id> [-v <vlan UUID>] [-d <description>] [-n <name>] [-s <speed>] [-t <tags>]`,
		Short: "Updates a virtualcircuit.",
		Long:  "Updates a specified virtualcircuit etiher of vlanID OR vrfID",
		Example: `  # Updates a specified virtualcircuit etiher of vlanID OR vrfID:

	metal vc update [-i <id>] [-n <name>] [-d <description>] [-v <vnid> ] [-s <speed> ] [-t <tags> ]

	metal vc update -i e2edb90b-a8ef-47cb-a577-63b0ba129c29 -d "test-inter-fri-dedicated"

	metal vc update [-i <id>] [-n <name>] [-d <description>] [-M <md5sum>] [-a <peer-asn>] [-S <subnet>] [-c <customer-ip>] [-m <metal-ip>] [-t <tags> ]`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			inc := []string{}
			var (
				virtualCircuitUpdateInput metal.VirtualCircuitUpdateInput
				err                       error
			)
			header := []string{"ID", "Name", "Type", "Speed", "Created"}

			vnidFlag, _ := cmd.Flags().GetString("vnid")
			if vnidFlag != "" {
				vlanUpdateInput := metal.NewVlanVirtualCircuitUpdateInput()
				updateParams(vlanUpdateInput, name, description, tags)

				virtualCircuitUpdateInput = updateVlanVirtualCircuit(vlanUpdateInput, vnid, speed)
			} else {
				vrfUpdateInput := metal.NewVrfVirtualCircuitUpdateInput()

				updateParams(vrfUpdateInput, name, description, tags)

				virtualCircuitUpdateInput = updateVrfVirtualCircuit(vrfUpdateInput, customerIP, metalIP, subnet, md5, speed, peerAsn)
			}
			vc, _, err := c.Service.UpdateVirtualCircuit(context.Background(), vcID).VirtualCircuitUpdateInput(virtualCircuitUpdateInput).Include(c.Servicer.Includes(inc)).Execute()
			if err != nil {
				return fmt.Errorf("could not update virtual-circuit: %w", err)
			}

			data := make([][]string, 1)
			if vnidFlag != "" {
				data[0] = []string{vc.VlanVirtualCircuit.GetId(), vc.VlanVirtualCircuit.GetName(), string(vc.VlanVirtualCircuit.GetType()), strconv.Itoa(int(vc.VlanVirtualCircuit.GetSpeed())), vc.VlanVirtualCircuit.GetCreatedAt().String()}
			} else {
				data[0] = []string{vc.VrfVirtualCircuit.GetId(), vc.VrfVirtualCircuit.GetName(), string(vc.VrfVirtualCircuit.GetType()), strconv.Itoa(int(vc.VrfVirtualCircuit.GetSpeed())), vc.VrfVirtualCircuit.GetCreatedAt().String()}
			}

			return c.Out.Output(vc, header, &data)
		},
	}

	// Virtual Circuit Params
	updateVirtualCircuitCmd.Flags().StringVarP(&vcID, "id", "i", "", "Specify the UUID of the virtual-circuit.")

	// Common params for both VlanVirtualCircuit and VrfVirtualCircit
	updateVirtualCircuitCmd.Flags().StringVarP(&speed, "speed", "s", "", "Adds or updates Speed can be changed only if it is an interconnection on a Dedicated Port")
	updateVirtualCircuitCmd.Flags().StringSliceVarP(&tags, "tags", "t", []string{}, `updates the tags for the virtual circuit --tags "tag1,tag2" OR --tags "tag1" --tags "tag2" (NOTE: --tags "" will remove all tags from the virtual circuit`)
	updateVirtualCircuitCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the Virtual Circuit")
	updateVirtualCircuitCmd.Flags().StringVarP(&description, "description", "d", "", "Description for a Virtual Circuit")

	//VlanVirtualCircuit update input params
	updateVirtualCircuitCmd.Flags().StringVarP(&vnid, "vnid", "v", "", "A Virtual Network record UUID or the VNID of a Metro Virtual Network in your project.")

	// VrfVirtualCircuit update input params
	updateVirtualCircuitCmd.Flags().StringVarP(&md5, "md5", "M", "", "The plaintext BGP peering password shared by neighbors as an MD5 checksum")
	updateVirtualCircuitCmd.Flags().IntVarP(&peerAsn, "peer-asn", "a", 0, "The peer ASN that will be used with the VRF on the Virtual Circuit.")
	updateVirtualCircuitCmd.Flags().StringVarP(&subnet, "subnet", "S", "", "The /30 or /31 subnet of one of the VRF IP Blocks that will be used with the VRF for the Virtual Circuit. ")
	updateVirtualCircuitCmd.Flags().StringVarP(&customerIP, "customer-ip", "c", "", "An IP address from the subnet that will be used on the Customer side")
	updateVirtualCircuitCmd.Flags().StringVarP(&metalIP, "metal-ip", "m", "", "An IP address from the subnet that will be used on the Metal side. ")

	_ = updateVirtualCircuitCmd.MarkFlagRequired("id")
	return updateVirtualCircuitCmd
}
