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

	metal "github.com/equinix-labs/metal-go/metal/v1"
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
			token := string(b)
			fmt.Println(token)

			c.Servicer.SetToken(token)
			fmt.Print("Set token is done\n")

			metalClient := c.Servicer.API(cmd)
			userRequest := metalClient.UsersApi.FindCurrentUser(context.Background())
			user, _, err := metalClient.UsersApi.FindCurrentUserExecute(userRequest)
			if err != nil {
				fmt.Printf("GET : Failed to get current user")
				return err
			}
			organization := user.AdditionalProperties["default_organization_id"]

			fmt.Printf("Organization ID [%s]: ", organization)

			userOrg := ""
			fmt.Scanln(&userOrg)
			if userOrg == "" {
				userOrg = fmt.Sprintf("value: %v", organization)
			}

			project := user.AdditionalProperties["default_project_id"]
			fmt.Printf("Default Project ID [%s]: ", project)

			userProj := ""
			fmt.Scanln(&userProj)
			if userProj == "" {
				userProj = fmt.Sprintf("value: %v", organization)
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
	API(*cobra.Command) *metal.APIClient
	SetToken(string)
	DefaultConfig(bool) string
}
