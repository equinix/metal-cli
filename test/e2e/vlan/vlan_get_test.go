package vlan

import (
	"io"
	"os"
	"strings"
	"testing"

	root "github.com/equinix/metal-cli/internal/cli"
	outputPkg "github.com/equinix/metal-cli/internal/outputs"
	"github.com/equinix/metal-cli/internal/vlan"
	"github.com/equinix/metal-cli/test/helper"
	"github.com/spf13/cobra"
)

func TestCli_Vlan_Get(t *testing.T) {
	var projectId, vlanId string
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
				projectId, err = helper.CreateTestProject("metal-cli-vlan-delete-pro")
				if err != nil {
					t.Error(err)
				}
				vlanId, err = helper.CreateTestVlanWithVxLan(projectId, 2023, "metal-cli-vlan-delete-test")
				if len(projectId) != 0 && len(vlanId) != 0 {
					root.SetArgs([]string{subCommand, "get", "-p", projectId})
					rescueStdout := os.Stdout
					r, w, _ := os.Pipe()
					os.Stdout = w
					if err := root.Execute(); err != nil {
						t.Error(err)
					}
					w.Close()
					out, _ := io.ReadAll(r)
					os.Stdout = rescueStdout
					if !strings.Contains(string(out[:]), "metal-cli-vlan-get-test") &&
						!strings.Contains(string(out[:]), "da") &&
						!strings.Contains(string(out[:]), "2023") {
						t.Error("expected output should include metal-cli-vlan-get-test, da and 2023 strings in the out string")
					}

					err = helper.CleanTestVlan(vlanId)
					if err != nil {
						t.Error(err)
					}
					err = helper.CleanTestProject(projectId)
					if err != nil {
						t.Error(err)
					}
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
