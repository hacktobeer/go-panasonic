package cloudcontrol_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hacktobeer/gopanasonic/cloudcontrol"
	pt "github.com/hacktobeer/gopanasonic/types"
)

//var server *httptest.Server
var (
	groupsBody  = `{"iaqStatus":{"statusCode":200},"groupCount":1,"groupList":[{"groupId":112867,"groupName":"My House","deviceList":[{"deviceGuid":"CZ-CAPWFC1+B8B7F1B3E326","deviceType":"4","deviceName":"Alaior-home","permission":3,"deviceModuleNumber":"S-125PU2E5B","deviceHashGuid":"f609023332bbeee157a5b868fe80b9fb14a1d883938c1836003796332150db16","summerHouse":0,"iAutoX":false,"nanoe":true,"autoMode":true,"heatMode":true,"fanMode":false,"dryMode":true,"coolMode":true,"ecoNavi":false,"powerfulMode":true,"quietMode":true,"airSwingLR":true,"ecoFunction":0,"temperatureUnit":0,"modeAvlList":{"autoMode":1,"fanMode":1},"autoTempMax":27,"autoTempMin":17,"dryTempMax":30,"dryTempMin":18,"coolTempMax":30,"coolTempMin":18,"heatTempMax":30,"heatTempMin":16,"fanSpeedMode":5,"fanDirectionMode":5,"parameters":{"operate":1,"operationMode":0,"temperatureSet":19.5,"fanSpeed":0,"fanAutoMode":1,"airSwingLR":2,"airSwingUD":3,"ecoMode":0,"ecoNavi":0,"nanoe":1,"iAuto":0,"actualNanoe":1,"airDirection":3,"ecoFunctionData":0}}]}]}`
	controlBody = `{"result":0}`
)

func serverMock() *httptest.Server {
	handler := http.NewServeMux()
	handler.HandleFunc(pt.URLLogin, sessionMock)
	handler.HandleFunc(pt.URLGroups, groupsMock)
	handler.HandleFunc(pt.URLControl, controlMock)

	srv := httptest.NewServer(handler)

	return srv
}

func sessionMock(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte(`{"uToken":"token12345","language":0,"result":0}`))
}

func groupsMock(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte(groupsBody))
}

func controlMock(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte(controlBody))
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
	server := serverMock()
	defer server.Close()

	var client cloudcontrol.Client
	client.CreateSession("", "", "", server.URL)
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
	server := serverMock()
	defer server.Close()

	var client cloudcontrol.Client
	client.CreateSession("", "", "", server.URL)
	groups, _ := client.GetGroups()
	want := "My House"
	got := groups.Groups[0].GroupName
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("TestGetGroups() mismatch (-want +got):\n%s", diff)
	}
}

func TestCreateSessionCustomToken(t *testing.T) {
	username := ""
	password := ""
	token := "token12345"

	server := serverMock()
	defer server.Close()

	var client cloudcontrol.Client
	body, _ := client.CreateSession(token, username, password, server.URL)
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

	server := serverMock()
	defer server.Close()

	var client cloudcontrol.Client
	client.CreateSession(token, username, password, server.URL)

	got := client.Utoken
	want := "token12345"
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("TestCreateSession() token mismatch (-want +got):\n%s", diff)
	}

	got = client.Server
	want = server.URL
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("TestCreateSession() server URL mismatch (-want +got):\n%s", diff)
	}
}
