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

// This the package for the Domoticz API
// See: http://www.domoticz.com/wiki/Domoticz_API/JSON_URL%27s

package domoticz

type DeviceResponse struct {
	Status     string `json:"status"`
	Title      string `json:"title"`
	ActTime    int    `json:"ActTime"`
	ServerTime string `json:"ServerTime"`
	Sunrise    string `json:"Sunrise"`
	Sunset     string `json:"Sunset"`
	Result     []struct {
		AddjMulti         float64 `json:"AddjMulti"`
		AddjMulti2        float64 `json:"AddjMulti2"`
		AddjValue         float64 `json:"AddjValue"`
		AddjValue2        float64 `json:"AddjValue2"`
		BatteryLevel      float64 `json:"BatteryLevel"`
		CustomImage       int     `json:"CustomImage"`
		Data              string  `json:"Data"`
		DewPoint          string  `json:"DewPoint"`
		Favorite          int     `json:"Favorite"`
		HardwareID        int     `json:"HardwareID"`
		HardwareName      string  `json:"HardwareName"`
		HaveTimeout       bool    `json:"HaveTimeout"`
		Humidity          int     `json:"Humidity"`
		HumidityStatus    string  `json:"HumidityStatus"`
		ID                string  `json:"ID"`
		LastUpdate        string  `json:"LastUpdate"`
		Name              string  `json:"Name"`
		Notifications     string  `json:"Notifications"`
		PlanID            string  `json:"PlanID"`
		Protected         bool    `json:"Protected"`
		ShowNotifications bool    `json:"ShowNotifications"`
		SignalLevel       int     `json:"SignalLevel"`
		SubType           string  `json:"SubType"`
		Temp              float64 `json:"Temp"`
		Timers            string  `json:"Timers"`
		Type              string  `json:"Type"`
		TypeImg           string  `json:"TypeImg"`
		Unit              int     `json:"Unit"`
		Used              int     `json:"Used"`
		XOffset           string  `json:"XOffset"`
		YOffset           string  `json:"YOffset"`
		Idx               string  `json:"idx"`
	} `json:"result"`
}

type DevicesResponse struct {
	Status  string `json:"status"`
	Title   string `json:"title"`
	Devices string `json:"devices"`
}
