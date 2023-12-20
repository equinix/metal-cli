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

func TestCli_Vlan_Get(t *testing.T) {
	var vlanId string
	var err error
	subCommand := "vlan"
	consumerToken := ""
	apiURL := ""
	Version := "metal"
	rootClient := root.NewClient(consumerToken, apiURL, Version)
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
			name: "get_virtual_network",
			fields: fields{
				MainCmd:  vlan.NewClient(rootClient, outputPkg.Outputer(&outputPkg.Standard{})).NewCommand(),
				Outputer: outputPkg.Outputer(&outputPkg.Standard{}),
			},
			want: &cobra.Command{},
			cmdFunc: func(t *testing.T, c *cobra.Command) {
				root := c.Root()
				projectName := "metal-cli-vlan-delete-pro" + helper.GenerateRandomString(5)
				project := helper.CreateTestProject(t, projectName)
				if err != nil {
					t.Error(err)
				}
				vlanId, err = helper.CreateTestVlanWithVxLan(t, project.GetId(), 2023, "metal-cli-vlan-delete-test")
				if len(vlanId) != 0 {
					root.SetArgs([]string{subCommand, "get", "-p", project.GetId()})

					out := helper.ExecuteAndCaptureOutput(t, root)

					if !strings.Contains(string(out[:]), "metal-cli-vlan-get-test") &&
						!strings.Contains(string(out[:]), "da") &&
						!strings.Contains(string(out[:]), "2023") {
						t.Error("expected output should include metal-cli-vlan-get-test, da and 2023 strings in the out string")
					}

					helper.CleanTestVlan(t, vlanId)
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
