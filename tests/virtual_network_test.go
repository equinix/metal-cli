package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
	"testing"

	"github.com/packethost/packngo"
)

var virtualNetworkID string

func TestVirtualNetworkOperations(t *testing.T) {
	client, _ = packngo.NewClientWithBaseURL(consumerToken, testToken(), nil, apiURL)
	projects, _, _ := client.Projects.List(nil)
	projectID = projects[0].ID
	setupTests := []Test{
		{
			"create virtual-network",
			[]string{
				"virtual-network", "create",
				"--facility", "ewr1",
				"--project-id", projectID,
				"-j",
			},
		},
	}
	tests := []Test{
		{"virtual-network list", []string{"virtual-network", "get", "-p", projectID}},
	}
	cleanUp := []Test{
		{"virtual-network delete", []string{"virtual-network", "delete", "-i"}},
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
			if strings.Contains(strings.ToLower(actual), "error:") {
				t.Fatal(actual)
			}

			if tt.args[0] == "virtual-network" && tt.args[1] == "create" {
				virtualNetwork := &packngo.VirtualNetwork{}
				err := json.Unmarshal([]byte(actual), virtualNetwork)
				if err != nil {
					t.Fatal(err)
				}

				virtualNetworkID = (*virtualNetwork).ID
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

	for _, tt := range cleanUp {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(tt.name, tt.args)

			dir, err := os.Getwd()
			if err != nil {
				t.Fatal(err)
			}
			if tt.name == "virtual-network delete" && virtualNetworkID != "" {
				tt.args = append(tt.args, virtualNetworkID)
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
