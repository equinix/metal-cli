package ipstest

import (
	"io"
	"os"
	"strings"
	"testing"

	root "github.com/equinix/metal-cli/internal/cli"
	"github.com/equinix/metal-cli/internal/ips"
	outputPkg "github.com/equinix/metal-cli/internal/outputs"
	"github.com/equinix/metal-cli/test/helper"
	"github.com/spf13/cobra"
)

func TestCli_Ips_Get(t *testing.T) {
	var projectId, ipsId string
	var err error
	subCommand := "ip"
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
			name: "get_ip_reservations",
			fields: fields{
				MainCmd:  ips.NewClient(rootClient, outputPkg.Outputer(&outputPkg.Standard{})).NewCommand(),
				Outputer: outputPkg.Outputer(&outputPkg.Standard{}),
			},
			want: &cobra.Command{},
			cmdFunc: func(t *testing.T, c *cobra.Command) {
				if true {
					t.Skip("Skipping this test because someCondition is true")
				}
				root := c.Root()
				projectName := "metal-cli-ips-get-pro" + helper.GenerateUUID()
				projectId, err = helper.CreateTestProject(t, projectName)
				if err != nil {
					t.Error(err)
				}
				ipsId, err = helper.CreateTestIps(t, projectId, 1, "public_ipv4")
				if len(projectId) != 0 && len(ipsId) != 0 {
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
					if !strings.Contains(string(out[:]), ipsId) &&
						!strings.Contains(string(out[:]), "da") {
						t.Error("expected output should include " + ipsId + " da strings in the out string")
					}

					err = helper.CleanTestIps(t, ipsId)
					if err != nil {
						t.Error(err)
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
