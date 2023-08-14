package devicestarttest

import (
	"fmt"
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
	var status bool
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
				projectId, err = helper.CreateTestProject("metal-cli-start-pro")
				if err != nil {
					t.Error(err)
				}
				deviceId, err = helper.CreateTestDevice(projectId, "metal-cli-start-dev")
				if err != nil {
					t.Error(err)
				}
				status, err = helper.IsDeviceStateActive(deviceId)
				if err != nil {
					_, err := helper.IsDeviceStateActive(deviceId)
					if err != nil {
						t.Error(err)
					} else {
						err = helper.StopTestDevice(deviceId)
						if err != nil {
							t.Error(err)
						}
						status, err = helper.IsDeviceStateActive(deviceId)
						if err == nil {
							t.Error(err)
						}
					}
				}

				if len(projectId) != 0 && len(deviceId) != 0 && !status {
					root.SetArgs([]string{subCommand, "start", "--id", deviceId})
					rescueStdout := os.Stdout
					r, w, _ := os.Pipe()
					os.Stdout = w
					if err := root.Execute(); err != nil {
						t.Error(err)
					}
					w.Close()
					out, _ := io.ReadAll(r)
					os.Stdout = rescueStdout
					if !strings.Contains(string(out[:]), "Device "+deviceId+" successfully started.") {
						t.Error("expected output should include" + "Device " + deviceId + " successfully started." + "in the out string ")
					} else {
						status, _ = helper.IsDeviceStateActive(deviceId)
						if err != nil || status {
							if !status {
								_, err = helper.IsDeviceStateActive(deviceId)
								if err != nil {
									t.Error(err)
								}
							}
							fmt.Print("Device is Active")
							err = helper.CleanTestDevice(deviceId)
							if err != nil {
								t.Error(err)
							}
							fmt.Print("Cleaned Test Device")
							err = helper.CleanTestProject(projectId)
							if err != nil {
								t.Error(err)
							}
							fmt.Print("Cleaned Test Project")
						}
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
