package hardwaretest

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	root "github.com/equinix/metal-cli/internal/cli"
	outputPkg "github.com/equinix/metal-cli/internal/outputs"
	"github.com/equinix/metal-cli/internal/twofa"
	"github.com/spf13/cobra"
)

var mockOtpUri = "otpauth://totp/foo"

func setupMock() *root.Client {
	mockAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var responseBody string
		if r.URL.Path == "/user/otp/sms/receive" {
			w.WriteHeader(http.StatusNoContent)
		} else if r.URL.Path == "/user/otp/app/receive" {
			w.Header().Add("Content-Type", "application/json")
			responseBody = fmt.Sprintf(`{"otp_uri": "%v"}`, mockOtpUri)

		} else {
			responseBody = fmt.Sprintf("no mock for endpoint %v", r.URL.Path)
			w.WriteHeader(http.StatusNotImplemented)
		}
		_, err := w.Write([]byte(responseBody))
		if err != nil {
			log.Fatalf("Failed to write mock response: %v", err)
		}
	}))
	mockClient := root.NewClient("", mockAPI.URL, "metal")
	return mockClient

}

func TestCli_Twofa(t *testing.T) {
	subCommand := "2fa"
	// Adjust this response as needed for your tests.

	rootClient := setupMock()

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
			name: "receive sms",
			fields: fields{
				MainCmd:  twofa.NewClient(rootClient, outputPkg.Outputer(&outputPkg.Standard{})).NewCommand(),
				Outputer: outputPkg.Outputer(&outputPkg.Standard{}),
			},
			want: &cobra.Command{},
			cmdFunc: func(t *testing.T, c *cobra.Command) {
				root := c.Root()
				root.SetArgs([]string{subCommand, "receive", "-s"})
				rescueStdout := os.Stdout
				r, w, _ := os.Pipe()
				os.Stdout = w
				if err := root.Execute(); err != nil {
					t.Error(err)
				}
				w.Close()
				out, _ := io.ReadAll(r)

				os.Stdout = rescueStdout
				if !strings.Contains(string(out[:]), "SMS token sent to your phone") {
					t.Error("expected output to include 'SMS token sent to your phone'.")
				}
			},
		},
		{
			name: "receive app",
			fields: fields{
				MainCmd:  twofa.NewClient(rootClient, outputPkg.Outputer(&outputPkg.Standard{})).NewCommand(),
				Outputer: outputPkg.Outputer(&outputPkg.Standard{}),
			},
			want: &cobra.Command{},
			cmdFunc: func(t *testing.T, c *cobra.Command) {
				root := c.Root()
				root.SetArgs([]string{subCommand, "receive", "-a"})
				rescueStdout := os.Stdout
				r, w, _ := os.Pipe()
				os.Stdout = w
				if err := root.Execute(); err != nil {
					t.Error(err)
				}
				w.Close()
				out, _ := io.ReadAll(r)

				os.Stdout = rescueStdout
				if !strings.Contains(string(out[:]), mockOtpUri) {
					t.Errorf("expected output to include %v", mockOtpUri)
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
