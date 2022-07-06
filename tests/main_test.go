package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
	"testing"
	"time"

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
	build := exec.Command("go", "build", "-o", binaryName)
	err = build.Run()
	if err != nil {
		fmt.Printf("could not build binary for %s: %v", binaryName, err)
		os.Exit(1)
	}

	os.Exit(m.Run())
}

var (
	projectID string
	deviceID  string
	client    *packngo.Client
)

type Test struct {
	name string
	args []string
}

func testToken() string {
	return os.Getenv(apiTokenEnvVar)
}

func TestCliArgs(t *testing.T) {
	client, _ = packngo.NewClientWithBaseURL(consumerToken, testToken(), nil, apiURL)
	projects, _, _ := client.Projects.List(nil)
	projectID = projects[0].ID
	tests := []Test{
		{"no arguments", []string{}},
		{"operating-systems get", []string{"operating-systems", "get", "-j"}},
		{"plan get", []string{"plan", "get", "-j"}},
		{"organization get", []string{"organization", "get", "-j"}},
		{"project get", []string{"project", "get", "-j"}},
		{
			"create device",
			[]string{
				"device", "create",
				"--hostname", "clitest",
				"--plan", "baremetal_1",
				"--facility", "ewr1",
				"--operating-system", "centos_7",
				"--project-id", projectID,
				"-j",
			},
		},
		{"devices get", []string{"device", "get", "project-id", projectID, "-j"}},
		{"device update", []string{"device", "update", "hostname", "updatedfromcli", "-i"}},
		{"device reboot", []string{"device", "reboot", "-i"}},
		{"device stop", []string{"device", "stop", "-i"}},
		{"device start", []string{"device", "start", "-i"}},
		{"device delete", []string{"device", "delete", "-i"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(tt.name, tt.args)

			dir, err := os.Getwd()
			if err != nil {
				t.Fatal(err)
			}

			if tt.name == "device delete" && deviceID != "" {
				tt.args = append(tt.args, deviceID, "-f")
			}

			if (tt.name == "device update" ||
				tt.name == "device reboot" ||
				tt.name == "device stop" ||
				tt.name == "device start") && deviceID != "" {
				tt.args = append(tt.args, deviceID)
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
			if len(tt.args) > 0 {
				if tt.args[0] == "project" {
					project := &[]packngo.Project{}
					err := json.Unmarshal([]byte(actual), project)
					if err != nil {
						t.Fatal(err)
					}
					projectID = (*project)[0].ID
				} else if tt.args[0] == "device" && tt.args[1] == "create" {
					device := &packngo.Device{}
					// fmt.Println(actual)
					err := json.Unmarshal([]byte(actual), device)
					if err != nil {
						t.Fatal(err)
					}

					deviceID = (*device).ID
					for {
						dev, _, err := client.Devices.Get(deviceID, nil)
						if err != nil {
							break
						}
						if dev.State == "active" {
							break
						}
						time.Sleep(2 * time.Second)
					}
				}
			}
		})
	}
}
