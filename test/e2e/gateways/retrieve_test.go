package gateways

import (
	"strconv"
	"testing"

	root "github.com/equinix/metal-cli/internal/cli"
	"github.com/equinix/metal-cli/internal/gateway"
	outputPkg "github.com/equinix/metal-cli/internal/outputs"
	"github.com/spf13/cobra"

	"github.com/equinix/metal-cli/test/helper"
)

func TestGateways_Retrieve(t *testing.T) {
	subCommand := "gateways"
	rootClient := root.NewClient(helper.ConsumerToken, helper.URL, helper.Version)
	projectName := "metal-cli-" + helper.GenerateRandomString(5) + "-gateway-get"
	project := helper.CreateTestProject(t, projectName)

	vlan := helper.CreateTestVLAN(t, project.GetId(), "sv")

	subnetSize := int32(8)
	metalGateway := helper.CreateTestGateway(t, project.GetId(), vlan.GetId(), &subnetSize)

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
				root.SetArgs([]string{subCommand, "get", "-p", project.GetId()})

				out := helper.ExecuteAndCaptureOutput(t, root)

				assertGatewaysCmdOutput(t, string(out[:]), metalGateway.GetId(), vlan.GetMetroCode(), strconv.Itoa(int(vlan.GetVxlan())))
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
