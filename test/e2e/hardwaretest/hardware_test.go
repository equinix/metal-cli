package hardwaretest

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	root "github.com/equinix/metal-cli/internal/cli"
	"github.com/equinix/metal-cli/internal/hardware"
	outputPkg "github.com/equinix/metal-cli/internal/outputs"
	"github.com/equinix/metal-cli/test/helper"
	"github.com/spf13/cobra"
)

func MockRootClient(responseBody string) *root.Client {
	mockAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		_, err := w.Write([]byte(responseBody))
		if err != nil {
			log.Fatalf("Failed to write mock response: %v", err)
		}
	}))
	mockClient := root.NewClient("", mockAPI.URL, "metal")
	return mockClient

}

func TestCli_Hardware(t *testing.T) {
	subCommand := "hardware-reservation"
	// Adjust this response as needed for your tests.
	mockResponse := `{
		"hardware_reservations": [
		  {
			"created_at": "2019-08-24T14:15:22Z",
			"custom_rate": 1050.5,
			"id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
			"facility": {
			  "code": "da",
			  "id": "497f6eca-6276-4993-bfeb-53cbbbba6f01",
			  "metro": {
				"code": "da",
				"country": "string",
				"id": "497f6eca-6276-4993-bfeb-53cbbbba6f02",
				"name": "string"
			  },
			  "name": "string"
			},
			"plan": {
			  "id": "497f6eca-6276-4993-bfeb-53cbbbba6f03",
			  "name": "m3.large.x86",
			  "slug": "m3.large.x86"
			},
			"project": {
			  "id": "497f6eca-6276-4993-bfeb-53cbbbba6f04"
			},
			"provisionable": true,
			"short_id": "string",
			"spare": true,
			"switch_uuid": "string",
			"termination_time": "2019-08-24T14:15:22Z"
		  }
		]
	  }`

	rootClient := MockRootClient(mockResponse)

	type fields struct {
		MainCmd  *cobra.Command
		Outputer outputPkg.Outputer
	}
	tests := []struct {
		name    string
		fields  fields
		want    *cobra.Command
		cmdFunc func(*testing.T, *cobra.Command)
	}{
		{
			name: "get",
			fields: fields{
				MainCmd:  hardware.NewClient(rootClient, outputPkg.Outputer(&outputPkg.Standard{})).NewCommand(),
				Outputer: outputPkg.Outputer(&outputPkg.Standard{}),
			},
			want: &cobra.Command{},
			cmdFunc: func(t *testing.T, c *cobra.Command) {
				root := c.Root()
				projectName := "metal-cli-" + helper.GenerateRandomString(5) + "-hardware-test"
				project := helper.CreateTestProject(t, projectName)
				root.SetArgs([]string{subCommand, "get", "-p", project.GetId()})
				rescueStdout := os.Stdout
				r, w, _ := os.Pipe()
				os.Stdout = w
				if err := root.Execute(); err != nil {
					t.Error(err)
				}
				w.Close()
				out, _ := io.ReadAll(r)

				os.Stdout = rescueStdout
				if !strings.Contains(string(out[:]), "da") &&
					!strings.Contains(string(out[:]), "m3.large.x86") {
					t.Error("expected output should include m3.large.x86.")
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rootCmd := rootClient.NewCommand()
			rootCmd.AddCommand(tt.fields.MainCmd)
			tt.cmdFunc(t, tt.fields.MainCmd)
		})
	}
}
