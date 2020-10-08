package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
	"testing"

	"github.com/packethost/packngo"
	"golang.org/x/crypto/ssh"
)

var sshKeyID string

func TestSSHKeyOperations(t *testing.T) {
	client, _ = packngo.NewClientWithBaseURL(consumerToken, testToken(), nil, apiURL)

	publicKey, err := generatePublicKey()
	if err != nil {
		fmt.Println("SSH Key generation error:", err)
	}

	tests := []Test{
		{"ssh-key create", []string{"ssh-key", "create", "-l", "test", "-k", publicKey}},
		{"ssh-key list", []string{"ssh-key", "get", "-j"}},
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
			if strings.Contains(strings.ToLower(actual), "error") {
				t.Fatal(actual)
			}
		})
	}

	sshKeys, _, _ := client.SSHKeys.List()
	for _, key := range sshKeys {
		if key.Label == "test" {
			sshKeyID = key.ID
			break
		}
	}
	fmt.Println("outside", sshKeyID)

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
			if strings.Contains(strings.ToLower(actual), "error:") {
				t.Fatal(actual)
			}
		})
	}
}

func generatePublicKey() (string, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return "", err
	}

	// Validate Private Key
	err = privateKey.Validate()
	if err != nil {
		return "", err
	}

	publicRsaKey, err := ssh.NewPublicKey(&privateKey.PublicKey)
	if err != nil {
		return "", err
	}

	pubKeyBytes := ssh.MarshalAuthorizedKey(publicRsaKey)
	return string(pubKeyBytes), nil
}
