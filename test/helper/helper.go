package helper

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	openapiclient "github.com/equinix-labs/metal-go/metal/v1"
)

func TestClient() *openapiclient.APIClient {
	configuration := openapiclient.NewConfiguration()
	configuration.AddDefaultHeader("X-Auth-Token", os.Getenv("METAL_AUTH_TOKEN"))
	apiClient := openapiclient.NewAPIClient(configuration)
	return apiClient
}

func CreateTestProject(name string) (string, error) {
	TestApiClient := TestClient()

	projectCreateFromRootInput := *openapiclient.NewProjectCreateFromRootInput(name) // ProjectCreateFromRootInput | Project to create

	projectResp, r, err := TestApiClient.ProjectsApi.CreateProject(context.Background()).ProjectCreateFromRootInput(projectCreateFromRootInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.CreateProject`: %v\n", err)
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
			Metro:           "sv",
			Plan:            "m3.small.x86",
			OperatingSystem: "ubuntu_20_04",
			Hostname:        &hostname,
		},
	}
	deviceResp, _, err := TestApiClient.DevicesApi.
		CreateDevice(context.Background(), projectId).
		CreateDeviceRequest(metroDeviceRequest).
		Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.CreateDevice`: %v\n", err)
		return "", err
	}
	return deviceResp.GetId(), nil
}

func CreateTestVLAN(projectId string) (*openapiclient.VirtualNetwork, error) {
	TestApiClient := TestClient()

	metro := "sv"
	vlanCreateInput := openapiclient.VirtualNetworkCreateInput{
		Metro: &metro,
	}
	vlan, _, err := TestApiClient.VLANsApi.
		CreateVirtualNetwork(context.Background(), projectId).
		VirtualNetworkCreateInput(vlanCreateInput).
		Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VLANsApi.CreateVirtualNetwork`: %v\n", err)
		return nil, err
	}
	return vlan, nil
}

func CreateTestGateway(projectId, vlanId string, privateIPv4SubnetSize *int32) (*openapiclient.MetalGateway, error) {
	TestApiClient := TestClient()

	gatewayCreateInput := openapiclient.CreateMetalGatewayRequest{
		MetalGatewayCreateInput: &openapiclient.MetalGatewayCreateInput{
			VirtualNetworkId:      vlanId,
			PrivateIpv4SubnetSize: privateIPv4SubnetSize,
		},
	}
	gateway, _, err := TestApiClient.MetalGatewaysApi.
		CreateMetalGateway(context.Background(), projectId).
		Include([]string{"ip_reservation"}).
		CreateMetalGatewayRequest(gatewayCreateInput).
		Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetalGatewaysApi.CreateMetalGateway`: %v\n", err)
		return nil, err
	}
	return gateway.MetalGateway, nil
}

func GetDeviceById(deviceId string) (*openapiclient.Device, error) {
	TestApiClient := TestClient()
	includes := []string{"network_ports"}

	device, _, err := TestApiClient.DevicesApi.
		FindDeviceById(context.Background(), deviceId).
		Include(includes).
		Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.FindDeviceById`: %v\n", err)
		return nil, err
	}

	return device, nil
}

func GetPortById(portId string) (*openapiclient.Port, error) {
	TestApiClient := TestClient()
	includes := []string{"virtual_network"}

	port, _, err := TestApiClient.PortsApi.
		FindPortById(context.Background(), portId).
		Include(includes).
		Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortsApi.FindPortById`: %v\n", err)
		return nil, err
	}

	return port, nil
}

func IsDeviceStateActive(deviceId string) (bool, error) {
	predefinedTime := 500 * time.Second // Adjust this as needed
	retryInterval := 10 * time.Second   // Adjust this as needed
	startTime := time.Now()
	for time.Since(startTime) < predefinedTime {
		device, err := GetDeviceById(deviceId)
		if err != nil {
			return false, err
		}
		if device.GetState() == "active" {
			return true, nil
		}

		// Sleep for the specified interval
		time.Sleep(retryInterval)
	}
	return false, fmt.Errorf("timed out waiting for device %v to become active", deviceId)
}

func WaitForAttachVlanToPort(portId string, attach bool) error {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	timeout := 300 * time.Second
	ctx, cancelFunc := context.WithTimeout(context.Background(), timeout)
	defer cancelFunc()

	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("Timeout exceeded for vlan assignment with port ID: %s", portId)
		case <-ticker.C:
			port, err := GetPortById(portId)
			if err != nil {
				return err
			}

			vlans := port.GetVirtualNetworks()
			if attach {
				if len(vlans) != 0 {
					return nil
				}
			} else {
				if len(vlans) == 0 {
					return nil
				}
			}
		}
	}
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
	_, err := TestApiClient.DevicesApi.
		DeleteDevice(context.Background(), deviceId).
		ForceDelete(forceDelete).
		Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesApi.DeleteDevice``: %v\n", err)
		return err
	}
	return nil
}

func CleanTestProject(projectId string) error {
	TestApiClient := TestClient()
	r, err := TestApiClient.ProjectsApi.
		DeleteProject(context.Background(), projectId).
		Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.DeleteProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		return err
	}
	return nil
}

func CreateTestIps(projectId string, quantity int, ipType string) (string, error) {
	TestApiClient := TestClient()
	metro := "da"
	var tags []string
	var facility string

	req := &openapiclient.IPReservationRequestInput{
		Metro:    &metro,
		Tags:     tags,
		Quantity: int32(quantity),
		Type:     ipType,
		Facility: &facility,
	}

	requestIPReservationRequest := &openapiclient.RequestIPReservationRequest{
		IPReservationRequestInput: req,
	}

	ipsresp, _, err := TestApiClient.IPAddressesApi.RequestIPReservation(context.Background(), projectId).RequestIPReservationRequest(*requestIPReservationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPAddressesApi.FindIPReservations``: %v\n", err)
		return "", err
	}
	return ipsresp.IPReservation.GetId(), nil
}

func CleanTestIps(ipsId string) error {
	TestApiClient := TestClient()
	_, err := TestApiClient.IPAddressesApi.DeleteIPAddress(context.Background(), ipsId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPAddressesApi.DeleteIPAddress``: %v\n", err)
		return err
	}
	return nil
}

func CreateTestVlanWithVxLan(projectId string, Id int, desc string) (string, error) {
	TestApiClient := TestClient()
	virtualNetworkCreateInput := *openapiclient.NewVirtualNetworkCreateInput()
	virtualNetworkCreateInput.SetDescription(desc)
	virtualNetworkCreateInput.SetMetro("da")
	virtualNetworkCreateInput.SetVxlan(int32(Id))

	vlanresp, _, err := TestApiClient.VLANsApi.CreateVirtualNetwork(context.Background(), projectId).VirtualNetworkCreateInput(virtualNetworkCreateInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VLANsApi.CreateVirtualNetwork``: %v\n", err)
		return "", err
	}
	return vlanresp.GetId(), nil
}

func CleanTestVlan(vlanId string) error {
	TestApiClient := TestClient()
	_, _, err := TestApiClient.VLANsApi.DeleteVirtualNetwork(context.Background(), vlanId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VLANsApi.DeleteVirtualNetwork``: %v\n", err)
		return err
	}

	return nil
}

func UnAssignPortVlan(portId, vlanId string) error {
	testClient := TestClient()
	_, _, err := testClient.PortsApi.
		UnassignPort(context.Background(), portId).
		PortAssignInput(openapiclient.PortAssignInput{Vnid: &vlanId}).
		Execute()
	return err
}

func CleanupProjectAndDevice(deviceId, projectId string) error {
	resp, err := IsDeviceStateActive(deviceId)
	if err == nil && resp {
		err = CleanTestDevice(deviceId)
		if err != nil {
			return err
		}
		err = CleanTestProject(projectId)
		if err != nil {
			return err
		}
	}

	return nil
}

//nolint:staticcheck
func SetupProjectAndDevice(t *testing.T, projectId, deviceId *string) *openapiclient.Device {
	t.Helper()
	projId, err := CreateTestProject("metal-cli-test-project")
	if err != nil {
		t.Fatal(err)
		return nil
	}
	*projectId = projId

	devId, err := CreateTestDevice(*projectId, "metal-cli-test-device")
	if err != nil {
		t.Fatal(err)
		return nil
	}
	*deviceId = devId

	active, err := IsDeviceStateActive(*deviceId)
	if err != nil {
		t.Fatal(err)
		return nil
	}
	if !active {
		t.Fatal("Timeout while waiting for device: %s to be active", *deviceId)
		return nil
	}

	device, err := GetDeviceById(*deviceId)
	if err != nil {
		t.Fatal(err)
		return nil
	}
	if len(device.NetworkPorts) < 3 {
		t.Fatal("All 3 ports doesnot exist for the created device: %s", device.GetId())
		return nil
	}

	return device
}

func CleanTestGateway(gatewayId string) error {
	TestApiClient := TestClient()
	_, _, err := TestApiClient.MetalGatewaysApi.
		DeleteMetalGateway(context.Background(), gatewayId).
		Include([]string{"ip_reservation"}).
		Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetalGatewaysApi.DeleteMetalGateway``: %v\n", err)
		return err
	}

	return nil
}
