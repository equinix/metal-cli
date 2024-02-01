package facilitiestest

import (
	"strings"
	"testing"

	root "github.com/equinix/metal-cli/internal/cli"
	"github.com/equinix/metal-cli/internal/facilities"
	outputPkg "github.com/equinix/metal-cli/internal/outputs"
	"github.com/spf13/cobra"

	"github.com/equinix/metal-cli/test/helper"
)

func TestFacilities_Retrieve(t *testing.T) {
	subCommand := "facilities"
	rootClient := root.NewClient(helper.ConsumerToken, helper.URL, helper.Version)
	randName := helper.GenerateRandomString(5)

	tests := []struct {
		name    string
		cmd     *cobra.Command
		want    *cobra.Command
		cmdFunc func(*testing.T, *cobra.Command)
	}{
		{
			name: "retrieve All facilities",
			cmd:  facilities.NewClient(rootClient, outputPkg.Outputer(&outputPkg.Standard{})).NewCommand(),
			want: &cobra.Command{},
			cmdFunc: func(t *testing.T, c *cobra.Command) {
				root := c.Root()

				projName := "metal-cli-" + randName + "-facilities-get-all-test"
				projectId := helper.CreateTestProject(t, projName)

				if projectId.GetId() != "" {

					root.SetArgs([]string{subCommand, "get"})

					out := helper.ExecuteAndCaptureOutput(t, root)

					if !strings.Contains(string(out[:]), "NAME") &&
						!strings.Contains(string(out[:]), "CODE") &&
						!strings.Contains(string(out[:]), "METRO") &&
						!strings.Contains(string(out[:]), "baremetal,backend_transfer,layer_2,global_ipv4,ibx") &&
						!strings.Contains(string(out[:]), "Singapore") {
						t.Error("expected output should include NAME CODE METRO Singapore and baremetal,backend_transfer,layer_2,global_ipv4,ibx, in the out string ")
					}
				}
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
