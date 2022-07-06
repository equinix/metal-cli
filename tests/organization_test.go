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

var organizationID string

func TestOrganizationOperations(t *testing.T) {
	client, _ = packngo.NewClientWithBaseURL(consumerToken, testToken(), nil, apiURL)
	setupTests := []Test{
		{
			"create organization",
			[]string{
				"organization", "create",
				"--name", "clitestOrg",
				"-j",
			},
		},
	}
	tests := []Test{
		{"organization list", []string{"organization", "get"}},
		{"organization get", []string{"organization", "get", "-i"}},
		{"organization update", []string{"organization", "update", "-n", "updatednamefromCLI", "-i"}},
	}
	cleanUp := []Test{
		{"organization delete", []string{"organization", "delete", "-i"}},
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

			if tt.args[0] == "organization" && tt.args[1] == "create" {
				organization := &packngo.Organization{}
				err := json.Unmarshal([]byte(actual), organization)
				if err != nil {
					t.Fatal(err)
				}

				organizationID = (*organization).ID
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

			if (tt.name == "organization get" ||
				tt.name == "organization update") && organizationID != "" {
				tt.args = append(tt.args, organizationID)
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
			if tt.name == "organization delete" && organizationID != "" {
				tt.args = append(tt.args, organizationID)
				tt.args = append(tt.args, "-f")
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
