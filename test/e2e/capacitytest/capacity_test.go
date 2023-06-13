package capacitytest

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/equinix/metal-cli/internal/capacity"
	root "github.com/equinix/metal-cli/internal/cli"
	outputPkg "github.com/equinix/metal-cli/internal/outputs"
	"github.com/spf13/cobra"
)

func TestCli_Capacity(t *testing.T) {
	subCommand := "capacity"
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
				MainCmd:  capacity.NewClient(rootClient, outputPkg.Outputer(&outputPkg.Standard{})).NewCommand(),
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
				if !strings.Contains(string(out[:]), "n3.xlarge.x86") &&
					!strings.Contains(string(out[:]), "m3.large.x86") &&
					!strings.Contains(string(out[:]), "s3.xlarge.x86") &&
					!strings.Contains(string(out[:]), "sv16") &&
					!strings.Contains(string(out[:]), "sv16") &&
					!strings.Contains(string(out[:]), "dc10") {
					t.Error("expected output should include n3.xlarge.x86, m3.large.x86, s3.xlarge.x86, dc10.")
				}
			},
		},
		{
			name: "get_by_plan_1",
			fields: fields{
				MainCmd:  capacity.NewClient(rootClient, outputPkg.Outputer(&outputPkg.Standard{})).NewCommand(),
				Outputer: outputPkg.Outputer(&outputPkg.Standard{}),
			},
			want: &cobra.Command{},
			cmdFunc: func(t *testing.T, c *cobra.Command) {
				root := c.Root()
				root.SetArgs([]string{subCommand, "get", "-m", "-P", "c3.small.x86"})
				rescueStdout := os.Stdout
				r, w, _ := os.Pipe()
				os.Stdout = w
				if err := root.Execute(); err != nil {
					t.Error(err)
				}
				w.Close()
				out, _ := io.ReadAll(r)
				os.Stdout = rescueStdout
				if !strings.Contains(string(out[:]), "c3.small.x86") &&
					!strings.Contains(string(out[:]), "mt") &&
					!strings.Contains(string(out[:]), "sv") &&
					!strings.Contains(string(out[:]), "md") &&
					!strings.Contains(string(out[:]), "sg") &&
					!strings.Contains(string(out[:]), "pa") {
					t.Error("expected output should include c3.small.x86, mt, sv, md, sg, pa")
				}
			},
		},
		{
			name: "get_by_plan_2",
			fields: fields{
				MainCmd:  capacity.NewClient(rootClient, outputPkg.Outputer(&outputPkg.Standard{})).NewCommand(),
				Outputer: outputPkg.Outputer(&outputPkg.Standard{}),
			},
			want: &cobra.Command{},
			cmdFunc: func(t *testing.T, c *cobra.Command) {
				root := c.Root()
				root.SetArgs([]string{subCommand, "get", "-m", "-P", "m3.large.x86"})
				rescueStdout := os.Stdout
				r, w, _ := os.Pipe()
				os.Stdout = w
				if err := root.Execute(); err != nil {
					t.Error(err)
				}
				w.Close()
				out, _ := io.ReadAll(r)
				os.Stdout = rescueStdout
				if !strings.Contains(string(out[:]), "m3.large.x86") &&
					!strings.Contains(string(out[:]), "mt") &&
					!strings.Contains(string(out[:]), "sv") &&
					!strings.Contains(string(out[:]), "md") &&
					!strings.Contains(string(out[:]), "sg") {
					t.Error("expected output should include c3.small.x86, mt, sv, md, sg")
				}
			},
		},
		{
			name: "check_by_multi_metro",
			fields: fields{
				MainCmd:  capacity.NewClient(rootClient, outputPkg.Outputer(&outputPkg.Standard{})).NewCommand(),
				Outputer: outputPkg.Outputer(&outputPkg.Standard{}),
			},
			want: &cobra.Command{},
			cmdFunc: func(t *testing.T, c *cobra.Command) {
				root := c.Root()
				root.SetArgs([]string{subCommand, "check", "-m", "ny,da", "-P", "c3.medium.x86", "-q", "10"})
				rescueStdout := os.Stdout
				r, w, _ := os.Pipe()
				os.Stdout = w
				if err := root.Execute(); err != nil {
					t.Error(err)
				}
				w.Close()
				out, _ := io.ReadAll(r)
				os.Stdout = rescueStdout
				if !strings.Contains(string(out[:]), "c3.medium.x86") &&
					!strings.Contains(string(out[:]), "ny") &&
					!strings.Contains(string(out[:]), "da") {
					t.Error("expected output should include c3.medium.x86, ny, da")
				}
			},
		},
		{
			name: "check_by_multi_plan",
			fields: fields{
				MainCmd:  capacity.NewClient(rootClient, outputPkg.Outputer(&outputPkg.Standard{})).NewCommand(),
				Outputer: outputPkg.Outputer(&outputPkg.Standard{}),
			},
			want: &cobra.Command{},
			cmdFunc: func(t *testing.T, c *cobra.Command) {
				root := c.Root()
				root.SetArgs([]string{subCommand, "check", "-m", "da", "-P", "c3.medium.x86,m3.large.x86", "-q", "10"})
				rescueStdout := os.Stdout
				r, w, _ := os.Pipe()
				os.Stdout = w
				if err := root.Execute(); err != nil {
					t.Error(err)
				}
				w.Close()
				out, _ := io.ReadAll(r)
				os.Stdout = rescueStdout
				if !strings.Contains(string(out[:]), "c3.medium.x86") &&
					!strings.Contains(string(out[:]), "m3.large.x86") &&
					!strings.Contains(string(out[:]), "ny") &&
					!strings.Contains(string(out[:]), "da") {
					t.Error("expected output should include c3.medium.x86, m3.large.x86, da")
				}
			},
		},
		{
			name: "check_by_multi_metro_and_plan",
			fields: fields{
				MainCmd:  capacity.NewClient(rootClient, outputPkg.Outputer(&outputPkg.Standard{})).NewCommand(),
				Outputer: outputPkg.Outputer(&outputPkg.Standard{}),
			},
			want: &cobra.Command{},
			cmdFunc: func(t *testing.T, c *cobra.Command) {
				root := c.Root()
				root.SetArgs([]string{subCommand, "check", "-m", "ny,da", "-P", "c3.medium.x86,m3.large.x86", "-q", "10"})
				rescueStdout := os.Stdout
				r, w, _ := os.Pipe()
				os.Stdout = w
				if err := root.Execute(); err != nil {
					t.Error(err)
				}
				w.Close()
				out, _ := io.ReadAll(r)
				os.Stdout = rescueStdout
				if !strings.Contains(string(out[:]), "c3.medium.x86") &&
					!strings.Contains(string(out[:]), "m3.large.x86") &&
					!strings.Contains(string(out[:]), "ny") &&
					!strings.Contains(string(out[:]), "da") {
					t.Error("expected output should include c3.medium.x86, m3.large.x86, ny, da")
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
