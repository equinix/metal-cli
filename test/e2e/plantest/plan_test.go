package plantest

import (
	"os"
	"strings"
	"testing"

	root "github.com/equinix/metal-cli/internal/cli"
	outputPkg "github.com/equinix/metal-cli/internal/outputs"
	"github.com/equinix/metal-cli/internal/plans"
	"github.com/equinix/metal-cli/test/helper"
	"github.com/spf13/cobra"
)

func TestCli_Plans(t *testing.T) {
	subCommand := "plans"
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
			name: "get",
			fields: fields{
				MainCmd:  plans.NewClient(rootClient, outputPkg.Outputer(&outputPkg.Standard{})).NewCommand(),
				Outputer: outputPkg.Outputer(&outputPkg.Standard{}),
			},
			want: &cobra.Command{},
			cmdFunc: func(t *testing.T, c *cobra.Command) {
				root := c.Root()
				root.SetArgs([]string{subCommand, "get"})

				out := helper.ExecuteAndCaptureOutput(t, root)

				if !strings.Contains(string(out[:]), "m3.small.x86") &&
					!strings.Contains(string(out[:]), "m3.large.x86") &&
					!strings.Contains(string(out[:]), "c3.medium.x86") &&
					!strings.Contains(string(out[:]), "c3.small.x86") &&
					!strings.Contains(string(out[:]), "x2.xlarge.x86") &&
					!strings.Contains(string(out[:]), "x3.xlarge.x86") {
					t.Error("expected output should include m3.small.x86, m3.large.x8, c3.medium.x86, c3.small.x86, x2.xlarge.x86 and x3.xlarge.x86 by SLUG")
				}
			},
		},
		{
			name: "get_by_slug",
			fields: fields{
				MainCmd:  plans.NewClient(rootClient, outputPkg.Outputer(&outputPkg.Standard{})).NewCommand(),
				Outputer: outputPkg.Outputer(&outputPkg.Standard{}),
			},
			want: &cobra.Command{},
			cmdFunc: func(t *testing.T, c *cobra.Command) {
				root := c.Root()
				root.SetArgs([]string{subCommand, "get", "--token", os.Getenv("METAL_AUTH_TOKEN"), "--filter", "slug=m3.small.x86"})

				out := helper.ExecuteAndCaptureOutput(t, root)

				if !strings.Contains(string(out[:]), "m3.small.x86") {
					t.Error("expected output should include m3.small.x86 by SLUG")
				}
			},
		},
		{
			name: "get_by_type",
			fields: fields{
				MainCmd:  plans.NewClient(rootClient, outputPkg.Outputer(&outputPkg.Standard{})).NewCommand(),
				Outputer: outputPkg.Outputer(&outputPkg.Standard{}),
			},
			want: &cobra.Command{},
			cmdFunc: func(t *testing.T, c *cobra.Command) {
				root := c.Root()
				root.SetArgs([]string{subCommand, "get", "--token", os.Getenv("METAL_AUTH_TOKEN"), "--filter", "type=standard"})

				out := helper.ExecuteAndCaptureOutput(t, root)

				if !strings.Contains(string(out[:]), "m3.small.x86") &&
					!strings.Contains(string(out[:]), "m3.large.x86") &&
					!strings.Contains(string(out[:]), "c3.medium.x86") &&
					!strings.Contains(string(out[:]), "c3.small.x86") {
					t.Error("expected output should include m3.small.x86, m3.large.x86, c3.medium.x86 and c3.small.x86 by SLUG")
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
