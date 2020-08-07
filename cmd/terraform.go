// Copyright © 2020 Packet, an Equinix Company <info@packet.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"errors"

	"github.com/hashicorp/terraform/registry"
	"github.com/hashicorp/terraform/registry/regsrc"

	"github.com/spf13/cobra"
)

// projectCmd represents the project command
var terraformCmd = &cobra.Command{
	Use: "terraform",
	// Aliases: []string{"registry"},
	Short: "Terraform operations",
	Long:  `Terraform operations: create, delete, update and get`,
}

func init() {
	createTerraformCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the project")
	retrieveTerraformCmd.Flags().StringVarP(&terraformName, "name", "n", "", "Name of the terraform module")

	_ = createTerraformCmd.MarkFlagRequired("name")
	terraformCmd.AddCommand(createTerraformCmd, retrieveTerraformCmd) //  deleteTerraformCmd, updateTerraformCmd

	rootCmd.AddCommand(terraformCmd)
}

// projectCreateCmd represents the projectCreate command
var createTerraformCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a terraform module project",
	Long: `Example:

packet terraform create --name [project_name]
  
  `,
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.New("not supported")
	},
}

var terraformName string

// retrieveTerraformCmd represents the retrieveTerraform command
var retrieveTerraformCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieves all available terraform modules or a single terraform modules",
	Long: `Example:

Retrieve all projects:
packet terraform get
  
Retrieve a specific project:
packet terraform get -n [project_name]
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if terraformName != "" {
			reg := registry.NewClient(nil, nil)
			modQuery := &regsrc.Module{RawNamespace: "displague", RawProvider: "packet"}
			resp, err := reg.ModuleVersions(modQuery)
			if err != nil {
				return errors.New("could not get module versions")
			}

			data := make([][]string, len(resp.Modules))
			for n, mod := range resp.Modules {
				for _, v := range mod.Versions {
					data[n] = []string{mod.Source, v.Version}
				}
			}
			header := []string{"ID", "Name", "Created"}
			return output(resp, header, &data)
		}

		return errors.New("list not supported")
	},
}
