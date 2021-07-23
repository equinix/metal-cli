// Copyright © 2021 Equinix Metal Developers <support@equinixmetal.com>
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

import (
	"os"
	"runtime/debug"

	"github.com/equinix/metal-cli/cmd"
)

func main() {
	cli := cmd.NewCli()
	if err := cli.MainCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	// look for the default version and replace it, if found, from runtime build info
	if cmd.Version != "devel" {
		return
	}
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		return
	}

	// Version is set in artifacts built with -X github.com/equinix/metal-cli/cmd.Version=1.2.3
	// Ensure version is also set when installed via go install github.com/equinix/metal-cli/cmd/metal
	cmd.Version = bi.Main.Version
}
