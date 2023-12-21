package helper

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
	"time"

	"golang.org/x/exp/rand"

	"github.com/equinix/equinix-sdk-go/services/metalv1"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func TestClient() *metalv1.APIClient {
	configuration := metalv1.NewConfiguration()
	configuration.AddDefaultHeader("X-Auth-Token", os.Getenv("METAL_AUTH_TOKEN"))
	configuration.UserAgent = fmt.Sprintf("metal-cli/test-helper %s", configuration.UserAgent)
	// For debug purpose
	// configuration.Debug = true
	apiClient := metalv1.NewAPIClient(configuration)
	return apiClient
}

func CreateTestProject(t *testing.T, name string) *metalv1.Project {
	t.Helper()
	TestApiClient := TestClient()

	projectCreateFromRootInput := *metalv1.NewProjectCreateFromRootInput(name) // ProjectCreateFromRootInput | Project to create

	project, _, err := TestApiClient.ProjectsApi.CreateProject(context.Background()).ProjectCreateFromRootInput(projectCreateFromRootInput).Execute()
	if err != nil {
		t.Fatalf("Error when calling `ProjectsApi.CreateProject`: %v\n", err)
		return nil
	}

	t.Cleanup(func() {
		CleanTestProject(t, project.GetId())
	})

	return project
}

func CreateTestDevice(t *testing.T, projectId, name string) *metalv1.Device {
	t.Helper()
	TestApiClient := TestClient()

	hostname := name
	termination := time.Now().Add(1 * time.Hour)
	metroDeviceRequest := metalv1.CreateDeviceRequest{
		DeviceCreateInMetroInput: &metalv1.DeviceCreateInMetroInput{
			Metro:           "sv",
			Plan:            "m3.small.x86",
			OperatingSystem: "ubuntu_20_04",
			Hostname:        &hostname,
			TerminationTime: &termination,
		},
	}
	device, _, err := TestApiClient.DevicesApi.
		CreateDevice(context.Background(), projectId).
		CreateDeviceRequest(metroDeviceRequest).
		Execute()
	if err != nil {
		t.Fatalf("Error when calling `DevicesApi.CreateDevice`: %v\n", err)
	}

	t.Cleanup(func() {
		CleanTestDevice(t, device.GetId())
	})

	return device
}

func CreateTestVLAN(t *testing.T, projectId string) *metalv1.VirtualNetwork {
	TestApiClient := TestClient()
	t.Helper()

	metro := "sv"
	vlanCreateInput := metalv1.VirtualNetworkCreateInput{
		Metro: &metro,
	}
	vlan, _, err := TestApiClient.VLANsApi.
		CreateVirtualNetwork(context.Background(), projectId).
		VirtualNetworkCreateInput(vlanCreateInput).
		Execute()
	if err != nil {
		t.Fatalf("Error when calling `VLANsApi.CreateVirtualNetwork`: %v\n", err)
	}

	t.Cleanup(func() {
		CleanTestVlan(t, vlan.GetId())
	})

	return vlan
}

func CreateTestGateway(t *testing.T, projectId, vlanId string, privateIPv4SubnetSize *int32) *metalv1.MetalGateway {
	TestApiClient := TestClient()
	t.Helper()

	gatewayCreateInput := metalv1.CreateMetalGatewayRequest{
		MetalGatewayCreateInput: &metalv1.MetalGatewayCreateInput{
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
		t.Fatalf("Error when calling `MetalGatewaysApi.CreateMetalGateway`: %v\n", err)
	}

	return gateway.MetalGateway
}

func GetDeviceById(t *testing.T, deviceId string) (*metalv1.Device, error) {
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

func GetPortById(t *testing.T, portId string) (*metalv1.Port, error) {
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

	deviceActionInput := *metalv1.NewDeviceActionInput("power_off")
	TestApiClient := TestClient()

	_, err := TestApiClient.DevicesApi.PerformAction(context.Background(), deviceId).DeviceActionInput(deviceActionInput).Execute()
	if err != nil {
		return fmt.Errorf("Error when calling `DevicesApi.PerformAction``: %v\n", err)
	}
	return nil
}

func CleanTestDevice(t *testing.T, deviceId string) {
	t.Helper()

	_, err := IsDeviceStateActive(t, deviceId)
	if err != nil {
		t.Fatal(err)
	}

	TestApiClient := TestClient()
	resp, err := TestApiClient.DevicesApi.
		DeleteDevice(context.Background(), deviceId).
		ForceDelete(true).
		Execute()

	// When deleting a device:
	// - ignore 404 (likely already deleted)
	// - ignore 403 (likely failed provision)
	if err != nil &&
		resp.StatusCode != http.StatusNotFound &&
		resp.StatusCode != http.StatusForbidden {
		t.Fatalf("Error when calling `DevicesApi.DeleteDevice`` for %v: %v\n", deviceId, err)
	}
}

func CleanTestProject(t *testing.T, projectId string) {
	t.Helper()
	TestApiClient := TestClient()
	resp, err := TestApiClient.ProjectsApi.
		DeleteProject(context.Background(), projectId).
		Execute()
	if err != nil && resp.StatusCode != http.StatusNotFound {
		t.Fatalf("Error when calling `ProjectsApi.DeleteProject`` for %v: %v\n", projectId, err)
	}
}

func CreateTestIps(t *testing.T, projectId string, quantity int, ipType string) (string, error) {
	t.Helper()
	TestApiClient := TestClient()
	metro := "da"
	var tags []string
	var facility string

	req := &metalv1.IPReservationRequestInput{
		Metro:    &metro,
		Tags:     tags,
		Quantity: int32(quantity),
		Type:     ipType,
		Facility: &facility,
	}

	requestIPReservationRequest := &metalv1.RequestIPReservationRequest{
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
	resp, err := TestApiClient.IPAddressesApi.DeleteIPAddress(context.Background(), ipsId).Execute()
	if err != nil && resp.StatusCode != http.StatusNotFound {
		return fmt.Errorf("Error when calling `IPAddressesApi.DeleteIPAddress`` for %v: %v\n", ipsId, err)
	}
	return nil
}

func CreateTestVlanWithVxLan(t *testing.T, projectId string, Id int, desc string) (string, error) {
	t.Helper()
	TestApiClient := TestClient()
	virtualNetworkCreateInput := *metalv1.NewVirtualNetworkCreateInput()
	virtualNetworkCreateInput.SetDescription(desc)
	virtualNetworkCreateInput.SetMetro("da")
	virtualNetworkCreateInput.SetVxlan(int32(Id))

	vlanresp, _, err := TestApiClient.VLANsApi.CreateVirtualNetwork(context.Background(), projectId).VirtualNetworkCreateInput(virtualNetworkCreateInput).Execute()
	if err != nil {
		return "", fmt.Errorf("Error when calling `VLANsApi.CreateVirtualNetwork``: %v\n", err)
	}
	return vlanresp.GetId(), nil
}

func CleanTestVlan(t *testing.T, vlanId string) {
	t.Helper()
	TestApiClient := TestClient()
	_, resp, err := TestApiClient.VLANsApi.DeleteVirtualNetwork(context.Background(), vlanId).Execute()
	if err != nil && resp.StatusCode != http.StatusNotFound {
		t.Fatalf("Error when calling `VLANsApi.DeleteVirtualNetwork`` for %v: %v\n", vlanId, err)
	}
}

func UnAssignPortVlan(t *testing.T, portId, vlanId string) error {
	t.Helper()
	testClient := TestClient()
	_, _, err := testClient.PortsApi.
		UnassignPort(context.Background(), portId).
		PortAssignInput(metalv1.PortAssignInput{Vnid: &vlanId}).
		Execute()
	return err
}

//nolint:staticcheck
func SetupProjectAndDevice(t *testing.T, projPrefix string) (*metalv1.Project, *metalv1.Device) {
	t.Helper()
	projectName := projPrefix + GenerateRandomString(5)
	project := CreateTestProject(t, projectName)

	device := CreateTestDevice(t, project.GetId(), "metal-cli-test-device")

	active, err := IsDeviceStateActive(t, device.GetId())
	if err != nil {
		t.Fatal(err)
		return nil, nil
	}
	if !active {
		t.Fatalf("Timeout while waiting for device: %s to be active", device.GetId())
		return nil, nil
	}

	device, err = GetDeviceById(t, device.GetId())
	if err != nil {
		t.Fatal(err)
		return nil, nil
	}
	if len(device.NetworkPorts) < 3 {
		t.Fatalf("All 3 ports doesnot exist for the created device: %s", device.GetId())
		return nil, nil
	}

	return project, device
}

func CleanTestGateway(t *testing.T, gatewayId string) error {
	t.Helper()

	TestApiClient := TestClient()
	_, _, err := TestApiClient.MetalGatewaysApi.
		DeleteMetalGateway(context.Background(), gatewayId).
		Include([]string{"ip_reservation"}).
		Execute()
	if err != nil {
		return fmt.Errorf("Error when calling `MetalGatewaysApi.DeleteMetalGateway`` for %v: %v\n", gatewayId, err)
	}

	return nil
}

func CreateTestOrganization(t *testing.T, name string) *metalv1.Organization {
	TestApiClient := TestClient()

	organizationInput := metalv1.NewOrganizationInput()
	organizationInput.Name = &name
	organizationInput.Description = &name

	addr := metalv1.NewAddressWithDefaults()
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
		t.Fatalf("Error when calling `OrganizationsApi.CreateOrganization``: %v\n", err)
	}

	t.Cleanup(func() {
		CleanTestOrganization(t, resp.GetId())
	})

	return resp
}

func CleanTestOrganization(t *testing.T, orgId string) {
	TestApiClient := TestClient()

	resp, err := TestApiClient.OrganizationsApi.DeleteOrganization(context.Background(), orgId).Execute()
	if err != nil && resp.StatusCode != http.StatusNotFound {
		t.Fatalf("Error when calling `OrganizationsApi.DeleteOrganization`` for %v: %v\n", orgId, err)
	}
}

func CreateTestBgpEnableTest(projId string) error {
	TestApiClient := TestClient()

	bgpConfigRequestInput := *metalv1.NewBgpConfigRequestInput(int32(65000), metalv1.BgpConfigRequestInputDeploymentType("local"))

	_, err := TestApiClient.BGPApi.RequestBgpConfig(context.Background(), projId).BgpConfigRequestInput(bgpConfigRequestInput).Execute()
	if err != nil {
		return fmt.Errorf("Error when calling `BGPApi.RequestBgpConfig``: %v\n", err)
	}
	return nil
}

//nolint:staticcheck
func GenerateRandomString(length int) string {
	charSet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	random := rand.New(rand.NewSource(uint64(time.Now().UnixNano())))

	result := make([]byte, length)
	for i := range result {
		result[i] = charSet[random.Intn(len(charSet))]
	}
	return string(result)
}

func ExecuteAndCaptureOutput(t *testing.T, root *cobra.Command) []byte {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	err := root.Execute()

	ioErr := w.Close()
	if ioErr != nil {
		t.Logf("error while capturing command output: %v", ioErr)
	}

	os.Stdout = rescueStdout

	if err != nil {
		t.Fatal(err)
	}

	out, ioErr := io.ReadAll(r)
	if ioErr != nil {
		t.Logf("error while reading command output: %v", ioErr)
	}

	return out
}
