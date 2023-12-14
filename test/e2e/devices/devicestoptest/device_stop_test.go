package devicestoptest

import (
	"io"
	"os"
	"strings"
	"testing"

	root "github.com/equinix/metal-cli/internal/cli"
	"github.com/equinix/metal-cli/internal/devices"
	outputPkg "github.com/equinix/metal-cli/internal/outputs"
	"github.com/equinix/metal-cli/test/helper"
	"github.com/spf13/cobra"
)

func TestCli_Devices_Update(t *testing.T) {
	var projectId, deviceId string
	var err error
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
			name: "stop_device",
			fields: fields{
				MainCmd:  devices.NewClient(rootClient, outputPkg.Outputer(&outputPkg.Standard{})).NewCommand(),
				Outputer: outputPkg.Outputer(&outputPkg.Standard{}),
			},
			want: &cobra.Command{},
			cmdFunc: func(t *testing.T, c *cobra.Command) {
				root := c.Root()
				projectName := "metal-cli-stop-pro" + helper.GenerateUUID()
				projectId, err = helper.CreateTestProject(t, projectName)
				t.Cleanup(func() {
					if err := helper.CleanTestProject(t, projectId); err != nil &&
						!strings.Contains(err.Error(), "Not Found") {
						t.Error(err)
					}
				})
				if err != nil {
					t.Fatal(err)
				}

				deviceId, err = helper.CreateTestDevice(t, projectId, "metal-cli-stop-dev")
				t.Cleanup(func() {
					if err := helper.CleanTestDevice(t, deviceId); err != nil &&
						!strings.Contains(err.Error(), "Not Found") {
						t.Error(err)
					}
				})
				if err != nil {
					t.Fatal(err)
				}

				status, err := helper.IsDeviceStateActive(t, deviceId)
				if err != nil {
					t.Fatal(err)
				}
				if len(projectId) != 0 && len(deviceId) != 0 && status {
					root.SetArgs([]string{subCommand, "stop", "--id", deviceId})
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
					if !strings.Contains(string(out[:]), "Device "+deviceId+" successfully stopped.") {
						t.Fatal("expected output should include" + "Device " + deviceId + " successfully stopped." + "in the out string ")
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
