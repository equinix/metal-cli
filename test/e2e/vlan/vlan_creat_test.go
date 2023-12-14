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

func TestCli_Vlan_Create(t *testing.T) {
	var projectId string
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
			name: "create_virtual_network",
			fields: fields{
				MainCmd:  vlan.NewClient(rootClient, outputPkg.Outputer(&outputPkg.Standard{})).NewCommand(),
				Outputer: outputPkg.Outputer(&outputPkg.Standard{}),
			},
			want: &cobra.Command{},
			cmdFunc: func(t *testing.T, c *cobra.Command) {
				root := c.Root()
				projectName := "metal-cli-vlan-create-pro" + helper.GenerateUUID()
				projectId, err = helper.CreateTestProject(t, projectName)
				if err != nil {
					t.Error(err)
				}
				if len(projectId) != 0 {
					root.SetArgs([]string{subCommand, "create", "-p", projectId, "-m", "da", "--vxlan", "2023", "-d", "metal-cli-vlan-test"})
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

					if !strings.Contains(string(out[:]), "metal-cli-vlan-test") &&
						!strings.Contains(string(out[:]), "da") &&
						!strings.Contains(string(out[:]), "2023") {
						t.Error("expected output should include metal-cli-vlan-test, da and 2023 strings in the out string")
					}

					err = helper.CleanTestProject(t, projectId)
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
