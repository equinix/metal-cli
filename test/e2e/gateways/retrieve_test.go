package gateways

import (
	"io"
	"os"
	"strconv"
	"strings"
	"testing"

	root "github.com/equinix/metal-cli/internal/cli"
	"github.com/equinix/metal-cli/internal/gateway"
	outputPkg "github.com/equinix/metal-cli/internal/outputs"
	"github.com/spf13/cobra"

	"github.com/equinix/metal-cli/test/helper"
)

func TestGateways_Retrieve(t *testing.T) {
	var projectId, deviceId string
	subCommand := "gateways"
	consumerToken := ""
	apiURL := ""
	Version := "devel"
	rootClient := root.NewClient(consumerToken, apiURL, Version)

	device := helper.SetupProjectAndDevice(t, &projectId, &deviceId)
	t.Cleanup(func() {
		if err := helper.CleanupProjectAndDevice(t, deviceId, projectId); err != nil {
			t.Error(err)
		}
	})
	if device == nil {
		return
	}

	vlan, err := helper.CreateTestVLAN(t, projectId)
	t.Cleanup(func() {
		if err := helper.CleanTestVlan(t, vlan.GetId()); err != nil {
			t.Error(err)
		}
	})
	if err != nil {
		t.Error(err)
		return
	}

	subnetSize := int32(8)
	metalGateway, err := helper.CreateTestGateway(t, projectId, vlan.GetId(), &subnetSize)
	t.Cleanup(func() {
		if err := helper.CleanTestGateway(t, metalGateway.GetId()); err != nil &&
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
			name: "retrieve gateways by projectId",
			cmd:  gateway.NewClient(rootClient, outputPkg.Outputer(&outputPkg.Standard{})).NewCommand(),
			want: &cobra.Command{},
			cmdFunc: func(t *testing.T, c *cobra.Command) {
				root := c.Root()

				// get using projectId
				root.SetArgs([]string{subCommand, "get", "-p", projectId})

				rescueStdout := os.Stdout
				r, w, _ := os.Pipe()
				os.Stdout = w
				t.Cleanup(func() {
					w.Close()
					os.Stdout = rescueStdout
				})

				if err := root.Execute(); err != nil {
					t.Fatal(err)
				}

				out, _ := io.ReadAll(r)

				assertGatewaysCmdOutput(t, string(out[:]), metalGateway.GetId(), device.Metro.GetCode(), strconv.Itoa(int(vlan.GetVxlan())))
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
