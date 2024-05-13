package helper

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/equinix/equinix-sdk-go/services/metalv1"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"golang.org/x/exp/rand"
)

const (
	ConsumerToken = ""
	URL           = "https://api.equinix.com/metal/v1/"
	Version       = "metal"
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

func CreateTestVLAN(t *testing.T, projectId, metro string) *metalv1.VirtualNetwork {
	TestApiClient := TestClient()
	t.Helper()

	if metro == "" {
		metro = "sv"
	}
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
	if gateway == nil {
		t.Fatal("Nil gateway returned. Error when calling `MetalGatewaysApi.CreateMetalGateway`")
	}

	t.Cleanup(func() {
		CleanTestGateway(t, gateway.MetalGateway.GetId())
	})

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
	return WaitForDeviceState(t, deviceId, metalv1.DEVICESTATE_ACTIVE)
}

func WaitForDeviceState(t *testing.T, deviceId string, states ...metalv1.DeviceState) (bool, error) {
	var device *metalv1.Device
	var err error
	t.Helper()
	predefinedTime := 900 * time.Second // Adjust this as needed
	retryInterval := 10 * time.Second   // Adjust this as needed
	startTime := time.Now()
	for time.Since(startTime) < predefinedTime {
		device, err = GetDeviceById(t, deviceId)
		if err != nil {
			return false, err
		}
		for _, state := range states {
			if device.GetState() == state {
				return true, nil
			}
		}

		// Sleep for the specified interval
		time.Sleep(retryInterval)
	}
	return false, fmt.Errorf("timed out waiting for device %v state %v to become one of %v", deviceId, device.GetState(), states)
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

	_, err := WaitForDeviceState(t, deviceId, metalv1.DEVICESTATE_ACTIVE, metalv1.DEVICESTATE_INACTIVE)
	// WaitForDeviceState doesn't return a response so we
	// look at the error message for now; we can revisit
	// this in a future refactoring
	if err != nil &&
		!strings.Contains(err.Error(), fmt.Sprint(http.StatusNotFound)) &&
		!strings.Contains(err.Error(), fmt.Sprint(http.StatusForbidden)) {
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

func CreateTestVlanWithVxLan(t *testing.T, projectId string, Id int, desc string) *metalv1.VirtualNetwork {
	t.Helper()
	TestApiClient := TestClient()
	virtualNetworkCreateInput := *metalv1.NewVirtualNetworkCreateInput()
	virtualNetworkCreateInput.SetDescription(desc)
	virtualNetworkCreateInput.SetMetro("da")
	virtualNetworkCreateInput.SetVxlan(int32(Id))

	vlanresp, _, err := TestApiClient.VLANsApi.CreateVirtualNetwork(context.Background(), projectId).VirtualNetworkCreateInput(virtualNetworkCreateInput).Execute()
	if err != nil {
		t.Fatalf("Error when calling `VLANsApi.CreateVirtualNetwork``: %v\n", err)
	}

	t.Cleanup(func() {
		CleanTestVlan(t, vlanresp.GetId())
	})

	return vlanresp
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

func AssertPortCmdOutput(t *testing.T, port *metalv1.Port, out, networkType string, bonded bool) {
	if !strings.Contains(out, port.GetId()) {
		t.Errorf("cmd output should contain ID of the port: %s", port.GetId())
	}

	if !strings.Contains(out, port.GetName()) {
		t.Errorf("cmd output should contain name of the port: %s", port.GetName())
	}

	if !strings.Contains(out, networkType) {
		t.Errorf("cmd output should contain type of the port: %s", string(port.GetNetworkType()))
	}

	if !strings.Contains(out, strconv.FormatBool(bonded)) {
		t.Errorf("cmd output should contain if port is bonded: %s", strconv.FormatBool(port.Data.GetBonded()))
	}
}

func CleanTestGateway(t *testing.T, gatewayId string) {
	t.Helper()

	TestApiClient := TestClient()
	_, resp, err := TestApiClient.MetalGatewaysApi.
		DeleteMetalGateway(context.Background(), gatewayId).
		Include([]string{"ip_reservation"}).
		Execute()
	if err != nil && resp.StatusCode != http.StatusNotFound {
		t.Fatalf("Error when calling `MetalGatewaysApi.DeleteMetalGateway`` for %v: %v\n", gatewayId, err)
	}

	if err := waitForVrfGatewayDeleted(TestApiClient, gatewayId, 5*time.Minute); err != nil {
		t.Fatal(err)
	}
}

func CreateTestInterConnection(t *testing.T, projectId, name string) *metalv1.Interconnection {
	t.Helper()
	TestApiClient := TestClient()

	createOrganizationInterconnectionRequest := metalv1.CreateOrganizationInterconnectionRequest{DedicatedPortCreateInput: metalv1.NewDedicatedPortCreateInput("da", name, "primary", metalv1.DedicatedPortCreateInputType("dedicated"))}

	resp, _, err := TestApiClient.InterconnectionsApi.CreateProjectInterconnection(context.Background(), projectId).CreateOrganizationInterconnectionRequest(createOrganizationInterconnectionRequest).Execute()
	if err != nil {
		t.Fatalf("Error when calling `InterconnectionsApi.CreateProjectInterconnection``: %v\n", err)
	}

	t.Cleanup(func() {
		CleanTestInterConnection(t, resp.GetId())
	})
	return resp
}

func CleanTestInterConnection(t *testing.T, connectionID string) {
	t.Helper()
	TestApiClient := TestClient()
	_, resp, err := TestApiClient.InterconnectionsApi.DeleteInterconnection(context.Background(), connectionID).Execute()
	if err != nil && resp.StatusCode != http.StatusNotFound {
		t.Fatalf("Error when calling `InterconnectionsApi.DeleteInterconnection``for %v : %v\n", connectionID, err)
	}
}

func GetInterconnPort(t *testing.T, connId string) string {
	t.Helper()
	TestApiClient := TestClient()
	resp, r, err := TestApiClient.InterconnectionsApi.ListInterconnectionPorts(context.Background(), connId).Execute()
	if err != nil && r.StatusCode != http.StatusNotFound {
		t.Fatalf("Error when calling `InterconnectionsApi.ListInterconnectionPorts``: %v\n", err)
	}

	return resp.Ports[len(resp.Ports)-1].GetId()
}

func CreateTestVirtualCircuit(t *testing.T, projectId, connId, portId, vlanId, name string) *metalv1.VirtualCircuit {
	t.Helper()
	TestApiClient := TestClient()

	vlanVCCreateInput := metalv1.NewVlanVirtualCircuitCreateInput(projectId)
	vlanVCCreateInput.SetVnid(vlanId)
	vlanVCCreateInput.SetSpeed("100")
	vlanVCCreateInput.SetNniVlan(1110)
	vlanVCCreateInput.SetName(name)

	virtualCircuitCreateInput := metalv1.VirtualCircuitCreateInput{VlanVirtualCircuitCreateInput: vlanVCCreateInput}

	resp, _, err := TestApiClient.InterconnectionsApi.CreateInterconnectionPortVirtualCircuit(context.Background(), connId, portId).VirtualCircuitCreateInput(virtualCircuitCreateInput).Execute()
	if err != nil {
		t.Fatalf("Error when calling `InterconnectionsApi.CreateInterconnectionPortVirtualCircuit``: %v\n", err)
	}

	t.Cleanup(func() {
		CleanTestVirtualCircuit(t, resp.VlanVirtualCircuit.GetId())
	})
	return resp
}

func CleanTestVirtualCircuit(t *testing.T, vcId string) {
	t.Helper()
	TestApiClient := TestClient()

	_, resp, err := TestApiClient.InterconnectionsApi.DeleteVirtualCircuit(context.Background(), vcId).Execute()
	if err != nil && resp.StatusCode != http.StatusNotFound {
		t.Fatalf("Error when calling `InterconnectionsApi.DeleteVirtualCircuit``for %v : %v\n", vcId, err)
	}
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

	bgpConfigRequestInput := *metalv1.NewBgpConfigRequestInput(65000, metalv1.BgpConfigRequestInputDeploymentType("local"))

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

	out, ioErr := io.ReadAll(r)
	if ioErr != nil {
		t.Logf("error while reading command output: %v", ioErr)
	}

	if err != nil {
		t.Fatalf("Failed to execute command with err: [%s]\n with output:\n%s", err.Error(), string(out))
	}

	return out
}

func CleanupInterconnectionVC(t *testing.T, connectionId string) {
	t.Helper()
	apiClient := TestClient()

	vcList, resp, err := apiClient.InterconnectionsApi.
		ListInterconnectionVirtualCircuits(context.Background(), connectionId).
		Execute()
	if err != nil && resp.StatusCode != http.StatusNotFound {
		t.Fatalf("Error when calling `InterconnectionsApi.ListInterconnectionVirtualCircuits`` for %v: %v\n", connectionId, err)
	}

	if vcList != nil && vcList.HasVirtualCircuits() {
		for _, vc := range vcList.GetVirtualCircuits() {
			_, resp, err := apiClient.InterconnectionsApi.
				DeleteVirtualCircuit(context.Background(), vc.VlanVirtualCircuit.GetId()).
				Execute()
			if err != nil && resp.StatusCode != http.StatusNotFound {
				t.Fatalf("Error when calling `InterconnectionsApi.DeleteVirtualCircuit`` for %v: %v\n", connectionId, err)
			}
		}
	}
}

func CleanupInterconnection(t *testing.T, connectionId string) {
	t.Helper()
	apiClient := TestClient()

	_, resp, err := apiClient.InterconnectionsApi.
		DeleteInterconnection(context.Background(), connectionId).
		Execute()
	if err != nil && resp.StatusCode != http.StatusNotFound {
		t.Fatalf("Error when calling `InterconnectionsApi.DeleteInterconnection`` for %v: %v\n", connectionId, err)
	}

	if err := waitForInterconnectionDeleted(apiClient, connectionId, 5*time.Minute); err != nil {
		t.Fatal(err)
	}

	CleanupInterconnectionVC(t, connectionId)
}

func waitForInterconnectionDeleted(apiClient *metalv1.APIClient, connId string, timeout time.Duration) error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), timeout)
	defer cancelFunc()

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return errors.New("Timeout while waiting for connection to be deleted")
		case <-ticker.C:
			conn, _, err := apiClient.InterconnectionsApi.GetInterconnection(context.Background(), connId).Execute()
			if err != nil {
				if strings.Contains(err.Error(), "Not Found") {
					return nil
				}
				return err
			}

			if conn == nil {
				return nil
			}

			fmt.Printf("Connection not deleted. Current status: [%s]", conn.GetStatus())
		}
	}
}

//nolint:staticcheck
func CreateTestVrfs(t *testing.T, projectId, name string, asn int64) *metalv1.Vrf {
	t.Helper()
	TestApiClient := TestClient()

	var IpRanges []string

	vrfCreateInput := *metalv1.NewVrfCreateInput("da", name)
	vrfCreateInput.SetLocalAsn(asn)
	IpRanges = append(IpRanges, "10.10.1.0/24")
	vrfCreateInput.SetIpRanges(IpRanges)

	resp, _, err := TestApiClient.VRFsApi.CreateVrf(context.Background(), projectId).VrfCreateInput(vrfCreateInput).Execute()
	if err != nil {
		t.Fatalf("Error when calling `VRFsApi.CreateVrf``: %v\n", err)
	}

	t.Cleanup(func() {
		CleanTestVrfs(t, resp.GetId())
	})

	return resp
}

//nolint:staticcheck
func CleanTestVrfs(t *testing.T, vrfId string) {
	t.Helper()
	TestApiClient := TestClient()

	resp, err := TestApiClient.VRFsApi.DeleteVrf(context.Background(), vrfId).Execute()
	if err != nil && resp.StatusCode != http.StatusNotFound {
		t.Fatalf("Error when calling `VRFsApi.DeleteVrf`` for ID: %v: with error: %v\n", vrfId, err)
	}
}

func CreateTestVrfIpRequest(t *testing.T, projectId, vrfId string) *metalv1.RequestIPReservation201Response {
	t.Helper()
	TestApiClient := TestClient()

	vrfReq := &metalv1.VrfIpReservationCreateInput{
		Type:    "vrf",
		Cidr:    int32(24),
		Network: "10.10.1.0",
		VrfId:   vrfId,
	}
	requestIPReservationRequest := &metalv1.RequestIPReservationRequest{
		VrfIpReservationCreateInput: vrfReq,
	}
	reservation, _, err := TestApiClient.IPAddressesApi.RequestIPReservation(context.Background(), projectId).RequestIPReservationRequest(*requestIPReservationRequest).Execute()
	if err != nil {
		t.Fatalf("Error when calling `IPAddressesApi.RequestIPReservation` for %v: %v\n", vrfId, err)
	}

	t.Cleanup(func() {
		CleanTestVrfIpRequest(t, reservation.VrfIpReservation.GetId())
	})

	return reservation
}

func CleanTestVrfIpRequest(t *testing.T, IPReservationId string) {
	t.Helper()
	TestApiClient := TestClient()
	resp, err := TestApiClient.IPAddressesApi.DeleteIPAddress(context.Background(), IPReservationId).Execute()
	if err != nil && resp.StatusCode != http.StatusNotFound {
		t.Fatalf("Error when calling `IPAddressesApi.DeleteIPAddress``for %v: %v\n", IPReservationId, err)
	}
}

func CreateTestVrfRoute(t *testing.T, vrfId string) *metalv1.VrfRoute {
	t.Helper()
	TestApiClient := TestClient()
	tags := []string{"foobar"}

	vrfRouteCreateInput := metalv1.VrfRouteCreateInput{
		Prefix:  "0.0.0.0/0",
		NextHop: "10.10.1.2",
		Tags:    tags,
	}

	vrfRoute, _, err := TestApiClient.VRFsApi.CreateVrfRoute(context.Background(), vrfId).VrfRouteCreateInput(vrfRouteCreateInput).Execute()
	if err != nil {
		t.Fatalf("Error when calling `VRFsApi.CreateVrfRoute`` for %v: %v\n", vrfId, err)
	}

	t.Cleanup(func() {
		CleanTestVrfIpRequest(t, vrfRoute.GetId())
	})

	return vrfRoute

}

func CleanTestVrfRoute(t *testing.T, vrfRouteId string) {
	t.Helper()

	TestApiClient := TestClient()

	_, resp, err := TestApiClient.VRFsApi.DeleteVrfRouteById(context.Background(), vrfRouteId).Execute()
	if err != nil && resp.StatusCode != http.StatusNotFound {
		t.Fatalf("Error when calling `VRFsApi.DeleteVrfRouteById`` for %v: %v\n", vrfRouteId, err)
	}
}

func CreateTestVrfGateway(t *testing.T, projectId, reservationId, vlanId string) *metalv1.VrfMetalGateway {
	t.Helper()
	TestApiClient := TestClient()
	includes := []string{"virtual_network", "ip_reservation"}

	gatewayCreateInput := metalv1.CreateMetalGatewayRequest{
		VrfMetalGatewayCreateInput: &metalv1.VrfMetalGatewayCreateInput{
			VirtualNetworkId: vlanId,
			IpReservationId:  reservationId,
		},
	}
	gateway, _, err := TestApiClient.MetalGatewaysApi.CreateMetalGateway(context.Background(), projectId).CreateMetalGatewayRequest(gatewayCreateInput).Include(includes).Execute()
	if err != nil {
		t.Fatalf("Error when calling `MetalGatewaysApi.CreateMetalGateway`: %v\n", err)
	}

	t.Cleanup(func() {
		CleanTestVrfGateway(t, gateway.VrfMetalGateway.GetId())
	})

	return gateway.VrfMetalGateway
}

func CleanTestVrfGateway(t *testing.T, gatewayId string) {
	t.Helper()

	TestApiClient := TestClient()
	includes := []string{"ip_reservation"}
	_, resp, err := TestApiClient.MetalGatewaysApi.DeleteMetalGateway(context.Background(), gatewayId).Include(includes).Execute()
	if err != nil && resp.StatusCode != http.StatusNotFound {
		t.Fatalf("Error when calling `MetalGatewaysApi.DeleteMetalGateway`` for %v: %v\n", gatewayId, err)
	}
}

func CreateTestBgpDynamicNeighbor(t *testing.T, gatewayId, iprange string, asn int64) *metalv1.BgpDynamicNeighbor {
	TestApiClient := TestClient()
	t.Helper()

	bgpNeighborCreateInput := metalv1.NewBgpDynamicNeighborCreateInput(iprange, asn)
	neighbor, _, err := TestApiClient.VRFsApi.
		CreateBgpDynamicNeighbor(context.Background(), gatewayId).
		BgpDynamicNeighborCreateInput(*bgpNeighborCreateInput).
		Include([]string{"created_by"}).
		Execute()
	if err != nil {
		t.Fatalf("Error when calling `VRFsApi.CreateBgpDynamicNeighbor`: %v\n", err)
	}

	t.Cleanup(func() {
		CleanTestBgpDynamicNeighbor(t, neighbor.GetId())
	})

	return neighbor
}

func CleanTestBgpDynamicNeighbor(t *testing.T, id string) {
	t.Helper()
	TestApiClient := TestClient()
	_, resp, err := TestApiClient.VRFsApi.DeleteBgpDynamicNeighborById(context.Background(), id).Include([]string{"created_by"}).Execute()
	if err != nil && resp.StatusCode != http.StatusNotFound {
		t.Fatalf("Error when calling `VRFsApi.DeleteBgpDynamicNeighborById``for %v: %v\n", id, err)
	}
}

func waitForVrfGatewayDeleted(apiClient *metalv1.APIClient, gatewayId string, timeout time.Duration) error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), timeout)
	defer cancelFunc()

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return errors.New("Timeout while waiting for gateway to be deleted")
		case <-ticker.C:
			gway, resp, err := apiClient.MetalGatewaysApi.FindMetalGatewayById(context.Background(), gatewayId).Execute()
			if err != nil {
				if strings.Contains(err.Error(), "Not Found") || resp.StatusCode == http.StatusNotFound {
					return nil
				}
				return err
			}

			if gway == nil {
				return nil
			}

			fmt.Printf("Gateway not deleted. Current status: [%s]", gway.VrfMetalGateway.GetId())
		}
	}
}
