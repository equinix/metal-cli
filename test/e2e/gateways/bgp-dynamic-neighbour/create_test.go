package bgp_dynamic_neighbour

import (
	"context"
	"strconv"
	"strings"
	"testing"

	metal "github.com/equinix/equinix-sdk-go/services/metalv1"
	"github.com/spf13/cobra"

	root "github.com/equinix/metal-cli/internal/cli"
	"github.com/equinix/metal-cli/internal/gateway"
	outputPkg "github.com/equinix/metal-cli/internal/outputs"
	"github.com/equinix/metal-cli/test/helper"
)

func TestBgpDynamicNeighbours_Create(t *testing.T) {
	subCommand := "gateways"
	rootClient := root.NewClient(helper.ConsumerToken, helper.URL, helper.Version)
	randomStr := helper.GenerateRandomString(5)
	projectName := "metal-cli-" + randomStr + "-gateways-create"

	project := helper.CreateTestProject(t, projectName)
	vlan := helper.CreateTestVLAN(t, project.GetId())
	vrf := helper.CreateTestVrfs(t, project.GetId(), "test-vrf-"+randomStr, vlan.GetVxlan())
	vrfIpRes := helper.CreateTestVrfIpRequest(t, project.GetId(), vrf.GetId())
	gway := helper.CreateTestVrfGateway(t, project.GetId(), vrfIpRes.VrfIpReservation.GetId(), vlan.GetId())

	tests := []struct {
		name    string
		cmd     *cobra.Command
		want    *cobra.Command
		cmdFunc func(*testing.T, *cobra.Command)
	}{
		{
			name: "create bgp dynamic neighbour",
			cmd:  gateway.NewClient(rootClient, outputPkg.Outputer(&outputPkg.Standard{})).NewCommand(),
			want: &cobra.Command{},
			cmdFunc: func(t *testing.T, c *cobra.Command) {
				root := c.Root()

				root.SetArgs([]string{subCommand, "create-bgp-dynamic-neighbours", "--gateway-id", gway.GetId(), "--bgp-neighbour-range", gway.IpReservation.GetAddress(), "--asn", "65000"})
				out := helper.ExecuteAndCaptureOutput(t, root)

				apiClient := helper.TestClient()
				neighbours, _, err := apiClient.VRFsApi.
					GetBgpDynamicNeighbors(context.Background(), gway.GetId()).
					Include([]string{"created_by"}).
					Execute()
				if err != nil {
					t.Fatal(err)
				}
				if len(neighbours.GetBgpDynamicNeighbors()) != 1 {
					t.Errorf("Bgp Dynamic Beigbours Not Found for gateway [%s]. Failed to create Bgp Dynamic Beigbour", gway.GetId())
					return
				}

				assertBgpDynamicNeighbourCmdOutput(t, string(out[:]), &neighbours.GetBgpDynamicNeighbors()[0])
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

func assertBgpDynamicNeighbourCmdOutput(t *testing.T, out string, neighbour *metal.BgpDynamicNeighbor) {
	if !strings.Contains(out, neighbour.GetId()) {
		t.Errorf("cmd output should contain ID of the BGP Dynamic Neighbour: [%s] \n output:\n%s", neighbour.GetId(), out)
	}

	if !strings.Contains(out, neighbour.GetBgpNeighborRange()) {
		t.Errorf("cmd output should contain IP Range: [%s] \n output:\n%s", neighbour.GetBgpNeighborRange(), out)
	}

	if !strings.Contains(out, strconv.Itoa(int(neighbour.GetBgpNeighborAsn()))) {
		t.Errorf("cmd output should contain asn: [%d] \n output:\n%s", neighbour.GetBgpNeighborAsn(), out)
	}

	if !(strings.Contains(out, "pending") || strings.Contains(out, "ready")) {
		t.Errorf("cmd output should contain [%s] state of the gateway, output:\n%s", string(neighbour.GetState()), out)
	}
}
