package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

//Config Struct for layout of config.json
type Config struct {
	Protocol       string `json:"Protocol"`
	UpdateInterval int    `json:"UpdateInterval"`
	Domain         struct {
		Name  string `json:"Name"`
		Token string `json:"Token"`
	} `json:"Domain"`
}

//deviceInfo Struct defines layout of JSON respone from IP API.
type deviceInfo struct {
	IP string `json:"ip"`
}

type device struct {
	IPv4 string
	IPv6 string
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

func setUpdateURL(cfg Config) string {
	var updateURL string
	var device device

	switch cfg.Protocol {
	case "ipv4":
		device.IPv4 = getDeviceInfo("v4")
		updateURL = fmt.Sprintf("https://www.duckdns.org/update?domains=%s&token=%s&ip=%s", cfg.Domain.Name, cfg.Domain.Token, device.IPv4)
	case "ipv6":
		device.IPv6 = getDeviceInfo("v6")
		updateURL = fmt.Sprintf("https://www.duckdns.org/update?domains=%s&token=%s&ipv6=%s", cfg.Domain.Name, cfg.Domain.Token, device.IPv6)
	case "both":
		device.IPv4 = getDeviceInfo("v4")
		device.IPv6 = getDeviceInfo("v6")
		updateURL = fmt.Sprintf("https://www.duckdns.org/update?domains=%s&token=%s&ip=%s&ipv6=%s", cfg.Domain.Name, cfg.Domain.Token, device.IPv4, device.IPv6)
	default:
		fmt.Println("Invalid Protocol defined.  Protocol should be either \"ipv4\", \"ipv6\", or \"both\".")
		os.Exit(1)
	}
	return updateURL
}

func updateDNS(updateURL string) {
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
			t := time.Now()
			fmt.Printf("%s Sucessfully updated IP.\n", t.Format("2006-01-02 15:04:05"))
		}
	}
}

func main() {
	var cfg Config
	loadConfig(&cfg)

	//Converts the update interval to milliseconds.
	updateInterval := cfg.UpdateInterval * 60 * 1000
	timer := time.Tick(time.Duration(updateInterval) * time.Millisecond)

	for range timer {
		updateURL := setUpdateURL(cfg)
		updateDNS(updateURL)
	}
}
