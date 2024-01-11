package ports

import (
	"strconv"
	"strings"
	"testing"

	root "github.com/equinix/metal-cli/internal/cli"
	outputPkg "github.com/equinix/metal-cli/internal/outputs"
	"github.com/equinix/metal-cli/internal/ports"
	"github.com/equinix/metal-cli/test/helper"

	metal "github.com/equinix/equinix-sdk-go/services/metalv1"
	"github.com/spf13/cobra"
)

func TestPorts_Convert(t *testing.T) {
	subCommand := "port"
	rootClient := root.NewClient(helper.ConsumerToken, helper.URL, helper.Version)

	_, device := helper.SetupProjectAndDevice(t, "metal-cli-port-convert")

	port := &device.GetNetworkPorts()[2]
	if port == nil {
		t.Error("bond0 Port not found on device")
		return
	}

	tests := []struct {
		name                string
		cmd                 *cobra.Command
		want                *cobra.Command
		cmdFunc             func(*testing.T, *cobra.Command)
		expectedNetworkType string
		expectedBonded      bool
	}{
		{
			name: "convert port layer-2 bonded false",
			cmd:  ports.NewClient(rootClient, outputPkg.Outputer(&outputPkg.Standard{})).NewCommand(),
			want: &cobra.Command{},
			cmdFunc: func(t *testing.T, c *cobra.Command) {
				root := c.Root()

				root.SetArgs([]string{subCommand, "convert", "-i", port.GetId(), "--layer2", "--bonded=false", "--force"})

				out := helper.ExecuteAndCaptureOutput(t, root)

				assertPortCmdOutput(t, port, string(out[:]), "layer2-individual", false)
			},
		},
		{
			name: "convert port layer-2 bonded true",
			cmd:  ports.NewClient(rootClient, outputPkg.Outputer(&outputPkg.Standard{})).NewCommand(),
			want: &cobra.Command{},
			cmdFunc: func(t *testing.T, c *cobra.Command) {
				root := c.Root()

				root.SetArgs([]string{subCommand, "convert", "-i", port.GetId(), "--layer2", "--bonded", "--force"})

				out := helper.ExecuteAndCaptureOutput(t, root)

				assertPortCmdOutput(t, port, string(out[:]), "layer2-bonded", true)
			},
		},
		{
			name: "convert port layer-3 bonded true",
			cmd:  ports.NewClient(rootClient, outputPkg.Outputer(&outputPkg.Standard{})).NewCommand(),
			want: &cobra.Command{},
			cmdFunc: func(t *testing.T, c *cobra.Command) {
				root := c.Root()

				root.SetArgs([]string{subCommand, "convert", "-i", port.GetId(), "-2=false", "--force"})

				out := helper.ExecuteAndCaptureOutput(t, root)

				assertPortCmdOutput(t, port, string(out[:]), "layer3", true)
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

func assertPortCmdOutput(t *testing.T, port *metal.Port, out, networkType string, bonded bool) {
	if !strings.Contains(out, port.GetId()) {
		t.Errorf("cmd output should contain ID of the port: %s", port.GetId())
	}

	if !strings.Contains(out, port.GetName()) {
		t.Errorf("cmd output should contain name of the port: %s", port.GetName())
	}

	if !strings.Contains(out, networkType) {
		t.Errorf("cmd output should contain type of the port: %s", string(port.GetNetworkType()))
	}

	if !strings.Contains(out, strconv.FormatBool(bonded)) {
		t.Errorf("cmd output should contain if port is bonded: %s", strconv.FormatBool(port.Data.GetBonded()))
	}
}
