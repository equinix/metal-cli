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

func TestCli_Vlan_Create(t *testing.T) {
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
			name: "Request_NewIP",
			fields: fields{
				MainCmd:  ips.NewClient(rootClient, outputPkg.Outputer(&outputPkg.Standard{})).NewCommand(),
				Outputer: outputPkg.Outputer(&outputPkg.Standard{}),
			},
			want: &cobra.Command{},
			cmdFunc: func(t *testing.T, c *cobra.Command) {
				if true {
					t.Skip("Skipping temporarily for now")
				}
				root := c.Root()
				projectName := "metal-cli-" + helper.GenerateRandomString(5) + "-ips-request-get"
				project := helper.CreateTestProject(t, projectName)

				root.SetArgs([]string{subCommand, "request", "-p", project.GetId(), "-t", "public_ipv4", "-m", "da", "-q", "4"})

				out := helper.ExecuteAndCaptureOutput(t, root)

				if !strings.Contains(string(out[:]), "ID") &&
					!strings.Contains(string(out[:]), "PUBLIC") &&
					!strings.Contains(string(out[:]), "true") {
					t.Error("expected output should include ID, PUBLIC and true strings in the out string")
				}
			},
		},

		{
			name: "Request_NewVRFIP",
			fields: fields{
				MainCmd:  ips.NewClient(rootClient, outputPkg.Outputer(&outputPkg.Standard{})).NewCommand(),
				Outputer: outputPkg.Outputer(&outputPkg.Standard{}),
			},
			want: &cobra.Command{},
			cmdFunc: func(t *testing.T, c *cobra.Command) {
				if true {
					t.Skip("Skipping temporarily for now")
				}
				root := c.Root()
				projectName := "metal-cli-" + helper.GenerateRandomString(5) + "-ips-request-vrf"
				project := helper.CreateTestProject(t, projectName)
				_ = helper.CreateTestVlanWithVxLan(t, project.GetId(), 3988, projectName)
				vrf := helper.CreateTestVrfs(t, project.GetId(), projectName, 3988)

				root.SetArgs([]string{subCommand, "request", "-v", vrf.GetId(), "-t", "vrf", "--cidr", "24", "-n", "10.10.1.0", "--tags", "foobar", "--tags", "barfoo"})

				out := helper.ExecuteAndCaptureOutput(t, root)

				if !strings.Contains(string(out[:]), "TYPE") &&
					!strings.Contains(string(out[:]), "vrf") &&
					!strings.Contains(string(out[:]), "PUBLIC") &&
					!strings.Contains(string(out[:]), "false") {
					t.Error("expected output should include TYPE, PUBLIC and false strings in the out string")
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
