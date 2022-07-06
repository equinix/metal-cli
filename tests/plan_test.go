package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
	"testing"
)

func TestPlanOperations(t *testing.T) {
	tests := []Test{
		{"plan get", []string{"plan", "get"}},
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
