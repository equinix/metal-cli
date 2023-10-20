package vlan

import (
	"regexp"
	"strings"
	"testing"

	root "github.com/equinix/metal-cli/internal/cli"
	outputPkg "github.com/equinix/metal-cli/internal/outputs"
	"github.com/equinix/metal-cli/internal/vlan"
	"github.com/equinix/metal-cli/test/helper"
	"github.com/spf13/cobra"
)

func TestCli_Vlan_Test(t *testing.T) {
	subCommand := "vlan"
	randName := helper.GenerateRandomString(5)
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
				projName := "metal-cli-" + randName + "-vlan-create-pro-test"

				projectId := helper.CreateTestProject(t, projName)

				root.SetArgs([]string{subCommand, "create", "-p", projectId.GetId(), "-m", "da", "--vxlan", "2023", "-d", projName})

				out := helper.ExecuteAndCaptureOutput(t, root)

				if !strings.Contains(string(out[:]), projName) &&
					!strings.Contains(string(out[:]), "da") &&
					!strings.Contains(string(out[:]), "2023") {
					t.Error("expected output should include metal-cli-vlan-test, da and 2023 strings in the out string")
				}
				idNamePattern := `(?m)^\| ([a-zA-Z0-9-]+) +\| *` + projName + ` *\|`

				// Find the match of the ID and NAME pattern in the table string
				match := regexp.MustCompile(idNamePattern).FindStringSubmatch(string(out[:]))

				// Extract the ID from the match
				if len(match) > 1 {
					vlanId := strings.TrimSpace(match[1])
					t.Cleanup(func() {
						helper.CleanTestVlan(t, vlanId)
					})
				}
			},
		},

		{
			name: "delete_virtual_network",
			fields: fields{
				MainCmd:  vlan.NewClient(rootClient, outputPkg.Outputer(&outputPkg.Standard{})).NewCommand(),
				Outputer: outputPkg.Outputer(&outputPkg.Standard{}),
			},
			want: &cobra.Command{},
			cmdFunc: func(t *testing.T, c *cobra.Command) {
				root := c.Root()

				projName := "metal-cli-" + randName + "-vlan-get-pro-test"

				projectId := helper.CreateTestProject(t, projName)

				vlanId := helper.CreateTestVlanWithVxLan(t, projectId.GetId(), 2023, projName)
				if vlanId.GetId() != "" {
					root.SetArgs([]string{subCommand, "delete", "-f", "-i", vlanId.GetId()})

					out := helper.ExecuteAndCaptureOutput(t, root)

					if !strings.Contains(string(out[:]), "Virtual Network "+vlanId.GetId()+" successfully deleted.") {
						t.Error("expected output should include Virtual Network " + vlanId.GetId() + "successfully deleted." + "in the out string")
					}
				}
			},
		},
		{
			name: "get_virtual_network",
			fields: fields{
				MainCmd:  vlan.NewClient(rootClient, outputPkg.Outputer(&outputPkg.Standard{})).NewCommand(),
				Outputer: outputPkg.Outputer(&outputPkg.Standard{}),
			},
			want: &cobra.Command{},
			cmdFunc: func(t *testing.T, c *cobra.Command) {
				root := c.Root()

				projName := "metal-cli-" + randName + "-vlan-delete-pro-test"

				projectId := helper.CreateTestProject(t, projName)
				vlanId := helper.CreateTestVlanWithVxLan(t, projectId.GetId(), 2023, projName)
				if vlanId.GetId() != "" {
					root.SetArgs([]string{subCommand, "get", "-p", projectId.GetId()})

					out := helper.ExecuteAndCaptureOutput(t, root)

					if !strings.Contains(string(out[:]), projName) &&
						!strings.Contains(string(out[:]), "da") &&
						!strings.Contains(string(out[:]), "2023") {
						t.Error("expected output should include " + projName + ", da and 2023 strings in the out string")
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
