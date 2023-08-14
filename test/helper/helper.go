package helper

import (
	"context"
	"fmt"
	"os"
	"time"

	openapiclient "github.com/equinix-labs/metal-go/metal/v1"
)

func TestClient() *openapiclient.APIClient {
	configuration := openapiclient.NewConfiguration()
	configuration.AddDefaultHeader("X-Auth-Token", os.Getenv("METAL_AUTH_TOKEN"))
	apiClient := openapiclient.NewAPIClient(configuration)
	return apiClient
}

// func Create_test_project(name string) string {
func CreateTestProject(name string) (string, error) {
	TestApiClient := TestClient()

	projectCreateFromRootInput := *openapiclient.NewProjectCreateFromRootInput(name) // ProjectCreateFromRootInput | Project to create

	projectResp, r, err := TestApiClient.ProjectsApi.CreateProject(context.Background()).ProjectCreateFromRootInput(projectCreateFromRootInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.CreateProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		return "", err
	}
	return projectResp.GetId(), nil
}

func CreateTestDevice(projectId, name string) (string, error) {
	TestApiClient := TestClient()

	hostname := name
	metroDeviceRequest := openapiclient.CreateDeviceRequest{
		DeviceCreateInMetroInput: &openapiclient.DeviceCreateInMetroInput{
			Metro:           "da",
			Plan:            "m3.small.x86",
			OperatingSystem: "ubuntu_20_04",
			Hostname:        &hostname,
		},
	}
	deviceResp, _, err := TestApiClient.DevicesApi.CreateDevice(context.Background(), projectId).CreateDeviceRequest(metroDeviceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.CreateDevice``: %v\n", err)
		return "", err
	}
	return deviceResp.GetId(), nil
}

func IsDeviceStateActive(deviceId string) (bool, error) {
	var err error
	var resp *openapiclient.Device
	TestApiClient := TestClient()
	predefinedTime := 500 * time.Second // Adjust this as needed
	retryInterval := 10 * time.Second   // Adjust this as needed
	startTime := time.Now()
	for time.Since(startTime) < predefinedTime {
		resp, _, err = TestApiClient.DevicesApi.FindDeviceById(context.Background(), deviceId).Execute()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.FindDeviceById``: %v\n", err)
			return false, err
		}
		if resp.GetState() == "active" {
			return true, nil
		}

		// Sleep for the specified interval
		time.Sleep(retryInterval)
	}
	return false, err
}

func StopTestDevice(deviceId string) error {

	deviceActionInput := *openapiclient.NewDeviceActionInput("power_off")
	TestApiClient := TestClient()

	_, err := TestApiClient.DevicesApi.PerformAction(context.Background(), deviceId).DeviceActionInput(deviceActionInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.PerformAction``: %v\n", err)
		return err
	}
	return nil
}

func CleanTestDevice(deviceId string) error {
	forceDelete := true // bool | Force the deletion of the device, by detaching any storage volume still active. (optional)

	TestApiClient := TestClient()
	_, err := TestApiClient.DevicesApi.DeleteDevice(context.Background(), deviceId).ForceDelete(forceDelete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.DeleteDevice``: %v\n", err)
		return err
	}
	return nil
}

func CleanTestProject(projectId string) error {
	TestApiClient := TestClient()
	r, err := TestApiClient.ProjectsApi.DeleteProject(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.DeleteProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		return err
	}
	return nil
}
