package e2e

import (
	"bytes"
	"testing"

	root "github.com/equinix/metal-cli/internal/cli"
	"github.com/equinix/metal-cli/internal/os"
	outputPkg "github.com/equinix/metal-cli/internal/outputs"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestCli_OperatingSystem(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
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
				MainCmd:  os.NewClient(rootClient, outputPkg.Outputer(&outputPkg.Standard{})).NewCommand(),
				Outputer: outputPkg.Outputer(&outputPkg.Standard{}),
			},
			want: &cobra.Command{},
			cmdFunc: func(t *testing.T, c *cobra.Command) {
				root := c.Root()
				out := &bytes.Buffer{}
				root.SetArgs([]string{subCommand, "get"})
				root.SetOutput(out)
				err := root.Execute()
				assert.NoError(t, err)
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
