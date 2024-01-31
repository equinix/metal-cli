package ports

import (
	"context"
	"strconv"
	"testing"

	root "github.com/equinix/metal-cli/internal/cli"
	outputPkg "github.com/equinix/metal-cli/internal/outputs"
	"github.com/equinix/metal-cli/internal/ports"
	"github.com/equinix/metal-cli/test/helper"

	metal "github.com/equinix/equinix-sdk-go/services/metalv1"
	"github.com/spf13/cobra"
)

func TestPorts_VLANs(t *testing.T) {
	subCommand := "port"
	rootClient := root.NewClient(helper.ConsumerToken, helper.URL, helper.Version)

	project, device := helper.SetupProjectAndDevice(t, "metal-cli-port-vlan")

	port := &device.GetNetworkPorts()[2]
	if port == nil {
		t.Error("bond0 Port not found on device")
		return
	}

	if err := convertToLayer2(port.GetId()); err != nil {
		t.Error(err)
		return
	}

	vlan := helper.CreateTestVLAN(t, project.GetId(), "sv")
	t.Cleanup(func() {
		if err := helper.UnAssignPortVlan(t, port.GetId(), vlan.GetId()); err != nil {
			t.Error(err)
		}
	})

	tests := []struct {
		name    string
		cmd     *cobra.Command
		want    *cobra.Command
		cmdFunc func(*testing.T, *cobra.Command)
	}{
		{
			name: "vlan assignment port",
			cmd:  ports.NewClient(rootClient, outputPkg.Outputer(&outputPkg.Standard{})).NewCommand(),
			want: &cobra.Command{},
			cmdFunc: func(t *testing.T, c *cobra.Command) {
				root := c.Root()

				vxLanStr := strconv.Itoa(int(vlan.GetVxlan()))
				// should be layer2-bonded
				root.SetArgs([]string{subCommand, "vlan", "-i", port.GetId(), "-a", vxLanStr})

				out := helper.ExecuteAndCaptureOutput(t, root)

				// wait for port to have vlans attached
				if err := helper.WaitForAttachVlanToPort(t, port.GetId(), true); err != nil {
					t.Fatal(err)
				}

				helper.AssertPortCmdOutput(t, port, string(out[:]), "layer2-bonded", true)
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

func convertToLayer2(portId string) error {
	apiClient := helper.TestClient()

	_, _, err := apiClient.PortsApi.ConvertLayer2(context.Background(), portId).
		PortAssignInput(*metal.NewPortAssignInput()).
		Execute()
	return err
}
