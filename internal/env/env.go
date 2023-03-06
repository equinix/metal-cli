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
package env

import (
	"fmt"

	"github.com/spf13/cobra"
)

type Tokener interface {
	Token() string
}

type Client struct {
	tokener        Tokener
	apiTokenEnvVar string
}

func NewClient(t Tokener, apiTokenEnvVar string) *Client {
	return &Client{
		tokener:        t,
		apiTokenEnvVar: apiTokenEnvVar,
	}
}

func (c *Client) NewCommand() *cobra.Command {
	// envCmd represents a command that, when run, generates a
	// set of environment variables, for use in shell environments
	// v := c.tokener.Config()
	// projectId := v.GetString("project-id")
	envCmd := &cobra.Command{
		Use:   `env [-p <project_id>]`,
		Short: "Prints or generates environment variables.",
		Long:  "Prints or generates environment variables. Currently emitted variables: METAL_AUTH_TOKEN, METAL_ORGANIZATION_ID, METAL_PROJECT_ID, METAL_CONFIG. Use the --project-id flag to set the METAL_PROJECT_ID variable. Use the --organization-id flag to set the METAL_ORGANIZATION_ID variable.",
		Example: `  # Print the current environment variables:
  metal env
  
  # Print the current environment variables in Terraform format:
  metal env --output terraform
  
    # Load environment variables in Bash, Zsh:
  source <(metal env)
  
  # Load environment variables in Bash 3.2.x:
  eval "$(metal env)"
  
  # Load environment variables in Fish:
  metal env | source`,

		DisableFlagsInUseLine: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			var formatter func(token, orgID, projID, conPath string) map[string]string

			organizationID, _ := cmd.Flags().GetString("organization-id")
			projectID, _ := cmd.Flags().GetString("project-id")
			config, _ := cmd.Flags().GetString("config")
			output, _ := cmd.Flags().GetString("output")
			export, _ := cmd.Flags().GetBool("export")
			prefix := ""

			switch output {
			case "sh":
				formatter = shellEnvWithClientToken(c)
			case "terraform":
				formatter = terraformEnv
			case "capp":
				formatter = cappEnv
			default:
				return fmt.Errorf("Unknown env output format %q", output)
			}

			if export {
				prefix = "export "
			}

			for k, v := range formatter(c.tokener.Token(), organizationID, projectID, config) {
				fmt.Printf("%s%s=%s\n", prefix, k, v)
			}
			return nil
		},
	}

	envCmd.PersistentFlags().StringP("output", "o", "sh", "Output format for environment variables (*sh, terraform, capp).")

	envCmd.Flags().StringP("project-id", "p", "", "A project UUID to set as an environment variable.")

	envCmd.Flags().StringP("organization-id", "O", "", "A organization UUID to set as an environment variable.")

	envCmd.Flags().Bool("export", false, "Export the environment variables.")

	return envCmd
}

func terraformEnv(token, orgID, projID, conPath string) map[string]string {
	return map[string]string{
		"TF_VAR_metal_auth_token":      token,
		"TF_VAR_metal_organization_id": orgID,
		"TF_VAR_metal_project_id":      projID,
		"TF_VAR_metal_config":          conPath,
	}
}

func cappEnv(token, orgID, projID, conPath string) map[string]string {
	return map[string]string{
		"PACKET_API_KEY":  token,
		"ORGANIZATION_ID": orgID,
		"PROJECT_ID":      projID,
	}
}

func shellEnvWithClientToken(c *Client) func(token, orgID, projID, conPath string) map[string]string {
	return func(token, orgID, projID, conPath string) map[string]string {
		return map[string]string{
			c.apiTokenEnvVar:        token,
			"METAL_ORGANIZATION_ID": orgID,
			"METAL_PROJECT_ID":      projID,
			"METAL_CONFIG":          conPath,
		}
	}
}
