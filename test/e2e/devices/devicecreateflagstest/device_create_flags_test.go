package devicecreateflagstest

import (
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
	var deviceId string
	var err error
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
			name: "create_device",
			fields: fields{
				MainCmd:  devices.NewClient(rootClient, outputPkg.Outputer(&outputPkg.Standard{})).NewCommand(),
				Outputer: outputPkg.Outputer(&outputPkg.Standard{}),
			},
			want: &cobra.Command{},
			cmdFunc: func(t *testing.T, c *cobra.Command) {
				root := c.Root()
				projectName := "metal-cli-" + helper.GenerateRandomString(5) + "-device-create-flags"
				project := helper.CreateTestProject(t, projectName)

				root.SetArgs([]string{subCommand, "create", "-p", project.GetId(), "-P", "m3.small.x86", "-m", "da", "-H", "metal-cli-create-flags-dev", "--operating-system", "custom_ipxe", "--always-pxe=true", "--ipxe-script-url", "https://boot.netboot.xyz/"})

				out := helper.ExecuteAndCaptureOutput(t, root)

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
					_, err = helper.IsDeviceStateActive(t, deviceId)
					t.Cleanup(func() {
						helper.CleanTestDevice(t, deviceId)
					})
					if err != nil {
						t.Fatal(err)
					}
				} else {
					t.Errorf("No match found for %v in %v", idNamePattern, string(out[:]))
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
