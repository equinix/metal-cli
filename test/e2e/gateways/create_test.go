package gateways

import (
	"context"
	"errors"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/spf13/cobra"

	root "github.com/equinix/metal-cli/internal/cli"
	"github.com/equinix/metal-cli/internal/gateway"
	outputPkg "github.com/equinix/metal-cli/internal/outputs"
	"github.com/equinix/metal-cli/test/helper"
)

func TestGateways_Create(t *testing.T) {
	var projectId, deviceId string
	subCommand := "gateways"
	consumerToken := ""
	apiURL := ""
	Version := "devel"
	rootClient := root.NewClient(consumerToken, apiURL, Version)

	device := helper.SetupProjectAndDevice(t, &projectId, &deviceId)
	t.Cleanup(func() {
		if err := helper.CleanupProjectAndDevice(deviceId, projectId); err != nil &&
			!strings.Contains(err.Error(), "Not Found") {
			t.Error(err)
		}
	})
	if device == nil {
		return
	}

	vlan, err := helper.CreateTestVLAN(projectId)
	t.Cleanup(func() {
		if err := helper.CleanTestVlan(vlan.GetId()); err != nil &&
			!strings.Contains(err.Error(), "Not Found") {
			t.Error(err)
		}
	})
	if err != nil {
		t.Error(err)
		return
	}

	tests := []struct {
		name    string
		cmd     *cobra.Command
		want    *cobra.Command
		cmdFunc func(*testing.T, *cobra.Command)
	}{
		{
			name: "create gateways",
			cmd:  gateway.NewClient(rootClient, outputPkg.Outputer(&outputPkg.Standard{})).NewCommand(),
			want: &cobra.Command{},
			cmdFunc: func(t *testing.T, c *cobra.Command) {
				root := c.Root()

				root.SetArgs([]string{subCommand, "create", "-p", projectId, "-v", vlan.GetId(), "-s", "8"})

				rescueStdout := os.Stdout
				r, w, _ := os.Pipe()
				os.Stdout = w
				if err := root.Execute(); err != nil {
					t.Error(err)
				}
				w.Close()
				out, _ := io.ReadAll(r)
				os.Stdout = rescueStdout

				apiClient := helper.TestClient()
				gateways, _, err := apiClient.MetalGatewaysApi.
					FindMetalGatewaysByProject(context.Background(), projectId).
					Execute()
				if err != nil {
					t.Error(err)
					return
				}
				if len(gateways.MetalGateways) != 1 {
					t.Error(errors.New("Gateway Not Found. Failed to create gateway"))
					return
				}

				assertGatewaysCmdOutput(t, string(out[:]), gateways.MetalGateways[0].MetalGateway.GetId(), device.Metro.GetCode(), strconv.Itoa(int(vlan.GetVxlan())))
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

func assertGatewaysCmdOutput(t *testing.T, out, gatewayId, metro, vxlan string) {
	if !strings.Contains(out, gatewayId) {
		t.Errorf("cmd output should contain ID of the gateway: [%s] \n output:\n%s", gatewayId, out)
	}

	if !strings.Contains(out, metro) {
		t.Errorf("cmd output should contain metro same as device: [%s] \n output:\n%s", metro, out)
	}

	if !strings.Contains(out, vxlan) {
		t.Errorf("cmd output should contain vxlan, gateway is attached with: [%s] \n output:\n%s", vxlan, out)
	}

	if !strings.Contains(out, "ready") {
		t.Errorf("cmd output should contain 'ready' state of the gateway, output:\n%s", out)
	}
}
