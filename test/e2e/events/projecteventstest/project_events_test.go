package eventsprojtest

import (
	"io"
	"os"
	"strings"
	"testing"

	root "github.com/equinix/metal-cli/internal/cli"
	"github.com/equinix/metal-cli/internal/events"
	outputPkg "github.com/equinix/metal-cli/internal/outputs"
	"github.com/equinix/metal-cli/test/helper"
	"github.com/spf13/cobra"
)

func TestCli_Events_Get(t *testing.T) {
	var projectId string
	var err error
	subCommand := "event"
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
			name: "get_events_by_proj_id",
			fields: fields{
				MainCmd:  events.NewClient(rootClient, outputPkg.Outputer(&outputPkg.Standard{})).NewCommand(),
				Outputer: outputPkg.Outputer(&outputPkg.Standard{}),
			},
			want: &cobra.Command{},
			cmdFunc: func(t *testing.T, c *cobra.Command) {
				root := c.Root()

				projectName := "metal-cli-events-pro" + helper.GenerateUUID()
				projectId, err = helper.CreateTestProject(t, projectName)
				if err != nil {
					t.Error(err)
				}
				root.SetArgs([]string{subCommand, "get", "-p", projectId})
				rescueStdout := os.Stdout
				r, w, _ := os.Pipe()
				os.Stdout = w
				t.Cleanup(func() {
					w.Close()
					os.Stdout = rescueStdout
				})

				if err := root.Execute(); err != nil {
					t.Fatal(err)
				}

				out, _ := io.ReadAll(r)
				if !strings.Contains(string(out[:]), "metal-cli-events-pro") {
					t.Error("expected output should include metal-cli-events-pro in output string")
				}
				err = helper.CleanTestProject(t, projectId)
				if err != nil {
					t.Error(err)
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
