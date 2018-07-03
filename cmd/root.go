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

package cmd

import (
	"fmt"
	"io"
	"os"
	"runtime"

	"github.com/packethost/packngo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// PacknGo client
	PacknGo     packngo.Client
	cfgFile     string
	isJSON      bool
	isYaml      bool
	packetToken string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:              "packet",
	Short:            "Command line interface for Packet Host",
	Long:             `Command line interface for Packet Host`,
	PersistentPreRun: packetConnect,
}

func packetConnect(cmd *cobra.Command, args []string) {
	client, err := packngo.NewClientWithBaseURL("Packet CLI", packetToken, nil, "https://api.packet.net/")
	if err != nil {
		fmt.Println("Client error:", err)
		return
	}

	PacknGo = *client
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", fmt.Sprint(userHomeDir(), "/.packet-cli.json"), "Path to JSON or YAML configuration file")
	rootCmd.Version = VERSION
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	}

	viper.AutomaticEnv()
	var r io.Reader
	r, _ = os.Open(cfgFile)

	viper.ReadConfig(r)

	if viper.GetString("token") != "" {
		packetToken = viper.GetString("token")
	} else {
		packetToken = os.Getenv("PACKET_TOKEN")
	}
	if packetToken == "" {
		fmt.Println("Packet authentication token not provided. Please either set the 'PACKET_TOKEN' environment variable or create a JSON configuration file.")
		os.Exit(1)
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
