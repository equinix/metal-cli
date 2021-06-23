// Copyright © 2018 Jasmin Gacic <jasmin@stackpointcloud.com>
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

package cli

import (
	"fmt"
	"log"
	"os"
	"path"
	"runtime"

	"github.com/packethost/packngo"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	outputPkg "github.com/equinix/metal-cli/internal/outputs"
)

type Client struct {
	// apiClient client
	apiClient *packngo.Client

	includes      *[]string // nolint:unused
	excludes      *[]string // nolint:unused
	search        string
	cfgFile       string
	isJSON        bool
	isYaml        bool
	metalToken    string
	consumerToken string
	apiURL        string
	Version       string
	rootCmd       *cobra.Command
}

func NewClient(consumerToken, apiURL, Version string) *Client {
	return &Client{
		consumerToken: consumerToken,
		apiURL:        apiURL,
		Version:       Version,
	}
}

type ResponseModifier interface {
	ListOptions(defaultIncludes, defaultExcludes []string) *packngo.ListOptions
}

func (c *Client) apiConnect() error {
	if c.metalToken == "" {
		return fmt.Errorf("Equinix Metal authentication token not provided. Please set the 'METAL_AUTH_TOKEN' environment variable or create a JSON or YAML configuration file.")
	}
	client, err := packngo.NewClientWithBaseURL(c.consumerToken, c.metalToken, nil, c.apiURL)
	if err != nil {
		return errors.Wrap(err, "Could not create Client")
	}
	client.UserAgent = fmt.Sprintf("metal-cli/%s %s", c.Version, client.UserAgent)
	c.apiClient = client
	return nil
}

func (c *Client) API() *packngo.Client {
	if c.apiClient == nil {
		err := c.apiConnect()
		if err != nil {
			log.Fatal(err)
		}
	}
	return c.apiClient
}

func (c *Client) Format() outputPkg.Format {
	format := outputPkg.FormatText

	// TODO(displague) remove isJSON and isYaml globals
	switch {
	case c.isJSON:
		format = outputPkg.FormatJSON
	case c.isYaml:
		format = outputPkg.FormatYAML
	}
	return format
}

func (c *Client) NewCommand() *cobra.Command {
	// rootCmd represents the base command when called without any subcommands
	var rootCmd = &cobra.Command{
		Use:               "metal",
		Short:             "Command line interface for Equinix Metal",
		Long:              `Command line interface for Equinix Metal`,
		DisableAutoGenTag: true,
	}
	rootCmd.PersistentFlags().StringVar(&c.cfgFile, "config", "", "Path to JSON or YAML configuration file")

	rootCmd.PersistentFlags().BoolVarP(&c.isJSON, "json", "j", false, "JSON output")
	rootCmd.PersistentFlags().BoolVarP(&c.isYaml, "yaml", "y", false, "YAML output")

	c.includes = rootCmd.PersistentFlags().StringSlice("include", nil, "Comma seperated Href references to expand in results, may be dotted three levels deep")
	c.excludes = rootCmd.PersistentFlags().StringSlice("exclude", nil, "Comma seperated Href references to collapse in results, may be dotted three levels deep")
	rootCmd.PersistentFlags().StringVar(&c.search, "search", "", "Search keyword for use in 'get' actions. Search is not supported by all resources.")

	rootCmd.Version = c.Version
	c.rootCmd = rootCmd
	return c.rootCmd
}

// ListOptions creates a packngo.ListOptions using the includes and excludes persistent
// flags. When not defined, the defaults given will be supplied.
func (c *Client) ListOptions(defaultIncludes, defaultExcludes []string) *packngo.ListOptions {
	listOptions := &packngo.ListOptions{
		Includes: defaultIncludes,
		Excludes: defaultExcludes,
	}
	if c.rootCmd.Flags().Changed("include") {
		listOptions.Includes = *c.includes
	}
	if c.rootCmd.Flags().Changed("exclude") {
		listOptions.Excludes = *c.excludes
	}
	if c.rootCmd.Flags().Changed("search") {
		listOptions.Search = c.search
	}

	return listOptions
}

// initConfig reads in config file and ENV variables if set.
func (c *Client) Init() {
	v := viper.New()
	if c.cfgFile != "" {
		// Use config file from the flag.
		v.SetConfigFile(c.cfgFile)
	} else {
		configDir := path.Join(userHomeDir(), "/.config/equinix")
		v.SetConfigName("metal")
		v.AddConfigPath(configDir)
	}

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			panic(fmt.Errorf("Could not read config: %s", err))
		}
	}

	v.SetEnvPrefix("METAL")
	v.AutomaticEnv()
	c.metalToken = v.GetString("token")
	envToken := v.GetString("auth_token")
	// TODO: are we ok with this being configured by file too? yes?
	if envToken != "" {
		c.metalToken = envToken
	}
}

func userHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}
