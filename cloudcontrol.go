package cloudcontrol

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	pt "github.com/hacktobeer/gopanasonic/types"
	"github.com/m7shapan/njson"
)

// Client is a Cloud Control client
type Client struct {
	Utoken     string
	DeviceGUID string
	Server     string
}

// intPtr is a helper function that returns a pointer to an int
func intPtr(i int) *int {
	return &i
}

// intPtr is a helper function that returns a pointer to an bool
func boolPtr(b bool) *bool {
	return &b
}

// SetDevice sets the device GUID on the client
func (c *Client) SetDevice(deviceGUID string) {
	c.DeviceGUID = deviceGUID
}

func (c *Client) setHeaders(req *http.Request) {
	if c.Utoken != "" {
		req.Header.Set("X-User-Authorization", c.Utoken)
	}
	req.Header.Set("X-APP-TYPE", "1")
	req.Header.Set("X-APP-VERSION", "1.9.0")
	req.Header.Set("User-Agent", "G-RAC")
	req.Header.Set("Accept", "application/json; charset=utf-8")
	req.Header.Set("Content-Type", "application/json")
}

func (c *Client) doPostRequest(url string, postbody []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", c.Server+url, bytes.NewBuffer(postbody))
	c.setHeaders(req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode > 200 {
		return body, fmt.Errorf("HTTP Error: %s", resp.Status)
	}

	return body, nil
}

func (c *Client) doGetRequest(url string) ([]byte, error) {
	log.Println(url)
	req, err := http.NewRequest("GET", c.Server+url, nil)
	c.setHeaders(req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode > 200 {
		return body, fmt.Errorf("HTTP Error: %s", resp.Status)
	}

	return body, nil
}

// CreateSession creates a client connection to Panasonic Cloud Control
func (c *Client) CreateSession(token string, username string, password string, server string) error {
	if username == "" {
		c.Utoken = token
	}
	if server != "" {
		c.Server = server
	} else {
		c.Server = "https://accsmart.panasonic.com/"
	}

	postBody, _ := json.Marshal(map[string]string{
		"language": "0",
		"loginId":  username,
		"password": password,
	})

	body, err := c.doPostRequest("auth/login", postBody)
	if err != nil {
		return fmt.Errorf("Error: %v %s", err, body)
	}

	session := pt.Session{}
	err = njson.Unmarshal([]byte(body), &session)
	if err != nil {
		log.Fatalf("unmarshal error %v: %s", err, body)
	}

	c.Utoken = session.Utoken

	return nil
}

// GetGroups gets all Panasonic Control groups associated to this account
func (c *Client) GetGroups() (pt.Groups, error) {
	body, err := c.doGetRequest("/device/group")
	if err != nil {
		return pt.Groups{}, fmt.Errorf("Error: %v %s", err, body)
	}
	groups := pt.Groups{}
	err = njson.Unmarshal([]byte(body), &groups)
	if err != nil {
		log.Fatalf("unmarshal error %v: %s", err, body)
	}

	return groups, nil
}

// ListDevices lists all available devices
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

// GetDevice gets all details on a specific device
func (c *Client) GetDevice(deviceGUID string) (pt.Device, error) {
	body, err := c.doGetRequest("/deviceStatus/now/" + url.QueryEscape(deviceGUID))
	if err != nil {
		return pt.Device{}, fmt.Errorf("Error: %v %s", err, body)
	}

	device := pt.Device{}
	err = njson.Unmarshal([]byte(body), &device)
	if err != nil {
		log.Fatalf("unmarshal error %v: %s", err, body)
	}

	return device, nil
}

// GetDeviceHistory will fetch historical device data from Panasonic
func (c *Client) GetDeviceHistory(timeFrame int) (pt.History, error) {
	postBody, _ := json.Marshal(map[string]string{
		"dataMode":   fmt.Sprint(timeFrame),
		"date":       time.Now().Format("20060102"),
		"deviceGuid": c.DeviceGUID,
		"osTimezone": "+01:00",
	})

	body, err := c.doPostRequest("deviceHistoryData", postBody)
	if err != nil {
		return pt.History{}, fmt.Errorf("Error: %v %s", err, body)
	}

	history := pt.History{}
	err = njson.Unmarshal([]byte(body), &history)
	if err != nil {
		log.Fatalf("unmarshal error %v: %s", err, body)
	}

	return history, nil
}

// control sends commands to the Panasonic cloud to control a device
func (c *Client) control(command pt.Command) error {
	postBody, _ := json.Marshal(command)

	log.Println("JSON to be sent:")
	log.Println(string(postBody))

	body, err := c.doPostRequest("deviceStatus/control", postBody)
	if err != nil {
		return fmt.Errorf("Error: %v %s", err, body)
	}

	log.Println(string(body))

	return nil
}

// SetTemperature will turn the Panasonic device off
func (c *Client) SetTemperature(temperature float64) error {
	command := pt.Command{
		DeviceGUID: c.DeviceGUID,
		Parameters: pt.CommandDetails{
			TemperatureSet: &temperature,
		},
	}

	return c.control(command)
}

// TurnOn will switch the device on or off
func (c *Client) TurnOn() error {
	command := pt.Command{
		DeviceGUID: c.DeviceGUID,
		Parameters: pt.CommandDetails{
			Operate: intPtr(1),
		},
	}

	return c.control(command)
}

// TurnOff will switch the device on or off
func (c *Client) TurnOff() error {
	command := pt.Command{
		DeviceGUID: c.DeviceGUID,
		Parameters: pt.CommandDetails{
			Operate: intPtr(0),
		},
	}

	return c.control(command)
}

// SetMode will set the device to the requested AC mode
func (c *Client) SetMode(mode int) error {
	command := pt.Command{
		DeviceGUID: c.DeviceGUID,
		Parameters: pt.CommandDetails{},
	}

	command.Parameters.OperationMode = intPtr(mode)

	return c.control(command)
}
