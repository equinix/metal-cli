package devicereinstalltest

import (
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
			name: "reinstall_device",
			fields: fields{
				MainCmd:  devices.NewClient(rootClient, outputPkg.Outputer(&outputPkg.Standard{})).NewCommand(),
				Outputer: outputPkg.Outputer(&outputPkg.Standard{}),
			},
			want: &cobra.Command{},
			cmdFunc: func(t *testing.T, c *cobra.Command) {
				root := c.Root()
				projectId, err = helper.CreateTestProject("metal-cli-reinstall-pro")
				if err != nil {
					t.Error(err)
				}
				deviceId, err = helper.CreateTestDevice(projectId, "metal-cli-reinstall-dev")
				if err != nil {
					t.Error(err)
				}
				status, err = helper.IsDeviceStateActive(deviceId)
				if err != nil {
					status, err = helper.IsDeviceStateActive(deviceId)
					if err != nil {
						t.Error(err)
					}
				}

				if len(projectId) != 0 && len(deviceId) != 0 && status {
					root.SetArgs([]string{subCommand, "reinstall", "--id", deviceId, "-O", "ubuntu_22_04", "--preserve-data"})
					err = root.Execute()
					if err != nil {
						t.Error(err)
					} else {
						status, err = helper.IsDeviceStateActive(deviceId)
						// The below case will excute in both Device Active and Non-active states.
						if err != nil || status {
							if !status {
								_, err = helper.IsDeviceStateActive(deviceId)
								if err != nil {
									t.Error(err)
								}
							}
							err = helper.CleanTestDevice(deviceId)
							if err != nil {
								t.Error(err)
							}
							err = helper.CleanTestProject(projectId)
							if err != nil {
								t.Error(err)
							}
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
