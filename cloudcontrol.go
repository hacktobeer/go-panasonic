// Package cloudcontrol package is a Go package to control
// Panasonic Comfort Cloud devices.
package cloudcontrol

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	pt "github.com/hacktobeer/go-panasonic/types"
	"github.com/m7shapan/njson"
	log "github.com/sirupsen/logrus"
)

// Client is a Panasonic Comfort Cloud client.
type Client struct {
	Utoken     string
	DeviceGUID string
	Server     string
}

// intPtr is a helper function that returns a pointer to an int.
func intPtr(i int) *int {
	return &i
}

// intPtr is a helper function that returns a pointer to an bool.
func boolPtr(b bool) *bool {
	return &b
}

// SetDevice sets the device GUID on the client.
func (c *Client) SetDevice(deviceGUID string) {
	c.DeviceGUID = deviceGUID
}

// setHeaders sets the required http request headers.
func (c *Client) setHeaders(req *http.Request) {
	if c.Utoken != "" {
		req.Header.Set("X-User-Authorization", c.Utoken)
	}
	req.Header.Set("X-APP-TYPE", "1")
	req.Header.Set("X-APP-VERSION", "1.9.0")
	req.Header.Set("User-Agent", "G-RAC")
	req.Header.Set("Accept", "application/json; charset=utf-8")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Connection", "Keep-Alive")

	log.Debugf("HTTP headers set to: %#v", req.Header)
}

// doPostRequest will send a HTTP POST request.
func (c *Client) doPostRequest(url string, postbody []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", c.Server+url, bytes.NewBuffer(postbody))
	c.setHeaders(req)

	log.Debugf("POST request URL: %#v\n", req.URL)
	log.Debugf("POST request body: %#v\n", string(postbody))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	log.Debugf("POST response body: %s", string(body))

	if resp.StatusCode > 200 {
		return body, fmt.Errorf("HTTP Error: %s", resp.Status)
	}

	return body, nil
}

// doGetRequest will send a HTTP GET request.
func (c *Client) doGetRequest(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", c.Server+url, nil)
	c.setHeaders(req)

	log.Debugf("GET request URL: %#v\n", req.URL)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	log.Debugf("GET response body: %s", string(body))

	if resp.StatusCode > 200 {
		return body, fmt.Errorf("HTTP Error: %s", resp.Status)
	}

	return body, nil
}

// NewClient creates a new Panasonic Comfort Cloud client.
func NewClient(server string) Client {
	client := Client{}
	if server != "" {
		client.Server = server
	} else {
		client.Server = pt.URLServer
	}

	log.Debugf("Created new client for %s", client.Server)

	return client
}

// ValidateSession checks if the session token is still valid.
func (c *Client) ValidateSession(token string) ([]byte, error) {
	c.Utoken = token
	body, err := c.doGetRequest(pt.URLValidate1)
	if err != nil {
		return body, fmt.Errorf("error: %v %s", err, body)
	}

	return body, nil
}

// CreateSession initialises a client session to Panasonic Comfort Cloud.
func (c *Client) CreateSession(username string, password string) ([]byte, error) {
	postBody, _ := json.Marshal(map[string]string{
		"language": "0",
		"loginId":  username,
		"password": password,
	})

	body, err := c.doPostRequest(pt.URLLogin, postBody)
	if err != nil {
		return nil, fmt.Errorf("error: %v %s", err, body)
	}

	session := pt.Session{}
	err = njson.Unmarshal([]byte(body), &session)
	if err != nil {
		log.Fatalf("unmarshal error %v: %s", err, body)
	}

	c.Utoken = session.Utoken

	return body, nil
}

// GetGroups gets all Panasonic Comfort Cloud groups associated to this account.
func (c *Client) GetGroups() (pt.Groups, error) {
	body, err := c.doGetRequest(pt.URLGroups)
	if err != nil {
		return pt.Groups{}, fmt.Errorf("error: %v %s", err, body)
	}
	groups := pt.Groups{}
	err = njson.Unmarshal([]byte(body), &groups)
	if err != nil {
		log.Fatalf("unmarshal error %v: %s", err, body)
	}

	return groups, nil
}

// ListDevices lists all available devices.
func (c *Client) ListDevices() ([]string, error) {
	available := []string{}
	groups, err := c.GetGroups()
	if err != nil {
		return nil, err
	}
	for _, group := range groups.Groups {
		for _, device := range group.Devices {
			available = append(available, device.DeviceGUID)
		}
	}

	return available, nil
}

// GetDeviceStatus gets all details for a specific device.
func (c *Client) GetDeviceStatus() (pt.Device, error) {
	body, err := c.doGetRequest(pt.URLDeviceStatus + url.QueryEscape(c.DeviceGUID))
	if err != nil {
		return pt.Device{}, fmt.Errorf("error: %v %s", err, body)
	}

	device := pt.Device{}
	err = njson.Unmarshal([]byte(body), &device)
	if err != nil {
		log.Fatalf("unmarshal error %v: %s", err, body)
	}

	return device, nil
}

// GetDeviceHistory will fetch historical device data from Panasonic.
func (c *Client) GetDeviceHistory(timeFrame int) (pt.History, error) {
	postBody, _ := json.Marshal(map[string]string{
		"dataMode":   fmt.Sprint(timeFrame),
		"date":       time.Now().Format("20060102"),
		"deviceGuid": c.DeviceGUID,
		"osTimezone": "+01:00",
	})

	body, err := c.doPostRequest(pt.URLHistory, postBody)
	if err != nil {
		return pt.History{}, fmt.Errorf("error: %v %s", err, body)
	}

	history := pt.History{}
	err = njson.Unmarshal([]byte(body), &history)
	if err != nil {
		log.Fatalf("unmarshal error %v: %s", err, body)
	}

	return history, nil
}

// control sends commands to the Panasonic cloud to control a device.
func (c *Client) control(command pt.Command) ([]byte, error) {
	postBody, _ := json.Marshal(command)

	log.Debugf("Command: %s", postBody)

	body, err := c.doPostRequest(pt.URLControl, postBody)
	if err != nil {
		return nil, fmt.Errorf("error: %v %s", err, body)
	}
	if string(body) != pt.SuccessResponse {
		return body, fmt.Errorf("error body: %v %s", err, body)
	}

	return body, nil
}

// SetTemperature will set the temperature for a device.
func (c *Client) SetTemperature(temperature float64) ([]byte, error) {
	command := pt.Command{
		DeviceGUID: c.DeviceGUID,
		Parameters: pt.DeviceControlParameters{
			TemperatureSet: &temperature,
		},
	}

	return c.control(command)
}

// TurnOn will switch the device on.
func (c *Client) TurnOn() ([]byte, error) {
	command := pt.Command{
		DeviceGUID: c.DeviceGUID,
		Parameters: pt.DeviceControlParameters{
			Operate: intPtr(1),
		},
	}

	return c.control(command)
}

// TurnOff will switch the device off.
func (c *Client) TurnOff() ([]byte, error) {
	command := pt.Command{
		DeviceGUID: c.DeviceGUID,
		Parameters: pt.DeviceControlParameters{
			Operate: intPtr(0),
		},
	}

	return c.control(command)
}

// SetMode will set the device to the requested AC mode.
func (c *Client) SetMode(mode int) ([]byte, error) {
	command := pt.Command{
		DeviceGUID: c.DeviceGUID,
		Parameters: pt.DeviceControlParameters{},
	}

	command.Parameters.OperationMode = intPtr(mode)

	return c.control(command)
}
