package projecttest

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"

	root "github.com/equinix/metal-cli/internal/cli"
	outputPkg "github.com/equinix/metal-cli/internal/outputs"
	"github.com/equinix/metal-cli/internal/projects"
	"github.com/equinix/metal-cli/test/helper"
	"github.com/spf13/cobra"
)

// setupTestOrganization initializes a test Orge and returns its ID along with a cleanup function.
func setupTestOrganization(t *testing.T, projectName string) (string, func()) {
	orgId, err := helper.CreateTestOrganization(projectName)
	if err != nil {
		t.Fatal(err)
	}

	teardown := func() {
		err := helper.CleanTestOrganization(orgId)
		if err != nil {
			t.Error(err)
		}
	}

	return orgId, teardown
}

// setupTestProject initializes a test project and returns its ID along with a cleanup function.
func setupTestProject(t *testing.T, projectName string) (string, func()) {
	projectId, err := helper.CreateTestProject(t, projectName)
	if err != nil {
		t.Fatal(err)
	}

	teardown := func() {
		err := helper.CleanTestProject(t, projectId)
		if err != nil {
			t.Error(err)
		}
	}

	return projectId, teardown
}

func TestCli_Project_Tests(t *testing.T) {
	subCommand := "project"
	consumerToken := ""
	apiURL := ""
	Version := "metal"
	rootClient := root.NewClient(consumerToken, apiURL, Version)
	randName := helper.GenerateRandomString(5)

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
				orgId, cleanupOrg := setupTestOrganization(t, projName)
				defer cleanupOrg()

				if orgId != "" {
					root.SetArgs([]string{subCommand, "create", "-O", orgId, "-n", projName})
					rescueStdout := os.Stdout
					r, w, _ := os.Pipe()
					os.Stdout = w
					t.Cleanup(func() {
						w.Close()
						os.Stdout = rescueStdout
					})

					if err := root.Execute(); err != nil {
						t.Fatal(err)
					}

					out, _ := io.ReadAll(r)
					if !strings.Contains(string(out[:]), projName) {
						t.Error("expected output should include " + projName + ", in the out string ")
					}
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

				projectId, cleanProject := setupTestProject(t, projName)
				defer cleanProject()
				updateProjName := projName + "-123"

				if projectId != "" {
					root.SetArgs([]string{subCommand, "update", "-i", projectId, "-n", updateProjName})
					rescueStdout := os.Stdout
					r, w, _ := os.Pipe()
					os.Stdout = w
					t.Cleanup(func() {
						w.Close()
						os.Stdout = rescueStdout
					})

					if err := root.Execute(); err != nil {
						t.Fatal(err)
					}

					out, _ := io.ReadAll(r)
					if !strings.Contains(string(out[:]), updateProjName) &&
						!strings.Contains(string(out[:]), projectId) {
						t.Error("expected output should include " + updateProjName + projectId + ", in the out string ")
					}
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

				projectId, cleanProject := setupTestProject(t, projName)
				defer cleanProject()

				if projectId != "" {
					root.SetArgs([]string{subCommand, "get"})
					rescueStdout := os.Stdout
					r, w, _ := os.Pipe()
					os.Stdout = w
					t.Cleanup(func() {
						w.Close()
						os.Stdout = rescueStdout
					})

					if err := root.Execute(); err != nil {
						t.Fatal(err)
					}

					out, _ := io.ReadAll(r)
					if !strings.Contains(string(out[:]), projName) &&
						!strings.Contains(string(out[:]), projectId) {
						t.Error("expected output should include " + projName + projectId + ", in the out string ")
					}
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

				projectId, cleanProject := setupTestProject(t, projName)
				defer cleanProject()

				if projectId != "" {
					root.SetArgs([]string{subCommand, "get", "-i", projectId})
					rescueStdout := os.Stdout
					r, w, _ := os.Pipe()
					os.Stdout = w
					t.Cleanup(func() {
						w.Close()
						os.Stdout = rescueStdout
					})

					if err := root.Execute(); err != nil {
						t.Fatal(err)
					}

					out, _ := io.ReadAll(r)
					if !strings.Contains(string(out[:]), projName) &&
						!strings.Contains(string(out[:]), projectId) {
						t.Error("expected output should include " + projName + projectId + ", in the out string ")
					}
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

				projectId, _ := setupTestProject(t, projName)

				if projectId != "" {
					root.SetArgs([]string{subCommand, "delete", "-i", projectId, "-f"})
					rescueStdout := os.Stdout
					r, w, _ := os.Pipe()
					os.Stdout = w
					t.Cleanup(func() {
						w.Close()
						os.Stdout = rescueStdout
					})

					if err := root.Execute(); err != nil {
						t.Fatal(err)
					}

					out, _ := io.ReadAll(r)
					expectedOut := fmt.Sprintf("Project %s successfully deleted.", projectId)
					if !strings.Contains(string(out[:]), expectedOut) {
						t.Error(fmt.Errorf("expected output: '%s' but got '%s'", expectedOut, string(out)))
					}
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

				projectId, cleanProject := setupTestProject(t, projName)
				defer cleanProject()
				asn := "65000"
				dtype := "local"

				if projectId != "" {
					root.SetArgs([]string{subCommand, "bgp-enable", "--project-id", projectId, "--deployment-type", dtype, "--asn", asn})
					rescueStdout := os.Stdout
					r, w, _ := os.Pipe()
					os.Stdout = w
					t.Cleanup(func() {
						w.Close()
						os.Stdout = rescueStdout
					})

					if err := root.Execute(); err != nil {
						t.Fatal(err)
					}

					out, _ := io.ReadAll(r)
					if !strings.Contains(string(out[:]), projectId) &&
						!strings.Contains(string(out[:]), asn) &&
						!strings.Contains(string(out[:]), dtype) {
						t.Error("expected output should include " + projectId + "," + asn + " and " + dtype + ", in the out string ")
					}
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

				projectId, cleanProject := setupTestProject(t, projName)
				defer cleanProject()

				err := helper.CreateTestBgpEnableTest(projectId)
				if err != nil {
					t.Error(err)
				}

				asn := "65000"
				dtype := "local"
				status := "enabled"

				if projectId != "" {
					root.SetArgs([]string{subCommand, "bgp-config", "--project-id", projectId})
					rescueStdout := os.Stdout
					r, w, _ := os.Pipe()
					os.Stdout = w
					t.Cleanup(func() {
						w.Close()
						os.Stdout = rescueStdout
					})

					if err := root.Execute(); err != nil {
						t.Fatal(err)
					}

					out, _ := io.ReadAll(r)
					if !strings.Contains(string(out[:]), projectId) &&
						!strings.Contains(string(out[:]), asn) &&
						!strings.Contains(string(out[:]), status) &&
						!strings.Contains(string(out[:]), dtype) {
						t.Error("expected output should include " + projectId + "," + asn + "," + dtype + " and " + status + ", in the out string ")
					}
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
