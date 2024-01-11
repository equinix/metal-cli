package projecttest

import (
	"fmt"
	"strings"
	"testing"

	root "github.com/equinix/metal-cli/internal/cli"
	outputPkg "github.com/equinix/metal-cli/internal/outputs"
	"github.com/equinix/metal-cli/internal/projects"
	"github.com/equinix/metal-cli/test/helper"
	"github.com/spf13/cobra"
)

func TestCli_Project_Tests(t *testing.T) {
	subCommand := "project"
	rootClient := root.NewClient(helper.ConsumerToken, helper.URL, helper.Version)
	randName := helper.GenerateRandomString(32)

	type fields struct {
		MainCmd  *cobra.Command
		Outputer outputPkg.Outputer
	}

	tests := []struct {
		name    string
		fields  fields
		want    *cobra.Command
		cmdFunc func(*testing.T, *cobra.Command, *cobra.Command, string)
	}{
		{
			name: "project-create-test",
			fields: fields{
				MainCmd:  projects.NewClient(rootClient, outputPkg.Outputer(&outputPkg.Standard{})).NewCommand(),
				Outputer: outputPkg.Outputer(&outputPkg.Standard{}),
			},
			want: &cobra.Command{},
			cmdFunc: func(t *testing.T, c *cobra.Command, rootCmd *cobra.Command, projectID string) {
				root := c.Root()
				projName := "metal-cli-" + randName + "-project-create-test"
				org := helper.CreateTestOrganization(t, projName)

				root.SetArgs([]string{subCommand, "create", "-O", org.GetId(), "-n", projName})

				out := helper.ExecuteAndCaptureOutput(t, root)

				if !strings.Contains(string(out[:]), projName) {
					t.Error("expected output should include " + projName + ", in the out string ")
				}
			},
		},

		{
			name: "project-update-test",
			fields: fields{
				MainCmd:  projects.NewClient(rootClient, outputPkg.Outputer(&outputPkg.Standard{})).NewCommand(),
				Outputer: outputPkg.Outputer(&outputPkg.Standard{}),
			},
			want: &cobra.Command{},
			cmdFunc: func(t *testing.T, c *cobra.Command, rootCmd *cobra.Command, projectID string) {
				root := c.Root()

				projName := "metal-cli-" + randName + "-project-update-test"

				project := helper.CreateTestProject(t, projName)
				updateProjName := projName + "-123"

				root.SetArgs([]string{subCommand, "update", "-i", project.GetId(), "-n", updateProjName})

				out := helper.ExecuteAndCaptureOutput(t, root)

				if !strings.Contains(string(out[:]), updateProjName) &&
					!strings.Contains(string(out[:]), project.GetId()) {
					t.Error("expected output should include " + updateProjName + project.GetId() + ", in the out string ")
				}
			},
		},

		{
			name: "project-get-test",
			fields: fields{
				MainCmd:  projects.NewClient(rootClient, outputPkg.Outputer(&outputPkg.Standard{})).NewCommand(),
				Outputer: outputPkg.Outputer(&outputPkg.Standard{}),
			},
			want: &cobra.Command{},
			cmdFunc: func(t *testing.T, c *cobra.Command, rootCmd *cobra.Command, projectID string) {
				root := c.Root()
				projName := "metal-cli-" + randName + "-project-get-test"

				project := helper.CreateTestProject(t, projName)

				root.SetArgs([]string{subCommand, "get"})

				out := helper.ExecuteAndCaptureOutput(t, root)

				if !strings.Contains(string(out[:]), projName) &&
					!strings.Contains(string(out[:]), project.GetId()) {
					t.Error("expected output should include " + projName + project.GetId() + ", in the out string ")
				}
			},
		},

		{
			name: "project-get-id-test",
			fields: fields{
				MainCmd:  projects.NewClient(rootClient, outputPkg.Outputer(&outputPkg.Standard{})).NewCommand(),
				Outputer: outputPkg.Outputer(&outputPkg.Standard{}),
			},
			want: &cobra.Command{},
			cmdFunc: func(t *testing.T, c *cobra.Command, rootCmd *cobra.Command, projectID string) {
				root := c.Root()
				projName := "metal-cli-" + randName + "-project-get-id-test"

				project := helper.CreateTestProject(t, projName)

				root.SetArgs([]string{subCommand, "get", "-i", project.GetId()})

				out := helper.ExecuteAndCaptureOutput(t, root)

				if !strings.Contains(string(out[:]), projName) &&
					!strings.Contains(string(out[:]), project.GetId()) {
					t.Error("expected output should include " + projName + project.GetId() + ", in the out string ")
				}
			},
		},

		{
			name: "project-delete-test",
			fields: fields{
				MainCmd:  projects.NewClient(rootClient, outputPkg.Outputer(&outputPkg.Standard{})).NewCommand(),
				Outputer: outputPkg.Outputer(&outputPkg.Standard{}),
			},
			want: &cobra.Command{},
			cmdFunc: func(t *testing.T, c *cobra.Command, rootCmd *cobra.Command, projectID string) {
				root := c.Root()

				projName := "metal-cli-" + randName + "-project-delete-test"

				project := helper.CreateTestProject(t, projName)

				root.SetArgs([]string{subCommand, "delete", "-i", project.GetId(), "-f"})

				out := helper.ExecuteAndCaptureOutput(t, root)

				expectedOut := fmt.Sprintf("Project %s successfully deleted.", project.GetId())
				if !strings.Contains(string(out[:]), expectedOut) {
					t.Error(fmt.Errorf("expected output: '%s' but got '%s'", expectedOut, string(out)))
				}
			},
		},

		{
			name: "project-bgpenbale-test",
			fields: fields{
				MainCmd:  projects.NewClient(rootClient, outputPkg.Outputer(&outputPkg.Standard{})).NewCommand(),
				Outputer: outputPkg.Outputer(&outputPkg.Standard{}),
			},
			want: &cobra.Command{},
			cmdFunc: func(t *testing.T, c *cobra.Command, rootCmd *cobra.Command, projectID string) {
				root := c.Root()

				projName := "metal-cli-" + randName + "-project-bgpenable-test"

				project := helper.CreateTestProject(t, projName)
				asn := "65000"
				dtype := "local"

				root.SetArgs([]string{subCommand, "bgp-enable", "--project-id", project.GetId(), "--deployment-type", dtype, "--asn", asn})

				out := helper.ExecuteAndCaptureOutput(t, root)

				if !strings.Contains(string(out[:]), project.GetId()) &&
					!strings.Contains(string(out[:]), asn) &&
					!strings.Contains(string(out[:]), dtype) {
					t.Error("expected output should include " + project.GetId() + "," + asn + " and " + dtype + ", in the out string ")
				}
			},
		},

		{
			name: "project-bgpconfig-test",
			fields: fields{
				MainCmd:  projects.NewClient(rootClient, outputPkg.Outputer(&outputPkg.Standard{})).NewCommand(),
				Outputer: outputPkg.Outputer(&outputPkg.Standard{}),
			},
			want: &cobra.Command{},
			cmdFunc: func(t *testing.T, c *cobra.Command, rootCmd *cobra.Command, projectID string) {
				root := c.Root()

				projName := "metal-cli-" + randName + "-project-bgpconfig-test"

				project := helper.CreateTestProject(t, projName)

				err := helper.CreateTestBgpEnableTest(project.GetId())
				if err != nil {
					t.Error(err)
				}

				asn := "65000"
				dtype := "local"
				status := "enabled"

				root.SetArgs([]string{subCommand, "bgp-config", "--project-id", project.GetId()})

				out := helper.ExecuteAndCaptureOutput(t, root)

				if !strings.Contains(string(out[:]), project.GetId()) &&
					!strings.Contains(string(out[:]), asn) &&
					!strings.Contains(string(out[:]), status) &&
					!strings.Contains(string(out[:]), dtype) {
					t.Error("expected output should include " + project.GetId() + "," + asn + "," + dtype + " and " + status + ", in the out string ")
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rootCmd := rootClient.NewCommand()
			rootCmd.AddCommand(tt.fields.MainCmd)
			tt.cmdFunc(t, tt.fields.MainCmd, rootCmd, "")
		})
	}
}
