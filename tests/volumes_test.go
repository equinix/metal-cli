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

var volumeID string

func TestVolumeOperations(t *testing.T) {
	client, _ = packngo.NewClientWithBaseURL("Packet CLI", os.Getenv("PACKET_TOKEN"), nil, "https://api.packet.net/")
	projects, _, _ := client.Projects.List(nil)
	projectID = projects[0].ID
	setupTests := []Test{
		{
			"create volume",
			[]string{
				"volume", "create",
				"--size", "25",
				"--plan", "storage_1",
				"--facility", "ewr1",
				"--project-id", projectID,
				"-j",
			},
		},
		{
			"create device",
			[]string{
				"device", "create",
				"--hostname", "clivolumestest",
				"--plan", "baremetal_1",
				"--facility", "ewr1",
				"--operating-system", "centos_7",
				"--project-id", projectID,
				"-j",
			},
		},
	}
	tests := []Test{

		{"volumes list", []string{"volume", "get", "project-id", projectID, "-j"}},
		{"volume get", []string{"volume", "get", "-i"}},
		{"volume attach", []string{"volume", "attach", "-i"}},
		{"volume detach", []string{"volume", "detach", "-i"}},
	}
	cleanUp := []Test{

		{"volume delete", []string{"volume", "delete", "-i"}},
		{"device delete", []string{"device", "delete", "-i"}},
	}

	for _, tt := range setupTests {
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
			if strings.Contains(actual, "Error:") {
				t.Fatal(actual)
			}
			if len(tt.args) > 0 {
				if tt.args[0] == "device" && tt.args[1] == "create" {
					device := &packngo.Device{}
					// fmt.Println(actual)
					err := json.Unmarshal([]byte(actual), device)
					if err != nil {
						t.Fatal(err)
					}

					deviceID = (*device).ID
					for {
						dev, _, err := client.Devices.Get(deviceID)
						if err != nil {
							break
						}
						if dev.State == "active" {
							break
						}
						time.Sleep(2 * time.Second)
					}
				}
				if tt.args[0] == "volume" && tt.args[1] == "create" {
					volume := &packngo.Volume{}
					// fmt.Println(actual)
					err := json.Unmarshal([]byte(actual), volume)
					if err != nil {
						t.Fatal(err)
					}

					volumeID = (*volume).ID
					for {
						vol, _, err := client.Volumes.Get(volumeID)
						if err != nil {
							break
						}
						if vol.State == "active" {
							break
						}
						time.Sleep(2 * time.Second)
					}
				}
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(tt.name, tt.args)

			dir, err := os.Getwd()
			if err != nil {
				t.Fatal(err)
			}

			if tt.name == "volume get" && volumeID != "" {
				tt.args = append(tt.args, volumeID, "-j")
			}

			if tt.name == "volume attach" && volumeID != "" {
				tt.args = append(tt.args, volumeID, "--device-id", deviceID)
			}

			if tt.name == "volume detach" && volumeID != "" {
				tt.args = append(tt.args, volumeID)
			}

			cmd := exec.Command(path.Join(dir, binaryName), tt.args...)

			output, err := cmd.CombinedOutput()
			if err != nil {
				t.Fatal(err)
			}

			actual := string(output)
			if strings.Contains(actual, "Error:") {
				t.Fatal(actual)
			}
			if len(tt.args) > 0 {
				if tt.args[0] == "volume" && (tt.args[1] == "attach" || tt.args[1] == "detach") {
					for {
						vol, _, err := client.Volumes.Get(volumeID)
						if err != nil {
							break
						}
						if vol.State == "active" {
							break
						}
						time.Sleep(2 * time.Second)
					}
				}
			}
		})
	}

	for _, tt := range cleanUp {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(tt.name, tt.args)

			dir, err := os.Getwd()
			if err != nil {
				t.Fatal(err)
			}
			if tt.name == "volume delete" && volumeID != "" {
				tt.args = append(tt.args, volumeID, "-f")
			}

			if tt.name == "device delete" && deviceID != "" {
				tt.args = append(tt.args, deviceID, "-f")
			}

			cmd := exec.Command(path.Join(dir, binaryName), tt.args...)

			output, err := cmd.CombinedOutput()
			if err != nil {
				t.Fatal(err)
			}

			actual := string(output)
			if strings.Contains(actual, "Error:") {
				t.Fatal(actual)
			}
		})
	}
}
