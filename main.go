package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//Config Struct for layout of config.json
type Config struct {
	UpdateInterval int `json:"UpdateInterval"`
	Domain         struct {
		Name  string `json:"Name"`
		Token string `json:"Token"`
	} `json:"Domain"`
}

//deviceInfo Struct defines layout of JSON respone from IP API.
type deviceInfo struct {
	IP string `json:"ip"`
}

func loadConfig(cfg *Config) {
	//Load a config.json file in the root directory of the application.
	f, err := os.Open("config.json")
	if err != nil {
		fmt.Println("Unable to read config file:", err)
		os.Exit(1)
	}
	defer f.Close()
	decoder := json.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		fmt.Println("Unable to decode JSON: ", err)
		os.Exit(1)
	}
}

func getDeviceInfo(protocol string) string {
	var apiURL string

	if protocol == "v6" {
		apiURL = "https://api6.ipify.org?format=json"
	} else if protocol == "v4" {
		apiURL = "https://api.ipify.org?format=json"
	}

	response, err := http.Get(apiURL)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	var device deviceInfo
	json.Unmarshal(responseData, &device)

	return device.IP
}

func main() {
	var cfg Config
	loadConfig(&cfg)

	deviceIPv4 := getDeviceInfo("v4")
	deviceIPv6 := getDeviceInfo("v6")

	updateURL := fmt.Sprintf("https://www.duckdns.org/update?domains=%s&token=%s&ip=%s&ipv6=%s", cfg.Domain.Name, cfg.Domain.Token, deviceIPv4, deviceIPv6)
	//fmt.Println(updateURL)

	updateResponse, err := http.Get(updateURL)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	updateResponseData, err := ioutil.ReadAll(updateResponse.Body)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	} else {
		if string(updateResponseData) != "OK" {
			fmt.Println("Unable to update IP, please check config")
			os.Exit(1)
		} else {
			fmt.Println("Sucessfully updated IP.")
		}
	}
}
