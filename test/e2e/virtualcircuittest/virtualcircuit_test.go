package virtualcircuittest

import (
	"regexp"
	"strings"
	"testing"

	root "github.com/equinix/metal-cli/internal/cli"
	outputPkg "github.com/equinix/metal-cli/internal/outputs"
	"github.com/equinix/metal-cli/internal/virtualcircuit"
	"github.com/equinix/metal-cli/test/helper"
	"github.com/spf13/cobra"
)

func TestCli_Vc_Test(t *testing.T) {
	subCommand := "vc"

	rootClient := root.NewClient(helper.ConsumerToken, helper.URL, helper.Version)
	randName := helper.GenerateRandomString(5)
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
			name: "create-virtual-circuit",
			fields: fields{
				MainCmd:  virtualcircuit.NewClient(rootClient, outputPkg.Outputer(&outputPkg.Standard{})).NewCommand(),
				Outputer: outputPkg.Outputer(&outputPkg.Standard{}),
			},
			want: &cobra.Command{},
			cmdFunc: func(t *testing.T, c *cobra.Command) {
				root := c.Root()

				projName := "metal-cli-" + randName + "-vc-create-test"
				projectId := helper.CreateTestProject(t, projName)
				if projectId.GetId() != "" {
					connId := helper.CreateTestInterConnection(t, projectId.GetId(), projName)

					vlanId := helper.CreateTestVlanWithVxLan(t, projectId.GetId(), 1110, projName)
					if connId.GetId() != "" && vlanId.GetId() != "" {
						portId := helper.GetInterconnPort(t, connId.GetId())

						root.SetArgs([]string{subCommand, "create", "-P", projectId.GetId(), "-V", "1110", "--vnid", vlanId.GetId(), "-n", projName, "-s", "100", "-p", portId, "-c", connId.GetId()})

						out := helper.ExecuteAndCaptureOutput(t, root)

						if !strings.Contains(string(out[:]), projName) {
							t.Error("expected output should include " + projName + ", in the out string ")
						}
						idNamePattern := `(?m)^\| ([a-zA-Z0-9-]+) +\| *` + projName + ` *\|`

						// Find the match of the ID and NAME pattern in the table string
						match := regexp.MustCompile(idNamePattern).FindStringSubmatch(string(out[:]))

						// Extract the ID from the match
						if len(match) > 1 {
							vcId := strings.TrimSpace(match[1])
							t.Cleanup(func() {
								helper.CleanTestVirtualCircuit(t, vcId)
							})
						}
					}
				}
			},
		},
		{
			name: "delete-virtual-circuit",
			fields: fields{
				MainCmd:  virtualcircuit.NewClient(rootClient, outputPkg.Outputer(&outputPkg.Standard{})).NewCommand(),
				Outputer: outputPkg.Outputer(&outputPkg.Standard{}),
			},
			want: &cobra.Command{},
			cmdFunc: func(t *testing.T, c *cobra.Command) {
				root := c.Root()
				projName := "metal-cli-" + randName + "-vc-delete-test"
				projectId := helper.CreateTestProject(t, projName)

				if projectId.GetId() != "" {
					connId := helper.CreateTestInterConnection(t, projectId.GetId(), projName)

					vlanId := helper.CreateTestVlanWithVxLan(t, projectId.GetId(), 1110, projName)

					portId := helper.GetInterconnPort(t, connId.GetId())

					vcId := helper.CreateTestVirtualCircuit(t, projectId.GetId(), connId.GetId(), portId, vlanId.GetId(), projName)

					if connId.GetId() != "" && portId != "" && vlanId.GetId() != "" && vcId.VlanVirtualCircuit.GetId() != "" {
						root.SetArgs([]string{subCommand, "delete", "-i", vcId.VlanVirtualCircuit.GetId()})

						out := helper.ExecuteAndCaptureOutput(t, root)

						if !strings.Contains(string(out[:]), "virtual-circuit deletion initiated. Please check 'metal virtual-circuit get -i "+vcId.VlanVirtualCircuit.GetId()+" ' for status") {
							t.Error("expected output should include virtual-circuit deletion initiated. Please check '" + "metal virtual-circuit get -i " + vcId.VlanVirtualCircuit.GetId() + " ' for status in the out string ")
						}
					}
				}
			},
		},

		{
			name: "get-virtual-circuit",
			fields: fields{
				MainCmd:  virtualcircuit.NewClient(rootClient, outputPkg.Outputer(&outputPkg.Standard{})).NewCommand(),
				Outputer: outputPkg.Outputer(&outputPkg.Standard{}),
			},
			want: &cobra.Command{},
			cmdFunc: func(t *testing.T, c *cobra.Command) {
				root := c.Root()
				projName := "metal-cli-" + randName + "-vc-get-test"

				projectId := helper.CreateTestProject(t, projName)
				if projectId.GetId() != "" {
					connId := helper.CreateTestInterConnection(t, projectId.GetId(), projName)

					vlanId := helper.CreateTestVlanWithVxLan(t, projectId.GetId(), 1110, projName)

					portId := helper.GetInterconnPort(t, connId.GetId())

					if connId.GetId() != "" && portId != "" && vlanId.GetId() != "" {
						vcId := helper.CreateTestVirtualCircuit(t, projectId.GetId(), connId.GetId(), portId, vlanId.GetId(), projName)
						if vcId.VlanVirtualCircuit.GetId() != "" {
							root.SetArgs([]string{subCommand, "get", "-i", vcId.VlanVirtualCircuit.GetId()})

							out := helper.ExecuteAndCaptureOutput(t, root)

							if !strings.Contains(string(out[:]), projName) &&
								!strings.Contains(string(out[:]), vcId.VlanVirtualCircuit.GetId()) {
								t.Error("expected output should include " + projName + ", " + vcId.VlanVirtualCircuit.GetId() + "in the out string ")
							}
						}
					}
				}
			},
		},

		{
			name: "update-virtual-circuit",
			fields: fields{
				MainCmd:  virtualcircuit.NewClient(rootClient, outputPkg.Outputer(&outputPkg.Standard{})).NewCommand(),
				Outputer: outputPkg.Outputer(&outputPkg.Standard{}),
			},
			want: &cobra.Command{},
			cmdFunc: func(t *testing.T, c *cobra.Command) {
				root := c.Root()
				projName := "metal-cli-" + randName + "-vc-update-test"
				projectId := helper.CreateTestProject(t, projName)

				if projectId.GetId() != "" {
					connId := helper.CreateTestInterConnection(t, projectId.GetId(), projName)

					vlanId := helper.CreateTestVlanWithVxLan(t, projectId.GetId(), 1110, projName)

					portId := helper.GetInterconnPort(t, connId.GetId())

					vvlanId := helper.CreateTestVlanWithVxLan(t, projectId.GetId(), 1111, projName)
					vcId := helper.CreateTestVirtualCircuit(t, projectId.GetId(), connId.GetId(), portId, vlanId.GetId(), projName)

					if connId.GetId() != "" && portId != "" && vlanId.GetId() != "" && vcId.VlanVirtualCircuit.GetId() != "" {

						root.SetArgs([]string{subCommand, "update", "-i", vcId.VlanVirtualCircuit.GetId(), "-n", projName, "-v", vvlanId.GetId()})

						out := helper.ExecuteAndCaptureOutput(t, root)

						if !strings.Contains(string(out[:]), vcId.VlanVirtualCircuit.GetId()) {
							t.Error("expected output should include " + vcId.VlanVirtualCircuit.GetId() + " the out string ")
						}
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
