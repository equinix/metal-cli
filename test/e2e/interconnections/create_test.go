package interconnections

import (
	"context"
	"fmt"
	"strings"
	"testing"

	root "github.com/equinix/metal-cli/internal/cli"
	"github.com/equinix/metal-cli/internal/interconnections"
	outputPkg "github.com/equinix/metal-cli/internal/outputs"
	"github.com/equinix/metal-cli/test/helper"

	metal "github.com/equinix/equinix-sdk-go/services/metalv1"
	"github.com/spf13/cobra"
)

func TestInterconnections_Create(t *testing.T) {
	subCommand := "interconnections"
	rootClient := root.NewClient(helper.ConsumerToken, helper.URL, helper.Version)
	randomString := helper.GenerateRandomString(5)

	project := helper.CreateTestProject(t, "metal-cli-interconnections-create-"+randomString)
	vlan := helper.CreateTestVLAN(t, project.GetId())

	apiClient := helper.TestClient()

	tests := []struct {
		name    string
		cmd     *cobra.Command
		want    *cobra.Command
		cmdFunc func(*testing.T, *cobra.Command)
	}{
		{
			name: "create shared vlan interconnection",
			cmd:  interconnections.NewClient(rootClient, outputPkg.Outputer(&outputPkg.Standard{})).NewCommand(),
			want: &cobra.Command{},
			cmdFunc: func(t *testing.T, c *cobra.Command) {
				root := c.Root()
				connName := "conn-1-" + randomString

				root.SetArgs([]string{subCommand, "create", "-p", project.GetId(), "--vlan", fmt.Sprintf("%d", vlan.GetVxlan()), "-n", connName, "-m", vlan.GetMetroCode(), "-r", "primary", "-t", "shared", "-T", "a_side", "-s", "50000000"})

				out := helper.ExecuteAndCaptureOutput(t, root)

				// Need to find the current user's default org
				// as orgId is not populated in project. Its created using default org
				user, _, err := apiClient.UsersApi.
					FindCurrentUser(context.Background()).
					Include([]string{"default_organization_id"}).
					Execute()
				if err != nil {
					t.Fatal(err)
				}

				conns, _, err := apiClient.InterconnectionsApi.
					OrganizationListInterconnections(context.Background(), user.GetDefaultOrganizationId()).
					Execute()
				if err != nil {
					t.Fatal(err)
				}
				if len(conns.GetInterconnections()) < 1 {
					t.Fatal("Interconnections Not Found. Failed to create Interconnections")
				}

				var conn *metal.Interconnection
				for index, c := range conns.GetInterconnections() {
					if c.GetName() == connName {
						conn = &conns.GetInterconnections()[index]
						break
					}
				}

				t.Cleanup(func() {
					helper.CleanupInterconnection(t, conn.GetId())
				})

				assertInterconnectionsCmdOutput(t, string(out[:]), conn)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rootCmd := rootClient.NewCommand()
			rootCmd.AddCommand(tt.cmd)
			tt.cmdFunc(t, tt.cmd)
		})
	}
}

func assertInterconnectionsCmdOutput(t *testing.T, out string, conn *metal.Interconnection) {
	if !strings.Contains(out, conn.GetId()) {
		t.Errorf("cmd output should contain ID of the Interconnection: [%s] \n output:\n%s", conn.GetId(), out)
	}

	if !strings.Contains(out, string(conn.GetType())) {
		t.Errorf("cmd output should contain type of interconection: [%s] \n output:\n%s", conn.GetType(), out)
	}

	if !strings.Contains(out, conn.GetName()) {
		t.Errorf("cmd output should contain name of Interconnection: [%s] \n output:\n%s", conn.GetName(), out)
	}

	if !strings.Contains(out, conn.GetCreatedAt().String()) {
		t.Errorf("cmd output should contain creation time of the Interconnection, expected time:%s, output:\n%s", conn.GetCreatedAt().String(), out)
	}
}
