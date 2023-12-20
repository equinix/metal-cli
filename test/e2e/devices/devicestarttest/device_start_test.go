package devicestarttest

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
			name: "start_device",
			fields: fields{
				MainCmd:  devices.NewClient(rootClient, outputPkg.Outputer(&outputPkg.Standard{})).NewCommand(),
				Outputer: outputPkg.Outputer(&outputPkg.Standard{}),
			},
			want: &cobra.Command{},
			cmdFunc: func(t *testing.T, c *cobra.Command) {
				root := c.Root()
				projectName := "metal-cli-device-start" + helper.GenerateRandomString(5)
				project := helper.CreateTestProject(t, projectName)
				device := helper.CreateTestDevice(t, project.GetId(), "metal-cli-start-dev")
				status, err := helper.IsDeviceStateActive(t, device.GetId())

				if err != nil {
					_, err := helper.IsDeviceStateActive(t, device.GetId())
					if err != nil {
						t.Error(err)
					} else {
						err = helper.StopTestDevice(t, device.GetId())
						if err != nil {
							t.Error(err)
						}
						status, err = helper.IsDeviceStateActive(t, device.GetId())
						if err == nil {
							t.Error(err)
						}
					}
				}

				if !status {
					root.SetArgs([]string{subCommand, "start", "--id", device.GetId()})

					out := helper.ExecuteAndCaptureOutput(t, root)

					if !strings.Contains(string(out[:]), "Device "+device.GetId()+" successfully started.") {
						t.Fatal("expected output should include" + "Device " + device.GetId() + " successfully started." + "in the out string ")
					}

					status, err = helper.IsDeviceStateActive(t, device.GetId())
					if err != nil {
						t.Fatal(err)
					}
					if !status {
						t.Fatalf("Device not yet active, %s", device.GetId())
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
