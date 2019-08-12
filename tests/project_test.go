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

func TestProjectOperations(t *testing.T) {
	client, _ = packngo.NewClientWithBaseURL("Packet CLI", os.Getenv("PACKET_TOKEN"), nil, "https://api.packet.net/")
	orgs, _, _ := client.Organizations.List()
	orgID := orgs[0].ID

	tests := []Test{
		{"project create", []string{"project", "create", "-n", "test", "-o", orgID, "-j"}},
		{"project get", []string{"project", "get"}},
	}

	sshKeys, _, _ := client.SSHKeys.List()
	for _, key := range sshKeys {
		if key.Label == "test" {
			sshKeyID = key.ID
			fmt.Println("sshkeyID", sshKeyID)

			break
		}
	}

	cleanup := []Test{
		{"project get", []string{"project", "get", "-i"}},
		{"project delete", []string{"project", "delete", "-i"}},
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
			if strings.Contains(actual, "Error:") {
				t.Fatal(actual)
			}

			if tt.name == "project create" {
				project := &packngo.Project{}
				err := json.Unmarshal([]byte(actual), project)
				if err != nil {
					t.Fatal(err)
				}

				projectID = (*project).ID
			}
		})
	}

	for _, tt := range cleanup {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "project get" && projectID != "" {
				tt.args = append(tt.args, projectID)
			}

			if tt.name == "project delete" && projectID != "" {
				tt.args = append(tt.args, projectID)
				tt.args = append(tt.args, "-f")
			}
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
		})
	}
}
