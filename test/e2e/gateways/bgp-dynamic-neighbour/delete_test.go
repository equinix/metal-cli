package bgp_dynamic_neighbour

import (
	"context"
	"strings"
	"testing"

	"github.com/spf13/cobra"

	root "github.com/equinix/metal-cli/internal/cli"
	"github.com/equinix/metal-cli/internal/gateway"
	outputPkg "github.com/equinix/metal-cli/internal/outputs"
	"github.com/equinix/metal-cli/test/helper"
)

func TestBgpDynamicNeighbours_Delete(t *testing.T) {
	subCommand := "gateways"
	rootClient := root.NewClient(helper.ConsumerToken, helper.URL, helper.Version)
	randomStr := helper.GenerateRandomString(5)
	projectName := "metal-cli-" + randomStr + "-gateway-delete"

	project := helper.CreateTestProject(t, projectName)
	vlan := helper.CreateTestVLAN(t, project.GetId())
	vrf := helper.CreateTestVrfs(t, project.GetId(), "test-vrf-"+randomStr, vlan.GetVxlan())
	vrfIpRes := helper.CreateTestVrfIpRequest(t, project.GetId(), vrf.GetId())
	gway := helper.CreateTestVrfGateway(t, project.GetId(), vrfIpRes.VrfIpReservation.GetId(), vlan.GetId())
	bgpDynamicNeighbour := helper.CreateTestBgpDynamicNeighbour(t, gway.GetId(), gway.IpReservation.GetAddress(), 65000)

	tests := []struct {
		name    string
		cmd     *cobra.Command
		want    *cobra.Command
		cmdFunc func(*testing.T, *cobra.Command)
	}{
		{
			name: "delete gateways",
			cmd:  gateway.NewClient(rootClient, outputPkg.Outputer(&outputPkg.Standard{})).NewCommand(),
			want: &cobra.Command{},
			cmdFunc: func(t *testing.T, c *cobra.Command) {
				root := c.Root()

				root.SetArgs([]string{subCommand, "delete-bgp-dynamic-neighbours", "-i", bgpDynamicNeighbour.GetId()})

				out := helper.ExecuteAndCaptureOutput(t, root)

				apiClient := helper.TestClient()
				_, _, err := apiClient.VRFsApi.
					BgpDynamicNeighborsIdGet(context.Background(), bgpDynamicNeighbour.GetId()).
					Include([]string{"created_by"}).
					Execute()
				if err != nil && !strings.Contains(err.Error(), "Not Found") {
					t.Fatal(err)
				}

				strings.Contains(string(out), "BGP Dynamic Neighbour deletion initiated. Please check 'metal gateway get-bgp-dynamic-neighbour -i "+bgpDynamicNeighbour.GetId()+"' for status")
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
