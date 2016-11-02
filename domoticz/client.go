// Copyright (C) 2016 Nicolas Lamirault <nicolas.lamirault@gmail.com>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package domoticz

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/prometheus/common/log"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/nlamirault/domoticz_exporter/version"
)

const (
	acceptHeader = "application/json"
	mediaType    = "application/json"
)

var (
	userAgent = fmt.Sprintf("domoticz-exporter/%s", version.Version)
)

type Client struct {
	Endpoint string
	Username string
	Password string
}

func NewClient(ip string, username string, password string) (*Client, error) {
	url, err := url.Parse(fmt.Sprintf("http://%s", ip))
	if err != nil || url.Scheme != "http" {
		return nil, fmt.Errorf("Invalid Domoticz address: %s", err)
	}
	log.Debugf("Domoticz client creation")
	return &Client{
		Endpoint: url.String(),
		Username: username,
		Password: password,
	}, nil
}

func (c *Client) setupHeaders(request *http.Request) {
	request.Header.Add("Content-Type", mediaType)
	request.Header.Add("Accept", acceptHeader)
	request.Header.Add("User-Agent", userAgent)
	request.SetBasicAuth(c.Username, c.Password)
}

func (client *Client) GetAllDevices() (*DeviceResponse, error) {
	log.Infof("[Domoticz] Get all devices")
	resp, err := http.Get(fmt.Sprintf("%s/json.htm?type=devices&filter=all&used=true&order=Name", client.Endpoint))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var devices DeviceResponse
	dec := json.NewDecoder(bytes.NewBuffer(body))
	if err := dec.Decode(&devices); err != nil {
		return nil, err
	}
	return &devices, nil
}

func (client *Client) GetDevice(id string) (*DeviceResponse, error) {
	log.Infof("[Domoticz] Get device: %s", id)
	resp, err := http.Get(fmt.Sprintf("%s/json.htm?type=devices&rid=%s", client.Endpoint, id))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var device DeviceResponse
	dec := json.NewDecoder(bytes.NewBuffer(body))
	if err := dec.Decode(&device); err != nil {
		return nil, err
	}
	fmt.Printf("[Domoticz] Device: %s", device)
	return &device, nil
}
