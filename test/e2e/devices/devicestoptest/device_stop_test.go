package devicestoptest

import (
	"strings"
	"testing"

	root "github.com/equinix/metal-cli/internal/cli"
	"github.com/equinix/metal-cli/internal/devices"
	outputPkg "github.com/equinix/metal-cli/internal/outputs"
	"github.com/equinix/metal-cli/test/helper"
	"github.com/spf13/cobra"
)

func TestCli_Devices_Update(t *testing.T) {
	subCommand := "device"
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
			name: "stop_device",
			fields: fields{
				MainCmd:  devices.NewClient(rootClient, outputPkg.Outputer(&outputPkg.Standard{})).NewCommand(),
				Outputer: outputPkg.Outputer(&outputPkg.Standard{}),
			},
			want: &cobra.Command{},
			cmdFunc: func(t *testing.T, c *cobra.Command) {
				root := c.Root()
				projectName := "metal-cli-" + helper.GenerateRandomString(5) + "-device-stop"
				project := helper.CreateTestProject(t, projectName)
				device := helper.CreateTestDevice(t, project.GetId(), "metal-cli-stop-dev")

				status, err := helper.IsDeviceStateActive(t, device.GetId())
				if err != nil {
					t.Fatal(err)
				}
				if status {
					root.SetArgs([]string{subCommand, "stop", "--id", device.GetId()})

					out := helper.ExecuteAndCaptureOutput(t, root)

					if !strings.Contains(string(out[:]), "Device "+device.GetId()+" successfully stopped.") {
						t.Fatal("expected output should include" + "Device " + device.GetId() + " successfully stopped." + "in the out string ")
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
