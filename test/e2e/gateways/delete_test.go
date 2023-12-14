package gateways

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/spf13/cobra"

	root "github.com/equinix/metal-cli/internal/cli"
	"github.com/equinix/metal-cli/internal/gateway"
	outputPkg "github.com/equinix/metal-cli/internal/outputs"
	"github.com/equinix/metal-cli/test/helper"
)

func TestGateways_Delete(t *testing.T) {
	var projectId, deviceId string
	subCommand := "gateways"
	consumerToken := ""
	apiURL := ""
	Version := "devel"
	rootClient := root.NewClient(consumerToken, apiURL, Version)

	defer func() {
		if err := helper.CleanupProjectAndDevice(deviceId, projectId); err != nil {
			t.Error(err)
		}
	}()
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
			t.Error("Error while cleaning up vLan", err)
		}
	})
	if err != nil {
		t.Error(err)
		return
	}

	subnetSize := int32(8)
	metalGateway, err := helper.CreateTestGateway(projectId, vlan.GetId(), &subnetSize)
	t.Cleanup(func() {
		if err := helper.CleanTestGateway(metalGateway.GetId()); err != nil &&
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
			name: "delete gateways",
			cmd:  gateway.NewClient(rootClient, outputPkg.Outputer(&outputPkg.Standard{})).NewCommand(),
			want: &cobra.Command{},
			cmdFunc: func(t *testing.T, c *cobra.Command) {
				root := c.Root()

				root.SetArgs([]string{subCommand, "delete", "-f", "-i", metalGateway.GetId()})

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
					FindMetalGatewayById(context.Background(), metalGateway.GetId()).
					Include([]string{"ip_reservation"}).
					Execute()
				if err != nil && !strings.Contains(err.Error(), "Not Found") {
					t.Error(err)
				}
				if gateways != nil && gateways.MetalGateway != nil {
					t.Error(fmt.Errorf("not able to delete metal gateway with ID: [%s]", metalGateway.GetId()))
				}

				expectedOut := fmt.Sprintf("Metal Gateway %s successfully deleted.", metalGateway.GetId())
				if !strings.Contains(string(out[:]), expectedOut) {
					t.Error(fmt.Errorf("expected output: '%s' but got '%s'", expectedOut, string(out)))
				}
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
