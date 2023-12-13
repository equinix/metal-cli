package devicecreatetest

import (
	"io"
	"os"
	"regexp"
	"strings"
	"testing"

	root "github.com/equinix/metal-cli/internal/cli"
	"github.com/equinix/metal-cli/internal/devices"
	outputPkg "github.com/equinix/metal-cli/internal/outputs"
	"github.com/equinix/metal-cli/test/helper"
	"github.com/spf13/cobra"
)

func TestCli_Devices_Create(t *testing.T) {
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
			name: "create_device",
			fields: fields{
				MainCmd:  devices.NewClient(rootClient, outputPkg.Outputer(&outputPkg.Standard{})).NewCommand(),
				Outputer: outputPkg.Outputer(&outputPkg.Standard{}),
			},
			want: &cobra.Command{},
			cmdFunc: func(t *testing.T, c *cobra.Command) {
				root := c.Root()
				projectName := "metal-cli-create-pro" + helper.GenerateUUID()
				projectId, err = helper.CreateTestProject(t, projectName)
				t.Cleanup(func() {
					if err := helper.CleanTestProject(t, projectId); err != nil &&
						!strings.Contains(err.Error(), "Not Found") {
						t.Error(err)
					}
				})
				if err != nil {
					t.Error(err)
					return
				}

				if len(projectId) != 0 {

					deviceName := "metal-cli-create-dev" + helper.GenerateUUID()
					root.SetArgs([]string{subCommand, "create", "-p", projectId, "-P", "m3.small.x86", "-m", "da", "-O", "ubuntu_20_04", "-H", deviceName})
					rescueStdout := os.Stdout
					r, w, _ := os.Pipe()
					os.Stdout = w
					if err := root.Execute(); err != nil {
						t.Error(err)
						return
					}
					w.Close()
					out, _ := io.ReadAll(r)
					os.Stdout = rescueStdout

					t.Cleanup(func() {
						if err := helper.CleanTestDevice(t, deviceId); err != nil &&
							!strings.Contains(err.Error(), "Not Found") {
							t.Error(err)
						}
					})

					if !strings.Contains(string(out[:]), deviceName) &&
						!strings.Contains(string(out[:]), "Ubuntu 20.04 LTS") &&
						!strings.Contains(string(out[:]), "queued") {
						t.Errorf("expected output should include %s, Ubuntu 20.04 LTS, and queued strings in the out string ", deviceName)
					}

					idNamePattern := `(?m)^\| ([a-zA-Z0-9-]+) +\| *` + deviceName + ` *\|`

					// Find the match of the ID and NAME pattern in the table string
					match := regexp.MustCompile(idNamePattern).FindStringSubmatch(string(out[:]))

					// Extract the ID from the match
					if len(match) > 1 {
						deviceId = strings.TrimSpace(match[1])
						resp, err = helper.IsDeviceStateActive(t, deviceId)
						if err != nil || resp {
							if !resp {
								resp, err = helper.IsDeviceStateActive(t, deviceId)
							}
							err = helper.CleanTestDevice(t, deviceId)
							if err != nil {
								t.Error(err)
							}
							err = helper.CleanTestProject(t, projectId)
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
