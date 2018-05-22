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

package main

import "github.com/StackPointCloud/packetcli/cmd"

func main() {
	cmd.NewCli()
}

// package main

// import (
// 	"fmt"
// 	"strings"

// 	"github.com/spf13/cobra"
// )

// func main() {
// 	var echoTimes int

// 	var cmdPrint = &cobra.Command{
// 		Use:   "print [string to print]",
// 		Short: "Print anything to the screen",
// 		Long: `print is for printing anything back to the screen.
// For many years people have printed back to the screen.`,
// 		Args: cobra.MinimumNArgs(1),
// 		Run: func(cmd *cobra.Command, args []string) {
// 			fmt.Println("Print: " + strings.Join(args, " "))
// 		},
// 	}

// 	var cmdEcho = &cobra.Command{
// 		Use:   "echo [string to echo]",
// 		Short: "Echo anything to the screen",
// 		Long: `echo is for echoing anything back.
// Echo works a lot like print, except it has a child command.`,
// 		Args: cobra.MinimumNArgs(1),
// 		Run: func(cmd *cobra.Command, args []string) {
// 			fmt.Println("Print: " + strings.Join(args, " "))
// 		},
// 	}

// 	var cmdTimes = &cobra.Command{
// 		Use:   "times [# times] [string to echo]",
// 		Short: "Echo anything to the screen more times",
// 		Long: `echo things multiple times back to the user by providing
// a count and a string.`,
// 		Args: cobra.MinimumNArgs(1),
// 		Run: func(cmd *cobra.Command, args []string) {
// 			fmt.Println(echoTimes)
// 			for i := 0; i < echoTimes; i++ {
// 				fmt.Println("Echo: " + strings.Join(args, " "))
// 			}
// 		},
// 	}

// 	cmdTimes.Flags().IntVarP(&echoTimes, "times", "t", 1, "times to echo the input")

// 	var rootCmd = &cobra.Command{Use: "app"}
// 	rootCmd.AddCommand(cmdPrint, cmdEcho)
// 	cmdEcho.AddCommand(cmdTimes)
// 	rootCmd.Execute()
// }
