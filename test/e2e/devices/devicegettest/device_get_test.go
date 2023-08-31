package devicestest

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

func TestCli_Devices_Get(t *testing.T) {
	var projectId, deviceId string
	var err error
	var resp bool
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
			name: "get_by_device_id",
			fields: fields{
				MainCmd:  devices.NewClient(rootClient, outputPkg.Outputer(&outputPkg.Standard{})).NewCommand(),
				Outputer: outputPkg.Outputer(&outputPkg.Standard{}),
			},
			want: &cobra.Command{},
			cmdFunc: func(t *testing.T, c *cobra.Command) {
				root := c.Root()
				projectId, err = helper.CreateTestProject("metal-cli-get-pro")
				if err != nil {
					t.Error(err)
				}
				deviceId, err = helper.CreateTestDevice(projectId, "metal-cli-get-dev")
				if err != nil {
					t.Error(err)
				}
				root.SetArgs([]string{subCommand, "get", "-i", deviceId})
				rescueStdout := os.Stdout
				r, w, _ := os.Pipe()
				os.Stdout = w
				if err := root.Execute(); err != nil {
					t.Error(err)
				}
				w.Close()
				out, _ := io.ReadAll(r)
				os.Stdout = rescueStdout
				if !strings.Contains(string(out[:]), deviceId) {
					t.Error("expected output should include " + deviceId)
				}
				if len(projectId) != 0 && len(deviceId) != 0 {
					resp, err = helper.IsDeviceStateActive(deviceId, "active")
					if err == nil && resp == true {
						err = helper.CleanTestDevice(deviceId)
						if err != nil {
							t.Error(err)
						}
						err = helper.CleanTestProject(projectId)
						if err != nil {
							t.Error(err)
						}
					} else {
						t.Error(err)
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
