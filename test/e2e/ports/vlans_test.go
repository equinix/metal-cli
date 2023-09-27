package ports

import (
	"io"
	"os"
	"strconv"
	"testing"

	root "github.com/equinix/metal-cli/internal/cli"
	outputPkg "github.com/equinix/metal-cli/internal/outputs"
	"github.com/equinix/metal-cli/internal/ports"
	"github.com/equinix/metal-cli/test/helper"

	"github.com/spf13/cobra"
)

func TestPorts_VLANs(t *testing.T) {
	var projectId, deviceId string
	subCommand := "port"
	consumerToken := ""
	apiURL := ""
	Version := "devel"
	rootClient := root.NewClient(consumerToken, apiURL, Version)

	portList := setupProjectAndDevice(t, &projectId, &deviceId)
	port := &portList[2]
	vlan, err := helper.CreateTestVLAN(projectId)
	if err != nil {
		t.Error(err)
	}

	defer func() {
		if err := helper.UnAssignPortVlan(port.GetId(), vlan.GetId()); err != nil {
			t.Error(err)
			return
		}
		if err := helper.CleanupProjectAndDevice(deviceId, projectId); err != nil {
			t.Error(err)
		}
	}()
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
			name: "vlan assignment port",
			cmd:  ports.NewClient(rootClient, outputPkg.Outputer(&outputPkg.Standard{})).NewCommand(),
			want: &cobra.Command{},
			cmdFunc: func(t *testing.T, c *cobra.Command) {
				root := c.Root()

				vxLanStr := strconv.Itoa(int(vlan.GetVxlan()))
				// should be hybrid-bonded
				root.SetArgs([]string{subCommand, "vlan", "-i", port.GetId(), "-a", vxLanStr})

				rescueStdout := os.Stdout
				r, w, _ := os.Pipe()
				os.Stdout = w
				if err := root.Execute(); err != nil {
					t.Error(err)
				}
				w.Close()
				out, _ := io.ReadAll(r)
				os.Stdout = rescueStdout

				// wait for port to have vlans attached
				if err := helper.WaitForAttachVlanToPort(port.GetId(), true); err != nil {
					t.Error(err)
					return
				}

				assertPortCmdOutput(t, port, string(out[:]), "hybrid-bonded", true)
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
