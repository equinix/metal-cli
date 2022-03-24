package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
	"testing"

	"github.com/packethost/packngo"
)

const binaryName = "metal"

const (
	consumerToken  = "Equinix Metal CLI"
	apiTokenEnvVar = "METAL_AUTH_TOKEN"
	apiURL         = "https://api.equinix.com/metal/v1/"
)

func TestMain(m *testing.M) {
	err := os.Chdir("..")
	if err != nil {
		fmt.Printf("could not change dir: %v", err)
		os.Exit(1)
	}
	fmt.Println("build")
	build := exec.Command("go", "build", "-v", "-o", binaryName, "./cmd/metal")
	err = build.Run()
	if err != nil {
		fmt.Printf("could not build binary for %s: %v", binaryName, err)
		os.Exit(1)
	}

	os.Exit(m.Run())
}

var projectID string
var deviceID string
var client *packngo.Client

type Test struct {
	name string
	args []string
}

func testToken() string {
	return os.Getenv(apiTokenEnvVar)
}

func TestCliArgs(t *testing.T) {
	client, _ = packngo.NewClientWithBaseURL(consumerToken, testToken(), nil, apiURL)
	//projects, _, _ := client.Projects.List(nil)
	//projectID = projects[0].ID
	tests := []Test{
		{"no arguments", []string{}},
		{"version", []string{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(tt.name, tt.args)

			dir, err := os.Getwd()
			if err != nil {
				t.Fatal(err)
			}

			cmd := exec.Command(path.Join(dir, binaryName), tt.args...)

			output, err := cmd.CombinedOutput()
			if err != nil {
				t.Fatal(err)
			}

			actual := string(output)
			if strings.Contains(strings.ToLower(actual), "error:") {
				t.Fatal(actual)
			}
		})
	}
}
