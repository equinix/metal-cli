package devicecreateflagstest

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

func TestCli_Devices_Create_Flags(t *testing.T) {
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
				projectId, err = helper.CreateTestProject("metal-cli-create-flags-pro")
				if err != nil {
					t.Error(err)
				}
				if len(projectId) != 0 {

					root.SetArgs([]string{subCommand, "create", "-p", projectId, "-P", "m3.small.x86", "-m", "da", "-H", "metal-cli-create-flags-dev", "--operating-system", "custom_ipxe", "--always-pxe=true", "--ipxe-script-url", "https://boot.netboot.xyz/"})
					rescueStdout := os.Stdout
					r, w, _ := os.Pipe()
					os.Stdout = w
					if err := root.Execute(); err != nil {
						t.Error(err)
					}
					w.Close()
					out, _ := io.ReadAll(r)
					os.Stdout = rescueStdout
					if !strings.Contains(string(out[:]), "metal-cli-create-flags-dev") &&
						!strings.Contains(string(out[:]), "Ubuntu 20.04 LTS") &&
						!strings.Contains(string(out[:]), "queued") {
						t.Error("expected output should include metal-cli-create-flags-dev, Ubuntu 20.04 LTS, and queued strings in the out string ")
					}
					name := "metal-cli-create-flags-dev"
					idNamePattern := `(?m)^\| ([a-zA-Z0-9-]+) +\| *` + name + ` *\|`

					// Find the match of the ID and NAME pattern in the table string
					match := regexp.MustCompile(idNamePattern).FindStringSubmatch(string(out[:]))

					// Extract the ID from the match
					if len(match) > 1 {
						deviceId = strings.TrimSpace(match[1])
						resp, err = helper.IsDeviceStateActive(deviceId)
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
