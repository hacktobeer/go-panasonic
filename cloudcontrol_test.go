package cloudcontrol_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hacktobeer/gopanasonic/cloudcontrol"
	pt "github.com/hacktobeer/gopanasonic/types"
)

// Example on how to use this package:
func Example() {
	// Create a new Panasonic Comfort Cloud client
	client := cloudcontrol.NewClient("")
	// Initiate a session with your username and password
	client.CreateSession("", "username", "password")
	// List the available devices in your account
	devices, _ := client.ListDevices()
	// Set the device we want to control
	client.SetDevice(devices[0])
	// Show the detailed device status
	status, _ := client.GetDeviceStatus()
	// Show the inside temperature measured by the device
	fmt.Println(status.Parameters.InsideTemperature)
	// Set the temperature on the device
	client.SetTemperature(19.5)
}

//var server *httptest.Server
var (
	client      cloudcontrol.Client
	sessionBody = `{"uToken":"token12345","language":0,"result":0}`
	groupsBody  = `{"iaqStatus":{"statusCode":200},"groupCount":1,"groupList":[{"groupId":112867,"groupName":"My House","deviceList":[{"deviceGuid":"CZ-CAPWFC1+B8B7F1B3E326","deviceType":"4","deviceName":"Alaior-home","permission":3,"deviceModuleNumber":"S-125PU2E5B","deviceHashGuid":"f609023332bbeee157a5b868fe80b9fb14a1d883938c1836003796332150db16","summerHouse":0,"iAutoX":false,"nanoe":true,"autoMode":true,"heatMode":true,"fanMode":false,"dryMode":true,"coolMode":true,"ecoNavi":false,"powerfulMode":true,"quietMode":true,"airSwingLR":true,"ecoFunction":0,"temperatureUnit":0,"modeAvlList":{"autoMode":1,"fanMode":1},"autoTempMax":27,"autoTempMin":17,"dryTempMax":30,"dryTempMin":18,"coolTempMax":30,"coolTempMin":18,"heatTempMax":30,"heatTempMin":16,"fanSpeedMode":5,"fanDirectionMode":5,"parameters":{"operate":1,"operationMode":0,"temperatureSet":19.5,"fanSpeed":0,"fanAutoMode":1,"airSwingLR":2,"airSwingUD":3,"ecoMode":0,"ecoNavi":0,"nanoe":1,"iAuto":0,"actualNanoe":1,"airDirection":3,"ecoFunctionData":0}}]}]}`
	controlBody = `{"result":0}`
)

func TestMain(m *testing.M) {
	server := serverMock()
	defer server.Close()

	client = cloudcontrol.NewClient(server.URL)

	os.Exit(m.Run())
}

func serverMock() *httptest.Server {
	handler := http.NewServeMux()
	handler.HandleFunc(pt.URLLogin, sessionMock)
	handler.HandleFunc(pt.URLGroups, groupsMock)
	handler.HandleFunc(pt.URLControl, controlMock)

	srv := httptest.NewServer(handler)

	return srv
}

func sessionMock(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte(sessionBody))
}

func groupsMock(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte(groupsBody))
}

func controlMock(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte(controlBody))
}

func TestNewClient(t *testing.T) {
	cases := []struct {
		input string
		want  string
	}{
		{
			input: "http:/customserver.com",
			want:  "http:/customserver.com",
		},
		{
			input: "",
			want:  pt.URLServer,
		},
	}
	for _, c := range cases {
		client := cloudcontrol.NewClient(c.input)
		got := client.Server
		if diff := cmp.Diff(c.want, got); diff != "" {
			t.Errorf("TestNewClient() mismatch (-want +got):\n%s", diff)
		}
	}
}

func TestSetDevice(t *testing.T) {
	device := "device12345"

	var client cloudcontrol.Client
	client.SetDevice(device)

	want := device
	got := client.DeviceGUID
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("TestSetDevice() mismatch (-want +got):\n%s", diff)
	}
}

func TestTurnOn(t *testing.T) {
	client.CreateSession("", "", "")
	body, err := client.TurnOn()
	if err != nil {
		t.Errorf("TestTurnOn() returned an error: %v", err)
	}
	want := string(pt.SuccessResponse)
	got := string(body)
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("TestTurnOn() mismatch (-want +got):\n%s", diff)
	}
}

func TestGetGroups(t *testing.T) {
	client.CreateSession("", "", "")
	groups, _ := client.GetGroups()
	want := "My House"
	got := groups.Groups[0].GroupName
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("TestGetGroups() mismatch (-want +got):\n%s", diff)
	}
	if len(groups.Groups[0].Devices) != 1 {
		t.Errorf("TestGetGroups() mismatch Devices, want 1, got %d", len(groups.Groups[0].Devices))
	}
}

func TestCreateSessionCustomToken(t *testing.T) {
	username := ""
	password := ""
	token := "token12345"

	body, _ := client.CreateSession(token, username, password)
	if body != nil {
		t.Error("TestCreateSessionCustomToken() got non-nil body")
	}

	got := client.Utoken
	want := "token12345"
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("TestCreateSessionCustomToken() token mismatch (-want +got):\n%s", diff)
	}
}

func TestCreateSession(t *testing.T) {
	username := "test@test.com"
	password := "secret1234"
	token := ""

	client.CreateSession(token, username, password)

	got := client.Utoken
	want := "token12345"
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("TestCreateSession() token mismatch (-want +got):\n%s", diff)
	}
}
