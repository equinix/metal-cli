package vrfstest

import (
	"strings"
	"testing"

	root "github.com/equinix/metal-cli/internal/cli"
	outputPkg "github.com/equinix/metal-cli/internal/outputs"
	"github.com/equinix/metal-cli/internal/vrf"
	"github.com/equinix/metal-cli/test/helper"
	"github.com/spf13/cobra"
)

func TestCli_Vrf_Create(t *testing.T) {
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
			name: "vrf-create-test",
			fields: fields{
				MainCmd:  vrf.NewClient(rootClient, outputPkg.Outputer(&outputPkg.Standard{})).NewCommand(),
				Outputer: outputPkg.Outputer(&outputPkg.Standard{}),
			},
			want: &cobra.Command{},
			cmdFunc: func(t *testing.T, c *cobra.Command) {
				root := c.Root()
				projName := "metal-cli-" + randName + "-vrf-create-test"
				projectId := helper.CreateTestProject(t, projName)
				if projectId.GetId() != "" {

					root.SetArgs([]string{subCommand, "create", "-p", projectId.GetId(), "-m", "da", "-n", projName, "-a", "3456", "-r", "10.0.1.0/24"})

					out := helper.ExecuteAndCaptureOutput(t, root)

					if !strings.Contains(string(out[:]), projName) {
						t.Error("expected output should include " + projName + ", in the out string ")
					}
				}
			},
		},
		{
			name: "vrf-delete-test",
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
					vlan := helper.CreateTestVlanWithVxLan(t, projectId.GetId(), 3983, projName)
					vrf := helper.CreateTestVrfs(t, projectId.GetId(), projName, 3983)
					if vrf.GetId() != "" && vlan.GetId() != "" {
						root.SetArgs([]string{subCommand, "delete", "-i", vrf.GetId(), "-f"})
						out := helper.ExecuteAndCaptureOutput(t, root)

						if !strings.Contains(string(out[:]), "VRF deletion initiated. Please check 'metal vrf get -i "+vrf.GetId()+" ' for status") {
							t.Error("expected output should include VRF deletion initiated. Please check 'metal vrf get -i " + vrf.GetId() + " ' for status, in the out string")
						}
					}
				}
			},
		},
		{
			name: "vrf-list-test",
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
					_ = helper.CreateTestVlanWithVxLan(t, projectId.GetId(), 3982, projName)
					vrf := helper.CreateTestVrfs(t, projectId.GetId(), projName, 3982)

					root.SetArgs([]string{subCommand, "get", "-p", projectId.GetId()})
					out := helper.ExecuteAndCaptureOutput(t, root)

					if !strings.Contains(string(out[:]), vrf.GetId()) &&
						!strings.Contains(string(out[:]), projName) {
						t.Error("expected output should include " + vrf.GetId() + ", in the out string ")
					}
				}
			},
		},
		{
			name: "vrf-getByProjectId-test",
			fields: fields{
				MainCmd:  vrf.NewClient(rootClient, outputPkg.Outputer(&outputPkg.Standard{})).NewCommand(),
				Outputer: outputPkg.Outputer(&outputPkg.Standard{}),
			},
			want: &cobra.Command{},
			cmdFunc: func(t *testing.T, c *cobra.Command) {
				root := c.Root()

				projName := "metal-cli-" + randName + "-vrf-get-test"
				projectId := helper.CreateTestProject(t, projName)

				if projectId.GetId() != "" {
					_ = helper.CreateTestVlanWithVxLan(t, projectId.GetId(), 3981, projName)
					vrf := helper.CreateTestVrfs(t, projectId.GetId(), projName, 3981)

					root.SetArgs([]string{subCommand, "get", "-p", projectId.GetId()})
					out := helper.ExecuteAndCaptureOutput(t, root)

					if !strings.Contains(string(out[:]), vrf.GetId()) &&
						!strings.Contains(string(out[:]), projName) {
						t.Error("expected output should include " + vrf.GetId() + ", in the out string ")
					}

				}
			},
		},
		{
			name: "vrf-getById-test",
			fields: fields{
				MainCmd:  vrf.NewClient(rootClient, outputPkg.Outputer(&outputPkg.Standard{})).NewCommand(),
				Outputer: outputPkg.Outputer(&outputPkg.Standard{}),
			},
			want: &cobra.Command{},
			cmdFunc: func(t *testing.T, c *cobra.Command) {
				root := c.Root()

				projName := "metal-cli-" + randName + "-vrf-get-test"
				projectId := helper.CreateTestProject(t, projName)

				if projectId.GetId() != "" {
					_ = helper.CreateTestVlanWithVxLan(t, projectId.GetId(), 3979, projName)
					vrf := helper.CreateTestVrfs(t, projectId.GetId(), projName, 3979)
					if vrf.GetId() != "" {
						root.SetArgs([]string{subCommand, "get", "-v", vrf.GetId()})
						out := helper.ExecuteAndCaptureOutput(t, root)

						if !strings.Contains(string(out[:]), vrf.GetId()) &&
							!strings.Contains(string(out[:]), projName) {
							t.Error("expected output should include " + vrf.GetId() + ", in the out string ")
						}
					}
				}
			},
		},
		{
			name: "vrf-ips-getById-test",
			fields: fields{
				MainCmd:  vrf.NewClient(rootClient, outputPkg.Outputer(&outputPkg.Standard{})).NewCommand(),
				Outputer: outputPkg.Outputer(&outputPkg.Standard{}),
			},
			want: &cobra.Command{},
			cmdFunc: func(t *testing.T, c *cobra.Command) {
				root := c.Root()

				projName := "metal-cli-" + randName + "-vrf-ips-get-test"
				project := helper.CreateTestProject(t, projName)

				if project.GetId() != "" {
					_ = helper.CreateTestVlanWithVxLan(t, project.GetId(), 3978, projName)
					vrf := helper.CreateTestVrfs(t, project.GetId(), projName, 3978)
					if vrf.GetId() != "" {

						ipRequest := helper.CreateTestVrfIpRequest(t, project.GetId(), vrf.GetId())
						root.SetArgs([]string{subCommand, "ips", "-v", vrf.GetId(), "-i", ipRequest.IPReservation.GetId()})
						out := helper.ExecuteAndCaptureOutput(t, root)

						if !strings.Contains(string(out[:]), ipRequest.IPReservation.GetId()) &&
							!strings.Contains(string(out[:]), projName) {
							t.Error("expected output should include " + ipRequest.IPReservation.GetId() + ", in the out string ")
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
