package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"strings"

	cfg "RequestScript/config/env"
)

var config cfg.Config

func init() {
	config = cfg.NewViperConfig()
	fmt.Println("Request tool started")
}

func main() {
	host := config.GetString(`common.host`)
	hostname := config.GetString(`common.hostname`)

	// Loading settings from url
	fmt.Println("Getting settings:")
	var getSettingsUrl strings.Builder
	getSettingsUrl.WriteString(host)
	getSettingsUrl.WriteString("/cgi-bin/set_network_conf.cgi?")

	settings := new(Settings)
	getJson(getSettingsUrl.String(), settings)

	var url strings.Builder
	url.WriteString(host)
	url.WriteString("/cgi-bin/set_network_conf.cgi?")

	url.WriteString("_ant_conf_nettype=")
	url.WriteString(settings.conf_nettype)
	url.WriteString("&")

	url.WriteString("_ant_conf_hostname=")
	url.WriteString(hostname)
	url.WriteString("&")

	url.WriteString("_ant_conf_ipaddress=")
	url.WriteString(settings.conf_ipaddress)
	url.WriteString("&")

	url.WriteString("_ant_conf_netmask=")
	url.WriteString(settings.conf_netmask)
	url.WriteString("&")

	url.WriteString("_ant_conf_gateway=")
	url.WriteString(settings.conf_gateway)
	url.WriteString("&")

	url.WriteString("_ant_conf_dnsservers=")
	url.WriteString(settings.conf_dnsservers)

	fmt.Println("Execute request to change hostname")
	fmt.Println(url.String())

	req, _ := http.NewRequest("POST", url.String(), nil)
	req.Header.Add("cache-control", "no-cache")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}
