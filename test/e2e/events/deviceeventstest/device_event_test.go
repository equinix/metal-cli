package eventsprojtest

import (
	"strings"
	"testing"

	root "github.com/equinix/metal-cli/internal/cli"
	"github.com/equinix/metal-cli/internal/events"
	outputPkg "github.com/equinix/metal-cli/internal/outputs"
	"github.com/equinix/metal-cli/test/helper"
	"github.com/spf13/cobra"
)

func TestCli_Events_Get(t *testing.T) {
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
			name: "get_events_by_dev_id",
			fields: fields{
				MainCmd:  events.NewClient(rootClient, outputPkg.Outputer(&outputPkg.Standard{})).NewCommand(),
				Outputer: outputPkg.Outputer(&outputPkg.Standard{}),
			},
			want: &cobra.Command{},
			cmdFunc: func(t *testing.T, c *cobra.Command) {
				root := c.Root()
				projectName := "metal-cli-device-events" + helper.GenerateRandomString(5)
				project := helper.CreateTestProject(t, projectName)
				device := helper.CreateTestDevice(t, project.GetId(), "metal-cli-events-dev")
				status, err := helper.IsDeviceStateActive(t, device.GetId())
				if err != nil || !status {
					status, err = helper.IsDeviceStateActive(t, device.GetId())
					if err != nil || !status {
						t.Fatal(err)
					}
				}
				root.SetArgs([]string{subCommand, "get", "-d", device.GetId()})

				out := helper.ExecuteAndCaptureOutput(t, root)

				if !strings.Contains(string(out[:]), "Queued for provisioning") &&
					!strings.Contains(string(out[:]), "Connected to magic install system") &&
					!strings.Contains(string(out[:]), "Provision complete! Your device is ready to go.") {
					t.Error("expected output should include Queued for provisioning in output string")
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
