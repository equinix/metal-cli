package plantest

import (
	"io"
	"os"
	"strings"
	"testing"

	root "github.com/equinix/metal-cli/internal/cli"
	outputPkg "github.com/equinix/metal-cli/internal/outputs"
	"github.com/equinix/metal-cli/internal/plans"
	"github.com/spf13/cobra"
)

func TestCli_Plans(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	subCommand := "plans"
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
				MainCmd:  plans.NewClient(rootClient, outputPkg.Outputer(&outputPkg.Standard{})).NewCommand(),
				Outputer: outputPkg.Outputer(&outputPkg.Standard{}),
			},
			want: &cobra.Command{},
			cmdFunc: func(t *testing.T, c *cobra.Command) {
				root := c.Root()
				root.SetArgs([]string{subCommand, "get"})
				rescueStdout := os.Stdout
				r, w, _ := os.Pipe()
				os.Stdout = w
				if err := root.Execute(); err != nil {
					t.Error(err)
				}
				w.Close()
				out, _ := io.ReadAll(r)
				os.Stdout = rescueStdout
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
				rescueStdout := os.Stdout
				r, w, _ := os.Pipe()
				os.Stdout = w
				if err := root.Execute(); err != nil {
					t.Error(err)
				}
				w.Close()
				out, _ := io.ReadAll(r)
				os.Stdout = rescueStdout
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
				rescueStdout := os.Stdout
				r, w, _ := os.Pipe()
				os.Stdout = w
				if err := root.Execute(); err != nil {
					t.Error(err)
				}
				w.Close()
				out, _ := io.ReadAll(r)
				os.Stdout = rescueStdout
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
