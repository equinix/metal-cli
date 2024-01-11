package ipstest

import (
	"strings"
	"testing"

	root "github.com/equinix/metal-cli/internal/cli"
	"github.com/equinix/metal-cli/internal/ips"
	outputPkg "github.com/equinix/metal-cli/internal/outputs"
	"github.com/equinix/metal-cli/test/helper"
	"github.com/spf13/cobra"
)

func TestCli_Ips_Get(t *testing.T) {
	var ipsId string
	var err error
	subCommand := "ip"
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
				projectName := "metal-cli-" + helper.GenerateRandomString(5) + "-ips-get"
				project := helper.CreateTestProject(t, projectName)
				ipsId, err = helper.CreateTestIps(t, project.GetId(), 1, "public_ipv4")
				if len(ipsId) != 0 {
					root.SetArgs([]string{subCommand, "get", "-p", project.GetId()})

					out := helper.ExecuteAndCaptureOutput(t, root)

					if !strings.Contains(string(out[:]), ipsId) &&
						!strings.Contains(string(out[:]), "da") {
						t.Error("expected output should include " + ipsId + " da strings in the out string")
					}

					err = helper.CleanTestIps(t, ipsId)
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
