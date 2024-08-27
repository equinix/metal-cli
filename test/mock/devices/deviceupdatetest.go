package deviceupdatetest

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	metal "github.com/equinix/equinix-sdk-go/services/metalv1"
	root "github.com/equinix/metal-cli/internal/cli"
	"github.com/equinix/metal-cli/internal/devices"
	outputPkg "github.com/equinix/metal-cli/internal/outputs"
	"github.com/equinix/metal-cli/test/helper"
	"github.com/spf13/cobra"
)

type mockRoundTripper struct {
	handler func(req *http.Request) (*http.Response, error)
}

func (m *mockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return m.handler(req)
}

func TestCli_Devices_Update(t *testing.T) {
	var setupMockClientOpt root.ClientOpt
	setupMockClientFn := func(c *root.Client, httpClient *http.Client) *metal.APIClient {
		cfg := metal.NewConfiguration()
		httpClient.Transport = &mockRoundTripper{handler: func(req *http.Request) (*http.Response, error) {
			if strings.Contains(req.URL.Path, "/devices/") && req.Method == http.MethodPut {
				body, _ := io.ReadAll(req.Body)
				if !strings.Contains(string(body), `"locked": true`) {
					t.FailNow()
				}

				return &http.Response{
					Body:       io.NopCloser(strings.NewReader(`{}`)),
					Header:     http.Header{"Content-Type": []string{"application/json"}},
					StatusCode: http.StatusOK,
				}, nil
			}
			return nil, fmt.Errorf("unknown request: %s %s", req.Method, req.URL.Path)
		},
		}
		cfg.HTTPClient = httpClient
		return metal.NewAPIClient(cfg)
	}
	setupMockClientOpt = func(c *root.Client) {
		c.SetMetalAPIConnect(setupMockClientFn)
	}

	subCommand := "device"
	rootClient := root.NewClient(helper.ConsumerToken, helper.URL, helper.Version, setupMockClientOpt)
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
			name: "update_device",
			fields: fields{
				MainCmd:  devices.NewClient(rootClient, outputPkg.Outputer(&outputPkg.Standard{})).NewCommand(),
				Outputer: outputPkg.Outputer(&outputPkg.Standard{}),
			},
			want: &cobra.Command{},
			cmdFunc: func(t *testing.T, c *cobra.Command) {
				root := c.Root()

				root.SetArgs([]string{subCommand, "update", "-i", "1234", "-H", "metal-cli-update-dev-test", "--locked", "true"})

				out := helper.ExecuteAndCaptureOutput(t, root)

				if !strings.Contains(string(out[:]), "metal-cli-update-dev-test") {
					t.Error("expected output should include metal-cli-update-dev-test in the out string ")
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
