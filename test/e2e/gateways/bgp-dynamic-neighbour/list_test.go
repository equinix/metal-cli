package bgp_dynamic_neighbor

import (
	"testing"

	"github.com/spf13/cobra"

	root "github.com/equinix/metal-cli/internal/cli"
	"github.com/equinix/metal-cli/internal/gateway"
	outputPkg "github.com/equinix/metal-cli/internal/outputs"
	"github.com/equinix/metal-cli/test/helper"
)

func TestBgpDynamicNeighbors_List(t *testing.T) {
	subCommand := "gateways"
	rootClient := root.NewClient(helper.ConsumerToken, helper.URL, helper.Version)
	randomStr := helper.GenerateRandomString(5)
	projectName := "metal-cli-" + randomStr + "-gateway-get"

	project := helper.CreateTestProject(t, projectName)
	vlan := helper.CreateTestVLAN(t, project.GetId(), "da")
	vrf := helper.CreateTestVrfs(t, project.GetId(), "test-vrf-"+randomStr, 65100)
	vrfIpRes := helper.CreateTestVrfIpRequest(t, project.GetId(), vrf.GetId())
	gway := helper.CreateTestVrfGateway(t, project.GetId(), vrfIpRes.VrfIpReservation.GetId(), vlan.GetId())
	bgpDynamicNeighbor := helper.CreateTestBgpDynamicNeighbor(t, gway.GetId(), gway.IpReservation.GetAddress(), 65000)

	tests := []struct {
		name    string
		cmd     *cobra.Command
		want    *cobra.Command
		cmdFunc func(*testing.T, *cobra.Command)
	}{
		{
			name: "list bgpDynamicNeighbor by ID",
			cmd:  gateway.NewClient(rootClient, outputPkg.Outputer(&outputPkg.Standard{})).NewCommand(),
			want: &cobra.Command{},
			cmdFunc: func(t *testing.T, c *cobra.Command) {
				root := c.Root()

				// get using projectId
				root.SetArgs([]string{subCommand, "list-bgp-dynamic-neighbors", "--id", gway.GetId()})

				out := helper.ExecuteAndCaptureOutput(t, root)

				assertBgpDynamicNeighborCmdOutput(t, string(out[:]), bgpDynamicNeighbor)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rootCmd := rootClient.NewCommand()
			rootCmd.AddCommand(tt.cmd)
			tt.cmdFunc(t, tt.cmd)
		})
	}
}
