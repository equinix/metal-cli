package helper

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	openapiclient "github.com/equinix/equinix-sdk-go/services/metalv1"
	"github.com/pkg/errors"
)

func TestClient() *openapiclient.APIClient {
	configuration := openapiclient.NewConfiguration()
	configuration.AddDefaultHeader("X-Auth-Token", os.Getenv("METAL_AUTH_TOKEN"))
	// For debug purpose
	//configuration.Debug = true
	apiClient := openapiclient.NewAPIClient(configuration)
	return apiClient
}

func CreateTestProject(t *testing.T, name string) (string, error) {
	t.Helper()
	TestApiClient := TestClient()

	projectCreateFromRootInput := *openapiclient.NewProjectCreateFromRootInput(name) // ProjectCreateFromRootInput | Project to create

	projectResp, _, err := TestApiClient.ProjectsApi.CreateProject(context.Background()).ProjectCreateFromRootInput(projectCreateFromRootInput).Execute()
	if err != nil {
		return "", fmt.Errorf("Error when calling `ProjectsApi.CreateProject`: %v\n", err)
	}
	return projectResp.GetId(), nil
}

func CreateTestDevice(t *testing.T, projectId, name string) (string, error) {
	t.Helper()
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
		return "", fmt.Errorf("Error when calling `DevicesApi.CreateDevice`: %v\n", err)
	}
	return deviceResp.GetId(), nil
}

func CreateTestVLAN(t *testing.T, projectId string) (*openapiclient.VirtualNetwork, error) {
	TestApiClient := TestClient()
	t.Helper()

	metro := "sv"
	vlanCreateInput := openapiclient.VirtualNetworkCreateInput{
		Metro: &metro,
	}
	vlan, _, err := TestApiClient.VLANsApi.
		CreateVirtualNetwork(context.Background(), projectId).
		VirtualNetworkCreateInput(vlanCreateInput).
		Execute()
	if err != nil {
		return nil, fmt.Errorf("Error when calling `VLANsApi.CreateVirtualNetwork`: %v\n", err)
	}
	return vlan, nil
}

func CreateTestGateway(t *testing.T, projectId, vlanId string, privateIPv4SubnetSize *int32) (*openapiclient.MetalGateway, error) {
	TestApiClient := TestClient()
	t.Helper()

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
		return nil, fmt.Errorf("Error when calling `MetalGatewaysApi.CreateMetalGateway`: %v\n", err)
	}
	return gateway.MetalGateway, nil
}

func GetDeviceById(t *testing.T, deviceId string) (*openapiclient.Device, error) {
	TestApiClient := TestClient()
	t.Helper()

	includes := []string{"network_ports"}

	device, _, err := TestApiClient.DevicesApi.
		FindDeviceById(context.Background(), deviceId).
		Include(includes).
		Execute()
	if err != nil {
		return nil, fmt.Errorf("Error when calling `DevicesApi.FindDeviceById`: %v\n", err)
	}

	return device, nil
}

func GetPortById(t *testing.T, portId string) (*openapiclient.Port, error) {
	t.Helper()
	TestApiClient := TestClient()
	includes := []string{"virtual_network"}

	port, _, err := TestApiClient.PortsApi.
		FindPortById(context.Background(), portId).
		Include(includes).
		Execute()
	if err != nil {
		return nil, fmt.Errorf("Error when calling `PortsApi.FindPortById`: %v\n", err)
	}

	return port, nil
}

func IsDeviceStateActive(t *testing.T, deviceId string) (bool, error) {
	t.Helper()
	predefinedTime := 500 * time.Second // Adjust this as needed
	retryInterval := 10 * time.Second   // Adjust this as needed
	startTime := time.Now()
	for time.Since(startTime) < predefinedTime {
		device, err := GetDeviceById(t, deviceId)
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

func WaitForAttachVlanToPort(t *testing.T, portId string, attach bool) error {
	t.Helper()
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
			port, err := GetPortById(t, portId)
			if err != nil {
				return errors.Wrapf(err, "Error while fetching the port using ID: %s", portId)
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

func StopTestDevice(t *testing.T, deviceId string) error {
	t.Helper()

	deviceActionInput := *openapiclient.NewDeviceActionInput("power_off")
	TestApiClient := TestClient()

	_, err := TestApiClient.DevicesApi.PerformAction(context.Background(), deviceId).DeviceActionInput(deviceActionInput).Execute()
	if err != nil {
		return fmt.Errorf("Error when calling `DevicesApi.PerformAction``: %v\n", err)
	}
	return nil
}

func CleanTestDevice(t *testing.T, deviceId string) error {
	t.Helper()

	TestApiClient := TestClient()
	_, err := TestApiClient.DevicesApi.
		DeleteDevice(context.Background(), deviceId).
		ForceDelete(true).
		Execute()
	if err != nil {
		return fmt.Errorf("Error when calling `DevicesApi.DeleteDevice``: %v\n", err)
	}
	return nil
}

func CleanTestProject(t *testing.T, projectId string) error {
	t.Helper()
	TestApiClient := TestClient()
	_, err := TestApiClient.ProjectsApi.
		DeleteProject(context.Background(), projectId).
		Execute()
	if err != nil {
		return fmt.Errorf("Error when calling `ProjectsApi.DeleteProject``: %v\n", err)
	}
	return nil
}

func CreateTestIps(t *testing.T, projectId string, quantity int, ipType string) (string, error) {
	t.Helper()
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
		return "", fmt.Errorf("Error when calling `IPAddressesApi.FindIPReservations``: %v\n", err)
	}
	return ipsresp.IPReservation.GetId(), nil
}

func CleanTestIps(t *testing.T, ipsId string) error {
	t.Helper()
	TestApiClient := TestClient()
	_, err := TestApiClient.IPAddressesApi.DeleteIPAddress(context.Background(), ipsId).Execute()
	if err != nil {
		return fmt.Errorf("Error when calling `IPAddressesApi.DeleteIPAddress``: %v\n", err)
	}
	return nil
}

func CreateTestVlanWithVxLan(t *testing.T, projectId string, Id int, desc string) (string, error) {
	t.Helper()
	TestApiClient := TestClient()
	virtualNetworkCreateInput := *openapiclient.NewVirtualNetworkCreateInput()
	virtualNetworkCreateInput.SetDescription(desc)
	virtualNetworkCreateInput.SetMetro("da")
	virtualNetworkCreateInput.SetVxlan(int32(Id))

	vlanresp, _, err := TestApiClient.VLANsApi.CreateVirtualNetwork(context.Background(), projectId).VirtualNetworkCreateInput(virtualNetworkCreateInput).Execute()
	if err != nil {
		return "", fmt.Errorf("Error when calling `VLANsApi.CreateVirtualNetwork``: %v\n", err)
	}
	return vlanresp.GetId(), nil
}

func CleanTestVlan(t *testing.T, vlanId string) error {
	t.Helper()
	TestApiClient := TestClient()
	_, _, err := TestApiClient.VLANsApi.DeleteVirtualNetwork(context.Background(), vlanId).Execute()
	if err != nil {
		return fmt.Errorf("Error when calling `VLANsApi.DeleteVirtualNetwork``: %v\n", err)
	}

	return nil
}

func UnAssignPortVlan(t *testing.T, portId, vlanId string) error {
	t.Helper()
	testClient := TestClient()
	_, _, err := testClient.PortsApi.
		UnassignPort(context.Background(), portId).
		PortAssignInput(openapiclient.PortAssignInput{Vnid: &vlanId}).
		Execute()
	return err
}

func CleanupProjectAndDevice(t *testing.T, deviceId, projectId string) error {
	t.Helper()
	resp, err := IsDeviceStateActive(t, deviceId)
	if err == nil && resp {
		err = CleanTestDevice(t, deviceId)
		if err != nil {
			return err
		}
		err = CleanTestProject(t, projectId)
		if err != nil {
			return err
		}
	}

	return nil
}

//nolint:staticcheck
func SetupProjectAndDevice(t *testing.T, projectId, deviceId *string, projPrefix string) *openapiclient.Device {
	t.Helper()
	projectName := projPrefix + GenerateRandomString(5)
	projId, err := CreateTestProject(t, projectName)
	if err != nil {
		t.Fatal(err)
		return nil
	}
	*projectId = projId

	devId, err := CreateTestDevice(t, *projectId, "metal-cli-test-device")
	if err != nil {
		t.Fatal(err)
		return nil
	}
	*deviceId = devId

	active, err := IsDeviceStateActive(t, *deviceId)
	if err != nil {
		t.Fatal(err)
		return nil
	}
	if !active {
		t.Fatalf("Timeout while waiting for device: %s to be active", *deviceId)
		return nil
	}

	device, err := GetDeviceById(t, *deviceId)
	if err != nil {
		t.Fatal(err)
		return nil
	}
	if len(device.NetworkPorts) < 3 {
		t.Fatalf("All 3 ports doesnot exist for the created device: %s", device.GetId())
		return nil
	}

	return device
}

func CleanTestGateway(t *testing.T, gatewayId string) error {
	t.Helper()

	TestApiClient := TestClient()
	_, _, err := TestApiClient.MetalGatewaysApi.
		DeleteMetalGateway(context.Background(), gatewayId).
		Include([]string{"ip_reservation"}).
		Execute()
	if err != nil {
		return fmt.Errorf("Error when calling `MetalGatewaysApi.DeleteMetalGateway``: %v\n", err)
	}

	return nil
}

func CreateTestOrganization(name string) (string, error) {
	TestApiClient := TestClient()

	organizationInput := openapiclient.NewOrganizationInput()
	organizationInput.Name = &name
	organizationInput.Description = &name

	addr := openapiclient.NewAddressWithDefaults()
	addr.SetAddress("Boston")
	addr.SetCity("Boston")
	addr.SetCountry("US")
	addr.SetZipCode("02108")
	organizationInput.Address = addr

	// API spec says organization address.address is required,
	// but the address is not included by default
	defaultIncludes := []string{"address", "billing_address"}

	resp, _, err := TestApiClient.OrganizationsApi.CreateOrganization(context.Background()).OrganizationInput(*organizationInput).Include(defaultIncludes).Execute()
	if err != nil {
		return "", fmt.Errorf("Error when calling `OrganizationsApi.CreateOrganization``: %v\n", err)
	}

	return resp.GetId(), nil
}

func CleanTestOrganization(orgId string) error {
	TestApiClient := TestClient()

	_, err := TestApiClient.OrganizationsApi.DeleteOrganization(context.Background(), orgId).Execute()
	if err != nil {
		return fmt.Errorf("Error when calling `OrganizationsApi.DeleteOrganization``: %v\n", err)
	}

	return nil
}

func CreateTestBgpEnableTest(projId string) error {
	TestApiClient := TestClient()

	bgpConfigRequestInput := *openapiclient.NewBgpConfigRequestInput(int32(65000), openapiclient.BgpConfigRequestInputDeploymentType("local"))

	_, err := TestApiClient.BGPApi.RequestBgpConfig(context.Background(), projId).BgpConfigRequestInput(bgpConfigRequestInput).Execute()
	if err != nil {
		return fmt.Errorf("Error when calling `BGPApi.RequestBgpConfig``: %v\n", err)
	}
	return nil
}

//nolint:staticcheck
func GenerateRandomString(length int) string {
	// Calculate the number of bytes needed for the given string length
	numBytes := (length * 3) / 4

	// Create a byte slice to store the random bytes
	randomBytes := make([]byte, numBytes)

	// Read random bytes from the crypto/rand package
	_, err := rand.Read(randomBytes)
	if err != nil {
		return strconv.Itoa(int(time.Now().UnixNano()))
	}

	// Encode the random bytes to base64 to get a string
	randomString := base64.URLEncoding.EncodeToString(randomBytes)

	// Trim any padding characters
	randomString = randomString[:length]

	return randomString
}
