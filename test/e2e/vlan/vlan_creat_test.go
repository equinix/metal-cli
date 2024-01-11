package vlan

import (
	"strings"
	"testing"

	root "github.com/equinix/metal-cli/internal/cli"
	outputPkg "github.com/equinix/metal-cli/internal/outputs"
	"github.com/equinix/metal-cli/internal/vlan"
	"github.com/equinix/metal-cli/test/helper"
	"github.com/spf13/cobra"
)

func TestCli_Vlan_Create(t *testing.T) {
	var err error
	subCommand := "vlan"
	rootClient := root.NewClient(helper.ConsumerToken, helper.URL, helper.Version)
	type fields struct {
		MainCmd  *cobra.Command
		Outputer outputPkg.Outputer
	}
	tests := []struct {
		name    string
		fields  fields
		want    *cobra.Command
		cmdFunc func(*testing.T, *cobra.Command)
	}{
		{
			name: "create_virtual_network",
			fields: fields{
				MainCmd:  vlan.NewClient(rootClient, outputPkg.Outputer(&outputPkg.Standard{})).NewCommand(),
				Outputer: outputPkg.Outputer(&outputPkg.Standard{}),
			},
			want: &cobra.Command{},
			cmdFunc: func(t *testing.T, c *cobra.Command) {
				root := c.Root()
				projectName := "metal-cli-" + helper.GenerateRandomString(5) + "-vlan-create-pro"
				project := helper.CreateTestProject(t, projectName)
				if err != nil {
					t.Error(err)
				}
				root.SetArgs([]string{subCommand, "create", "-p", project.GetId(), "-m", "da", "--vxlan", "2023", "-d", "metal-cli-vlan-test"})

				out := helper.ExecuteAndCaptureOutput(t, root)

				if !strings.Contains(string(out[:]), "metal-cli-vlan-test") &&
					!strings.Contains(string(out[:]), "da") &&
					!strings.Contains(string(out[:]), "2023") {
					t.Error("expected output should include metal-cli-vlan-test, da and 2023 strings in the out string")
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rootCmd := rootClient.NewCommand()
			rootCmd.AddCommand(tt.fields.MainCmd)
			tt.cmdFunc(t, tt.fields.MainCmd)
		})
	}
}
