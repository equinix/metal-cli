package ostest

import (
	"strings"
	"testing"

	root "github.com/equinix/metal-cli/internal/cli"
	metalos "github.com/equinix/metal-cli/internal/os"
	outputPkg "github.com/equinix/metal-cli/internal/outputs"
	"github.com/equinix/metal-cli/test/helper"
	"github.com/spf13/cobra"
)

func TestCli_OperatingSystem(t *testing.T) {
	subCommand := "operating-systems"
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
			name: "get",
			fields: fields{
				MainCmd:  metalos.NewClient(rootClient, outputPkg.Outputer(&outputPkg.Standard{})).NewCommand(),
				Outputer: outputPkg.Outputer(&outputPkg.Standard{}),
			},
			want: &cobra.Command{},
			cmdFunc: func(t *testing.T, c *cobra.Command) {
				root := c.Root()
				root.SetArgs([]string{subCommand, "get"})

				out := helper.ExecuteAndCaptureOutput(t, root)

				if !strings.Contains(string(out[:]), "RedHat Enterprise Linux 7") &&
					!strings.Contains(string(out[:]), "RancherOS") &&
					!strings.Contains(string(out[:]), "VMware ESXi 8.0") &&
					!strings.Contains(string(out[:]), "Windows 2022 Standard") &&
					!strings.Contains(string(out[:]), "Debian 10") &&
					!strings.Contains(string(out[:]), "Dell Appliance") {
					t.Error("expected output should include RedHat Enterprise Linux 7, RancherOS, VMware ESXi 8.0, Windows 2022 Standard, Debian 10, Dell Appliance")
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
