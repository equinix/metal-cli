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

func TestCli_Devices_create(t *testing.T) {
	var projectId, deviceId string
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
				projectId = helper.Create_test_project("create-device-pro")
				if len(projectId) != 0 {

					root.SetArgs([]string{subCommand, "create", "-p", projectId, "-P", "c3.small.x86", "-m", "da", "-O", "ubuntu_20_04", "-H", "create-device-dev"})
					rescueStdout := os.Stdout
					r, w, _ := os.Pipe()
					os.Stdout = w
					if err := root.Execute(); err != nil {
						t.Error(err)
					}
					w.Close()
					out, _ := io.ReadAll(r)
					os.Stdout = rescueStdout
					if !strings.Contains(string(out[:]), "create-device-dev") &&
						!strings.Contains(string(out[:]), "Ubuntu 20.04 LTS") &&
						!strings.Contains(string(out[:]), "queued") {
						t.Error("expected output should include create-device-dev, Ubuntu 20.04 LTS, and queued strings in the out string ")
					}
					name := "create-device-dev"
					idNamePattern := `(?m)^\| ([a-zA-Z0-9-]+) +\| *` + name + ` *\|`

					// Find the match of the ID and NAME pattern in the table string
					match := regexp.MustCompile(idNamePattern).FindStringSubmatch(string(out[:]))

					// Extract the ID from the match
					if len(match) > 1 {
						deviceId = strings.TrimSpace(match[1])
						if helper.Is_Device_state_active(deviceId) {
							helper.Clean_test_device(deviceId)
							helper.Clean_test_project(projectId)
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
