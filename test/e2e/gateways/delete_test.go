package gateways

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/spf13/cobra"

	root "github.com/equinix/metal-cli/internal/cli"
	"github.com/equinix/metal-cli/internal/gateway"
	outputPkg "github.com/equinix/metal-cli/internal/outputs"
	"github.com/equinix/metal-cli/test/helper"
)

func TestGateways_Delete(t *testing.T) {
	subCommand := "gateways"
	rootClient := root.NewClient(helper.ConsumerToken, helper.URL, helper.Version)
	projectName := "metal-cli-" + helper.GenerateRandomString(5) + "-gateway-delete"
	project := helper.CreateTestProject(t, projectName)

	vlan := helper.CreateTestVLAN(t, project.GetId())

	subnetSize := int32(8)
	metalGateway := helper.CreateTestGateway(t, project.GetId(), vlan.GetId(), &subnetSize)

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

				out := helper.ExecuteAndCaptureOutput(t, root)

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
