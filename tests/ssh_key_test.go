package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"
	"testing"

	"github.com/packethost/packngo"
)

var sshKeyID string

func TestSSHKeyOperations(t *testing.T) {
	client, _ = packngo.NewClientWithBaseURL("Packet CLI", os.Getenv("PACKET_TOKEN"), nil, "https://api.packet.net/")

	data, err := ioutil.ReadFile(os.Getenv("HOME") + "/.ssh/id_rsa.pub")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	key := string(data)

	tests := []Test{
		{"ssh-key create", []string{"ssh-key", "create", "-l", "test", "-k", key}},
		{"ssh-key list", []string{"ssh-key", "get", "-j"}},
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
		{"ssh-key get", []string{"ssh-key", "get", "-i"}},
		{"ssh-key delete", []string{"ssh-key", "delete", "-i"}},
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
		})
	}

	for _, tt := range cleanup {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "ssh-key get" && sshKeyID != "" {
				tt.args = append(tt.args, sshKeyID)
			}

			if tt.name == "ssh-key delete" && sshKeyID != "" {
				tt.args = append(tt.args, sshKeyID)
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
