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
	"net/http"
	"os"
	"path"
	"runtime"
	"strings"

	metal "github.com/equinix/equinix-sdk-go/services/metalv1"
	"github.com/packethost/packngo"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	outputPkg "github.com/equinix/metal-cli/internal/outputs"
)

const (
	envPrefix                  = "METAL"
	configFileWithoutExtension = "metal"
	debugVar                   = "PACKNGO_DEBUG"
)

type Client struct {
	// apiClient client
	apiClient      *packngo.Client
	metalApiClient *metal.APIClient

	includes      *[]string // nolint:unused
	excludes      *[]string // nolint:unused
	filters       *[]string
	search        string
	sortBy        string
	sortDir       string
	cfgFile       string
	outputFormat  string
	metalToken    string
	consumerToken string
	apiURL        string
	Version       string
	rootCmd       *cobra.Command
	viper         *viper.Viper
}

type headerTransport struct {
	header http.Header
}

func (t *headerTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	for key, values := range t.header {
		for _, value := range values {
			r.Header.Add(key, value)
		}
	}
	return http.DefaultTransport.RoundTrip(r)
}

func NewClient(consumerToken, apiURL, Version string) *Client {
	return &Client{
		consumerToken: consumerToken,
		apiURL:        apiURL,
		Version:       Version,
	}
}

// This function provides backwards compatibility for the packngo
// debug environment variable while allowing us to introduce a new
// debug variable in the future that is not tied to packngo
func checkEnvForDebug() bool {
	return os.Getenv(debugVar) != ""
}

func (c *Client) apiConnect(httpClient *http.Client) error {
	client, err := packngo.NewClientWithBaseURL(c.consumerToken, c.metalToken, httpClient, c.apiURL)
	if err != nil {
		return fmt.Errorf("could not create client: %w", err)
	}
	client.UserAgent = fmt.Sprintf("metal-cli/%s %s", c.Version, client.UserAgent)
	c.apiClient = client
	return nil
}

func (c *Client) metalApiConnect(httpClient *http.Client) error {
	configuration := metal.NewConfiguration()
	configuration.Debug = checkEnvForDebug()
	configuration.AddDefaultHeader("X-Auth-Token", c.Token())
	configuration.UserAgent = fmt.Sprintf("metal-cli/%s %s", c.Version, configuration.UserAgent)
	configuration.Servers = metal.ServerConfigurations{
		metal.ServerConfiguration{
			URL: c.apiURL,
		},
	}
	metalgoClient := metal.NewAPIClient(configuration)
	c.metalApiClient = metalgoClient
	return nil
}

func (c *Client) Config(cmd *cobra.Command) *viper.Viper {
	if c.viper == nil {
		v := viper.New()
		v.AutomaticEnv()

		replacer := strings.NewReplacer("-", "_", ".", "_")
		v.SetEnvKeyReplacer(replacer)

		if c.cfgFile != "" {
			// Use config file from the flag.
			v.SetConfigFile(c.cfgFile)
		} else if envCfgFile := os.Getenv(envPrefix + "_CONFIG"); envCfgFile != "" {
			// METAL_CONFIG must be checked since cfgFile does not benefit from AutomaticEnv
			v.SetConfigFile(envCfgFile)
		} else {
			configDir := defaultConfigPath()
			v.SetConfigName(configFileWithoutExtension)
			v.AddConfigPath(configDir)
		}
		if err := v.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
				panic(fmt.Errorf("could not read config: %s", err))
			}
		}
		c.cfgFile = v.ConfigFileUsed()

		v.SetEnvPrefix(envPrefix)
		c.viper = v
	}

	bindFlags(cmd, c.viper)

	flagToken := cmd.Flag("token").Value.String()
	envToken := cmd.Flag("auth-token").Value.String()
	// TODO: are we ok with this being configured by file too? yes?
	// TODO: let cli arg take higher priority
	c.metalToken = flagToken
	if envToken != "" {
		c.metalToken = envToken
	}
	return c.viper
}

// Credit to https://carolynvanslyck.com/blog/2020/08/sting-of-the-viper/
func bindFlags(cmd *cobra.Command, v *viper.Viper) {
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		// Environment variables can't have dashes in them, so bind them to their equivalent
		// keys with underscores, e.g. --favorite-color to STING_FAVORITE_COLOR
		if strings.Contains(f.Name, "-") {
			envVarSuffix := strings.ToUpper(strings.ReplaceAll(f.Name, "-", "_"))
			_ = v.BindEnv(f.Name, fmt.Sprintf("%s_%s", envPrefix, envVarSuffix))
		}

		// Apply the viper config value to the flag when the flag is not set and viper has a value
		if !f.Changed && v.IsSet(f.Name) {
			val := v.Get(f.Name)
			_ = cmd.Flags().Set(f.Name, fmt.Sprintf("%v", val))
		}
	})
}

func (c *Client) API(cmd *cobra.Command) *packngo.Client {
	if c.metalToken == "" {
		log.Fatal("Equinix Metal authentication token not provided. Please set the 'METAL_AUTH_TOKEN' environment variable or create a configuration file using 'metal init'.")
	}

	if c.apiClient == nil {
		httpClient := &http.Client{
			Transport: &headerTransport{
				header: getAdditionalHeaders(cmd),
			},
		}

		err := c.apiConnect(httpClient)
		if err != nil {
			log.Fatal(err)
		}
	}
	return c.apiClient
}

func (c *Client) MetalAPI(cmd *cobra.Command) *metal.APIClient {
	if c.metalToken == "" {
		log.Fatal("Equinix Metal authentication token not provided. Please set the 'METAL_AUTH_TOKEN' environment variable or create a configuration file using 'metal init'.")
	}

	if c.metalApiClient == nil {
		httpClient := &http.Client{
			Transport: &headerTransport{
				header: getAdditionalHeaders(cmd),
			},
		}

		err := c.metalApiConnect(httpClient)
		if err != nil {
			log.Fatal(err)
		}
	}
	return c.metalApiClient
}

func (c *Client) Token() string {
	return c.metalToken
}

func (c *Client) SetToken(token string) {
	c.metalToken = token
}

func (c *Client) Format() outputPkg.Format {
	format := outputPkg.FormatTable

	switch f := outputPkg.Format(c.outputFormat); f {
	case "":
		break
	case outputPkg.FormatTable,
		outputPkg.FormatTerraform,
		outputPkg.FormatCrossplane,
		outputPkg.FormatJSON,
		outputPkg.FormatYAML:
		format = f
	default:
		log.Printf("unknown format: %q. Using default.", f)
	}
	return format
}

func (c *Client) NewCommand() *cobra.Command {
	// rootCmd represents the base command when called without any subcommands
	rootCmd := &cobra.Command{
		Use:               "metal",
		Short:             "Command line interface for Equinix Metal",
		Long:              `Command line interface for Equinix Metal`,
		DisableAutoGenTag: true,

		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			c.Config(cmd)
		},
	}
	rootCmd.PersistentFlags().String("token", "", "Metal API Token (METAL_AUTH_TOKEN)")
	rootCmd.PersistentFlags().String("auth-token", "", "Metal API Token (Alias)")
	rootCmd.PersistentFlags().StringSlice("http-header", nil, "Headers to add to requests (in format key=value)")
	authtoken := rootCmd.PersistentFlags().Lookup("auth-token")
	authtoken.Hidden = true
	rootCmd.PersistentFlags().StringVar(&c.cfgFile, "config", c.cfgFile, "Path to JSON or YAML configuration file (METAL_CONFIG)")
	rootCmd.PersistentFlags().StringVarP(&c.outputFormat, "output", "o", "", "Output format (*table, json, yaml). env output formats are (*sh, terraform, capp).")
	c.includes = rootCmd.PersistentFlags().StringSlice("include", nil, "Comma separated Href references to expand in results, may be dotted three levels deep")
	c.excludes = rootCmd.PersistentFlags().StringSlice("exclude", nil, "Comma separated Href references to collapse in results, may be dotted three levels deep")
	c.filters = rootCmd.PersistentFlags().StringArray("filter", nil, "Filter 'get' actions with name value pairs. Filter is not supported by all resources and is implemented as request query parameters.")
	rootCmd.PersistentFlags().StringVar(&c.search, "search", "", "Search keyword for use in 'get' actions. Search is not supported by all resources.")
	rootCmd.PersistentFlags().StringVar(&c.sortBy, "sort-by", "", "Sort fields for use in 'get' actions. Sort is not supported by all resources.")
	rootCmd.PersistentFlags().StringVar(&c.sortDir, "sort-dir", "", "Sort field direction for use in 'get' actions. Sort is not supported by all resources.")

	rootCmd.Version = c.Version
	c.rootCmd = rootCmd
	return c.rootCmd
}

func (c *Client) Includes(defaultIncludes []string) (incl []string) {
	var inc []string

	inc = defaultIncludes

	if c.rootCmd.Flags().Changed("include") {
		inc = *c.includes
	}

	return inc
}

func (c *Client) Excludes(defaultExcludes []string) (excl []string) {
	var exc []string

	exc = defaultExcludes

	if c.rootCmd.Flags().Changed("exclude") {
		exc = *c.excludes
	}

	return exc
}

func (c *Client) Search() (sea string) {
	var search string

	if c.rootCmd.Flags().Changed("search") {
		search = c.search
	}

	return search
}

func (c *Client) SortBy() (sBy string) {
	var sortBy string

	if c.rootCmd.Flags().Changed("sort-by") {
		sortBy = c.sortBy
	}

	return sortBy
}

func (c *Client) SortDirection() (sDir string) {
	var sortDir string

	if c.rootCmd.Flags().Changed("sort-dir") {
		sortDir = c.sortDir
	}

	return sortDir
}

func (c *Client) Filters() map[string]string {
	mapFilt := make(map[string]string)
	if c.rootCmd.Flags().Changed("filter") {
		for _, kv := range *c.filters {
			var k, v string
			tokens := strings.SplitN(kv, "=", 2)
			k = strings.TrimSpace(tokens[0])
			if len(tokens) != 1 {
				v = strings.TrimSpace(tokens[1])
			}
			mapFilt[k] = v
		}
	}
	return mapFilt
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
	if c.rootCmd.Flags().Changed("filter") {
		for _, kv := range *c.filters {
			var k, v string
			tokens := strings.SplitN(kv, "=", 2)
			k = strings.TrimSpace(tokens[0])
			if len(tokens) != 1 {
				v = strings.TrimSpace(tokens[1])
			}
			listOptions = listOptions.Filter(k, v)
		}
	}
	if c.rootCmd.Flags().Changed("search") {
		listOptions.Search = c.search
	}
	if c.rootCmd.Flags().Changed("sort-by") {
		listOptions.SortBy = c.sortBy
	}
	if c.rootCmd.Flags().Changed("sort-dir") {
		listOptions.SortDirection = packngo.ListSortDirection(c.sortDir)
	}

	return listOptions
}

// initConfig reads in config file and ENV variables if set.
func (c *Client) Init(cmd *cobra.Command) {
	// v := c.Config(cmd)
	c.Config(cmd)
	// c.metalToken = v.GetString("token")
	// envToken := v.GetString("auth_token")
	// TODO: are we ok with this being configured by file too? yes?
	// if envToken != "" {
	//		c.metalToken = envToken
	//	}
}

func defaultConfigPath() string {
	return path.Join(userHomeDir(), "/.config/equinix")
}

func (c *Client) DefaultConfig(withExtension bool) string {
	dir := defaultConfigPath()
	config := path.Join(dir, configFileWithoutExtension)
	if withExtension {
		config = config + ".yaml"
	}
	return config
}

func getAdditionalHeaders(cmd *cobra.Command) http.Header {
	header := make(http.Header)

	v, err := cmd.Flags().GetStringSlice("http-header")
	if err != nil {
		return header
	}

	for _, headerStr := range v {
		s := strings.SplitN(headerStr, "=", 2)
		if len(s) != 2 {
			// Ignore any malformed header strings.
			continue
		}

		for _, value := range strings.Split(s[1], ",") {
			header.Add(s[0], value)
		}
	}

	return header
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
