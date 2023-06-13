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

func Create_test_project(name string) string {
	TestApiClient := TestClient()

	projectCreateFromRootInput := *openapiclient.NewProjectCreateFromRootInput(name) // ProjectCreateFromRootInput | Project to create
	include := []string{"Inner_example"}
	exclude := []string{"Inner_example"}

	projectResp, r, err := TestApiClient.ProjectsApi.CreateProject(context.Background()).ProjectCreateFromRootInput(projectCreateFromRootInput).Include(include).Exclude(exclude).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.CreateProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		return ""
	}
	return projectResp.GetId()
}

func Create_test_device(projectId, name string) string {
	include := []string{"Inner_example"}
	exclude := []string{"Inner_example"}

	TestApiClient := TestClient()

	hostname := name
	metroDeviceRequest := openapiclient.CreateDeviceRequest{
		DeviceCreateInMetroInput: &openapiclient.DeviceCreateInMetroInput{
			Metro:           "da",
			Plan:            "c3.small.x86",
			OperatingSystem: "ubuntu_20_04",
			Hostname:        &hostname,
		},
	}
	deviceResp, _, err := TestApiClient.DevicesApi.CreateDevice(context.Background(), projectId).CreateDeviceRequest(metroDeviceRequest).Include(include).Exclude(exclude).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.CreateDevice``: %v\n", err)
		return ""
	}
	return deviceResp.GetId()
}

func Is_Device_state_active(deviceId string) bool {
	TestApiClient := TestClient()
	predefinedTime := 200 * time.Second // Adjust this as needed
	retryInterval := 10 * time.Second   // Adjust this as needed
	startTime := time.Now()
	for time.Since(startTime) < predefinedTime {
		resp, _, err := TestApiClient.DevicesApi.FindDeviceById(context.Background(), deviceId).Execute()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.FindDeviceById``: %v\n", err)
		}
		if resp.GetState() == "active" {
			return true
		}

		// Sleep for the specified interval
		time.Sleep(retryInterval)
	}
	return false
}

func Clean_test_device(deviceId string) {
	forceDelete := true // bool | Force the deletion of the device, by detaching any storage volume still active. (optional)

	TestApiClient := TestClient()
	_, err := TestApiClient.DevicesApi.DeleteDevice(context.Background(), deviceId).ForceDelete(forceDelete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.DeleteDevice``: %v\n", err)
	}
}

func Clean_test_project(projectId string) {
	TestApiClient := TestClient()
	r, err := TestApiClient.ProjectsApi.DeleteProject(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.DeleteProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
