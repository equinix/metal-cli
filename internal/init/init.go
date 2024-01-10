/*
Copyright © 2020 Equinix Metal Developers <support@equinixmetal.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package init

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"syscall"

	metal "github.com/equinix/equinix-sdk-go/services/metalv1"
	"github.com/spf13/cobra"
	"golang.org/x/term"
	"sigs.k8s.io/yaml"
)

type Client struct {
	Servicer       Servicer
	UserService    metal.UsersApiService
	ProjectService metal.ProjectsApiService
}

func NewClient(s Servicer) *Client {
	return &Client{
		Servicer: s,
	}
}

type configFormat struct {
	Token          string `json:"token,omitempty"`
	ProjectID      string `json:"project-id,omitempty"`
	OrganizationID string `json:"organization-id,omitempty"`
}

func (c *Client) NewCommand() *cobra.Command {
	// initCmd represents a command that, when run, generates a
	// set of initironment variables, for use in shell initironments
	// v := c.tokener.Config()
	// projectId := v.GetString("project-id")
	initCmd := &cobra.Command{
		Use:   `init`,
		Short: "Create a configuration file.",
		Long:  "Init will prompt for account settings and store the values as defaults in a configuration file that may be shared with other Equinix Metal tools. This file is typically stored in $HOME/.config/equinix/metal.yaml. Any Metal CLI command line argument can be specified in the config file. Be careful not to define options that you do not intend to use as defaults. The configuration file written to can be changed with METAL_CONFIG and --config.",

		Example: `  # Example config:
  --
  token: foo
  project-id: 1857dc19-76a5-4589-a9b6-adb729a7d18b
  organization-id: 253e9cf1-5b3d-41f5-a4fa-839c130c8c1d`,

		DisableFlagsInUseLine: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			cmd.SilenceUsage = true
			config, _ := cmd.Flags().GetString("config")
			if config == "" {
				config = c.Servicer.DefaultConfig(true)
			}

			fmt.Print("Equinix Metal API Tokens can be obtained through the portal at https://console.equinix.com/profile/api-keys\nSee https://metal.equinix.com/developers/docs/accounts/users/ for more details.\n\n")
			fmt.Print("User Token (hidden): ")
			b, err := term.ReadPassword(int(syscall.Stdin))
			if err != nil {
				return err
			}
			fmt.Println()
			token := string(b)
			c.Servicer.SetToken(token)
			metalGoClient := c.Servicer.MetalAPI(cmd)
			c.UserService = *metalGoClient.UsersApi
			c.ProjectService = *metalGoClient.ProjectsApi

			include := []string{} // []string | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. (optional)
			exclude := []string{"devices", "members", "memberships", "invitations", "ssh_keys", "volumes", "backend_transfer_enabled", "updated_at", "customdata", "event_alert_configuration",
				"timezone", "features", "avatar_url", "avatar_thumb_url", "two_factor_auth", "mailing_address", "max_projects", "verification_stage", "emails", "phone_number", "restricted",
				"full_name", "email", "social_accounts", "opt_in_updated_at", "opt_in", "first_name", "last_name", "last_login_at"}
			user, _, err := c.UserService.FindCurrentUser(context.Background()).Include(include).Exclude(exclude).Execute()
			if err != nil {
				return err
			}
			organization := user.GetDefaultOrganizationId()
			project := user.GetDefaultProjectId()
			fmt.Printf("Organization ID [%s]: ", organization)

			userOrg := ""
			fmt.Scanln(&userOrg)
			if userOrg == "" {
				userOrg = organization
			}

			// Choose the first project in the preferred org
			if project == "" {
				project, err = getFirstProjectID(c.ProjectService, userOrg)
				if err != nil {
					return err
				}
			}

			fmt.Printf("Project ID [%s]: ", project)

			userProj := ""
			fmt.Scanln(&userProj)
			if userProj == "" {
				userProj = project
			}

			b, err = formatConfig(userProj, userOrg, token)
			if err != nil {
				return err
			}
			return writeConfig(config, b)
		},
	}

	return initCmd
}

func getFirstProjectID(s metal.ProjectsApiService, userOrg string) (string, error) {
	include := []string{"organization"} // []string | Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects. (optional)
	exclude := []string{"devices", "members", "memberships", "invitations", "ssh_keys", "volumes", "backend_transfer_enabled", "updated_at", "customdata", "event_alert_configuration"}
	resp, err := s.FindProjects(context.Background()).Include(include).Exclude(exclude).ExecuteWithPagination()
	if err != nil {
		return "", err
	}

	projects := resp.Projects
	for _, p := range projects {
		if p.Organization.GetId() == userOrg {
			return p.GetId(), nil
		}
	}

	return "", nil // it's ok to have no projects and no default project
}

func formatConfig(userProj, userOrg, token string) ([]byte, error) {
	f := configFormat{ProjectID: userProj, OrganizationID: userOrg, Token: token}
	b, err := yaml.Marshal(f)
	if err != nil {
		return nil, err
	}
	b = append([]byte("---\n"), b...)
	return b, err
}

func writeConfig(config string, b []byte) error {
	fmt.Fprintf(os.Stderr, "\nWriting %s\n", config)
	dir := filepath.Dir(config)
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return fmt.Errorf("could not make directory %q: %s", dir, err)
	}
	return os.WriteFile(config, b, 0o600)
}

type Servicer interface {
	MetalAPI(*cobra.Command) *metal.APIClient
	SetToken(string)
	DefaultConfig(bool) string
}
