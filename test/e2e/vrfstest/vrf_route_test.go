package vrfstest

import (
	"regexp"
	"strings"
	"testing"
	"time"

	root "github.com/equinix/metal-cli/internal/cli"
	outputPkg "github.com/equinix/metal-cli/internal/outputs"
	"github.com/equinix/metal-cli/internal/vrf"
	"github.com/equinix/metal-cli/test/helper"
	"github.com/spf13/cobra"
)

func TestCli_Vrf_Route(t *testing.T) {
	subCommand := "vrf"
	rootClient := root.NewClient(helper.ConsumerToken, helper.URL, helper.Version)
	randName := helper.GenerateRandomString(5)

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
			name: "vrf-create-route-test",
			fields: fields{
				MainCmd:  vrf.NewClient(rootClient, outputPkg.Outputer(&outputPkg.Standard{})).NewCommand(),
				Outputer: outputPkg.Outputer(&outputPkg.Standard{}),
			},
			want: &cobra.Command{},
			cmdFunc: func(t *testing.T, c *cobra.Command) {
				root := c.Root()
				projName := "metal-cli-" + randName + "-vrf-create-route-test"
				projectId := helper.CreateTestProject(t, projName)
				if projectId.GetId() != "" {
					vlan := helper.CreateTestVlanWithVxLan(t, projectId.GetId(), 3987, projName)
					vrf := helper.CreateTestVrfs(t, projectId.GetId(), projName, 3987)

					ipReservation := helper.CreateTestVrfIpRequest(t, projectId.GetId(), vrf.GetId())
					_ = helper.CreateTestVrfGateway(t, projectId.GetId(), ipReservation.VrfIpReservation.GetId(), vlan.GetId())

					if vlan.GetId() != "" && vrf.GetId() != "" {
						root.SetArgs([]string{subCommand, "create-route", "-i", vrf.GetId(), "-p", "0.0.0.0/0", "-n", "10.10.1.2", "-t", "foo"})

						out := helper.ExecuteAndCaptureOutput(t, root)

						if !strings.Contains(string(out[:]), "TYPE") &&
							!strings.Contains(string(out[:]), "static") &&
							!strings.Contains(string(out[:]), "PREFIX") &&
							!strings.Contains(string(out[:]), "0.0.0.0/0") {
							t.Error("expected output should include TYPE static PREFIX and 0.0.0.0/0, in the out string ")
						}

						idNamePattern := `(?m)^\| ([a-zA-Z0-9-]+) +\| *`

						// Find the match of the ID and NAME pattern in the table string
						match := regexp.MustCompile(idNamePattern).FindStringSubmatch(string(out[:]))
						// Extract the ID from the match
						if len(match) > 1 {
							routeId := strings.TrimSpace(match[1])
							helper.CleanTestVrfRoute(t, routeId)
						} else {
							t.Errorf("No match found for %v in %v", idNamePattern, string(out[:]))
						}
					}
				}
			},
		},
		{
			name: "vrf-delete-route-test",
			fields: fields{
				MainCmd:  vrf.NewClient(rootClient, outputPkg.Outputer(&outputPkg.Standard{})).NewCommand(),
				Outputer: outputPkg.Outputer(&outputPkg.Standard{}),
			},
			want: &cobra.Command{},
			cmdFunc: func(t *testing.T, c *cobra.Command) {
				root := c.Root()

				projName := "metal-cli-" + randName + "-vrf-delete-test"
				projectId := helper.CreateTestProject(t, projName)

				if projectId.GetId() != "" {
					vlan := helper.CreateTestVlanWithVxLan(t, projectId.GetId(), 3987, projName)
					vrf := helper.CreateTestVrfs(t, projectId.GetId(), projName, 3987)

					ipReservation := helper.CreateTestVrfIpRequest(t, projectId.GetId(), vrf.GetId())
					_ = helper.CreateTestVrfGateway(t, projectId.GetId(), ipReservation.VrfIpReservation.GetId(), vlan.GetId())
					vrfRoute := helper.CreateTestVrfRoute(t, vrf.GetId())
					if vlan.GetId() != "" && vrf.GetId() != "" && vrfRoute.GetId() != "" {
						root.SetArgs([]string{subCommand, "delete-route", "-i", vrfRoute.GetId()})

						out := helper.ExecuteAndCaptureOutput(t, root)

						if !strings.Contains(string(out[:]), vrfRoute.GetId()) {
							t.Error("expected output should include VRF Route deletion initiated. Please check 'metal vrf GetRoute -i " + vrfRoute.GetId() + " ' for status, in the out string ")
						}
					}
				}
			},
		},
		{
			name: "vrf-update-route-test",
			fields: fields{
				MainCmd:  vrf.NewClient(rootClient, outputPkg.Outputer(&outputPkg.Standard{})).NewCommand(),
				Outputer: outputPkg.Outputer(&outputPkg.Standard{}),
			},
			want: &cobra.Command{},
			cmdFunc: func(t *testing.T, c *cobra.Command) {
				root := c.Root()

				projName := "metal-cli-" + randName + "-vrf-list-test"
				projectId := helper.CreateTestProject(t, projName)

				if projectId.GetId() != "" {
					vlan := helper.CreateTestVlanWithVxLan(t, projectId.GetId(), 3987, projName)
					vrf := helper.CreateTestVrfs(t, projectId.GetId(), projName, 3987)

					ipReservation := helper.CreateTestVrfIpRequest(t, projectId.GetId(), vrf.GetId())
					_ = helper.CreateTestVrfGateway(t, projectId.GetId(), ipReservation.VrfIpReservation.GetId(), vlan.GetId())
					route := helper.CreateTestVrfRoute(t, vrf.GetId())

					// We literally need to sleep for 5 minutes; the API will reject any
					// VRF route update request that comes in less than 5 minutes after
					// the VRF route was last updated
					time.Sleep(300 * time.Second)

					root.SetArgs([]string{subCommand, "update-route", "-i", route.GetId(), "-t", "foobar"})

					out := helper.ExecuteAndCaptureOutput(t, root)

					if !strings.Contains(string(out[:]), "TYPE") &&
						!strings.Contains(string(out[:]), "static") &&
						!strings.Contains(string(out[:]), "PREFIX") &&
						!strings.Contains(string(out[:]), "0.0.0.0/0") {
						t.Error("expected output should include TYPE static PREFIX and 0.0.0.0/0, in the out string ")
					}
				}
			},
		},
		{
			name: "vrf-get-route-test",
			fields: fields{
				MainCmd:  vrf.NewClient(rootClient, outputPkg.Outputer(&outputPkg.Standard{})).NewCommand(),
				Outputer: outputPkg.Outputer(&outputPkg.Standard{}),
			},
			want: &cobra.Command{},
			cmdFunc: func(t *testing.T, c *cobra.Command) {
				root := c.Root()

				projName := "metal-cli-" + randName + "-vrf-list-test"
				projectId := helper.CreateTestProject(t, projName)

				if projectId.GetId() != "" {
					vlan := helper.CreateTestVlanWithVxLan(t, projectId.GetId(), 3987, projName)
					vrf := helper.CreateTestVrfs(t, projectId.GetId(), projName, 3987)

					ipReservation := helper.CreateTestVrfIpRequest(t, projectId.GetId(), vrf.GetId())
					_ = helper.CreateTestVrfGateway(t, projectId.GetId(), ipReservation.VrfIpReservation.GetId(), vlan.GetId())
					// vrfRoute := helper.CreateTestVrfRoute(t, vrf.GetId())
					_ = helper.CreateTestVrfRoute(t, vrf.GetId())

					if vlan.GetId() != "" && vrf.GetId() != "" {
						root.SetArgs([]string{subCommand, "get-route", "-i", vrf.GetId()})

						out := helper.ExecuteAndCaptureOutput(t, root)

						if !strings.Contains(string(out[:]), "TYPE") &&
							!strings.Contains(string(out[:]), "static") &&
							!strings.Contains(string(out[:]), "PREFIX") &&
							!strings.Contains(string(out[:]), "ID") &&
							!strings.Contains(string(out[:]), vrf.GetId()) &&
							!strings.Contains(string(out[:]), "0.0.0.0/0") {
							t.Error("expected output should include TYPE static PREFIX and 0.0.0.0/0, in the out string ")
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
