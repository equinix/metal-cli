package ipstest

import (
	"io"
	"os"
	"strings"
	"testing"
	"time"

	root "github.com/equinix/metal-cli/internal/cli"
	"github.com/equinix/metal-cli/internal/ips"
	outputPkg "github.com/equinix/metal-cli/internal/outputs"
	"github.com/equinix/metal-cli/test/helper"
	"github.com/spf13/cobra"
)

func TestCli_Vlan_Create(t *testing.T) {
	var projectId string
	var err error
	subCommand := "ip"
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
			name: "Request_NewIP",
			fields: fields{
				MainCmd:  ips.NewClient(rootClient, outputPkg.Outputer(&outputPkg.Standard{})).NewCommand(),
				Outputer: outputPkg.Outputer(&outputPkg.Standard{}),
			},
			want: &cobra.Command{},
			cmdFunc: func(t *testing.T, c *cobra.Command) {
				if true {
					t.Skip("Skipping temporarily for now")
				}
				root := c.Root()
				projectName := "metal-cli-ips-get-pro" + helper.GenerateUUID()
				projectId, err = helper.CreateTestProject(t, projectName)
				if err != nil {
					t.Error(err)
				}
				time.Sleep(10 * time.Second)
				if len(projectId) != 0 {
					root.SetArgs([]string{subCommand, "request", "-p", projectId, "-t", "public_ipv4", "-m", "da", "-q", "4"})
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
					if !strings.Contains(string(out[:]), "ID") &&
						!strings.Contains(string(out[:]), "PUBLIC") &&
						!strings.Contains(string(out[:]), "true") {
						t.Error("expected output should include ID, PUBLIC and true strings in the out string")
					}
				}
				err = helper.CleanTestProject(t, projectId)
				if err != nil {
					t.Error(err)
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
