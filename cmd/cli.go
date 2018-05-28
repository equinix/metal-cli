package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/packethost/packngo"
)

//Cli struct
type Cli struct {
	Client  *packngo.Client
	MainCmd *cobra.Command
}

//NewCli asdfaf
func NewCli() *Cli {
	var err error
	cli := &Cli{}
	cli.Client, err = packngo.NewClientWithBaseURL("Packet CLI", os.Getenv("PACKET_TOKEN"), nil, "https://api.packet.net/")
	if err != nil {
		fmt.Println("Client error:", err)
		return nil
	}
	rootCmd.Execute()
	rootCmd.DisableSuggestions = false
	cli.MainCmd = rootCmd
	return cli
}
