package ports

import (
	"strings"
	"testing"

	root "github.com/equinix/metal-cli/internal/cli"
	outputPkg "github.com/equinix/metal-cli/internal/outputs"
	"github.com/equinix/metal-cli/internal/ports"
	"github.com/equinix/metal-cli/test/helper"

	"github.com/spf13/cobra"
)

func TestPorts_Retrieve(t *testing.T) {
	subCommand := "port"
	rootClient := root.NewClient(helper.ConsumerToken, helper.URL, helper.Version)

	_, device := helper.SetupProjectAndDevice(t, "metal-cli-port-get")

	port := &device.GetNetworkPorts()[2]
	if port == nil {
		t.Error("bond0 Port not found on device")
		return
	}

	tests := []struct {
		name    string
		cmd     *cobra.Command
		want    *cobra.Command
		cmdFunc func(*testing.T, *cobra.Command)
	}{
		{
			name: "retrieve port",
			cmd:  ports.NewClient(rootClient, outputPkg.Outputer(&outputPkg.Standard{})).NewCommand(),
			want: &cobra.Command{},
			cmdFunc: func(t *testing.T, c *cobra.Command) {
				root := c.Root()
				root.SetArgs([]string{subCommand, "get", "-i", port.GetId()})

				out := helper.ExecuteAndCaptureOutput(t, root)

				if !strings.Contains(string(out[:]), port.Data.GetMac()) {
					t.Errorf("cmd output should contain MAC address of the port: %s", port.Data.GetMac())
				}

				assertPortCmdOutput(t, port, string(out[:]), string(port.GetNetworkType()), port.Data.GetBonded())
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
