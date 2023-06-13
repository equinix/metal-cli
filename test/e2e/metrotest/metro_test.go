package metrotest

import (
	"io"
	"os"
	"strings"
	"testing"

	root "github.com/equinix/metal-cli/internal/cli"
	"github.com/equinix/metal-cli/internal/metros"
	outputPkg "github.com/equinix/metal-cli/internal/outputs"
	"github.com/spf13/cobra"
)

func TestCli_Metros(t *testing.T) {
	subCommand := "metros"
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
				MainCmd:  metros.NewClient(rootClient, outputPkg.Outputer(&outputPkg.Standard{})).NewCommand(),
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
				if !strings.Contains(string(out[:]), "Seattle") &&
					!strings.Contains(string(out[:]), "Tokyo") &&
					!strings.Contains(string(out[:]), "Sydney") &&
					!strings.Contains(string(out[:]), "Toronto") &&
					!strings.Contains(string(out[:]), "Chicago") &&
					!strings.Contains(string(out[:]), "Dallas") {
					t.Error("expected output should include Seattle, Tokyo, Sydney, Toronto, Chicago and Dallas by NAME")
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
