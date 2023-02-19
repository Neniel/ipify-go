// Package ipify implements a simple library for using the services provided by ipify.
//
// Using ipify is ridiculously simple. You have three options. You can get your public IP directly (in plain text),
// you can get your public IP in JSON format, or you can get your public IP information in JSONP format (useful for Javascript developers).
//
// Visit [ipify official website] for more information.
//
// [ipify official website]: https://www.ipify.org/
package ipify

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Neniel/ipify-go/model"
)

const (
	jsonFormat              = "?format=json"
	jsonpFormat             = "?format=jsonp"
	jsonpFormatWithCallback = "?format=jsonp&callback="
)

func getIPAsString(url string) (string, error) {
	response, err := http.Get(url)
	defer func(body io.ReadCloser) {
		err := body.Close()
		if err != nil {
			log.Printf("Error when closing body: %v\n", err.Error())
		}
	}(response.Body)

	if err != nil {
		return "", err
	}

	if response.StatusCode == http.StatusOK {
		bs, err := io.ReadAll(response.Body)
		if err != nil {
			return "", err
		}
		return string(bs), nil
	}

	return "", fmt.Errorf(response.Status)
}

func getIPAsJSON(url string) (string, error) {
	url = fmt.Sprintf(url+"%s", jsonFormat)
	return getIPAsString(url)
}

func getIPAsJSONP(url string, callback ...string) (string, error) {
	if len(callback) == 0 {
		url = fmt.Sprintf(url+"%s", jsonpFormat)
	}

	if len(callback) > 1 {
		return "", fmt.Errorf("only one callback name is allowed")
	}

	if len(callback) == 1 {
		url = fmt.Sprintf(url+"%s%s", jsonpFormatWithCallback, callback[0])
	}

	return getIPAsString(url)
}

func getIP(url string) (*model.IP, error) {
	ipString, err := getIPAsJSON(url)
	if err != nil {
		return nil, err
	}

	ip := model.NewIP()
	err = json.Unmarshal([]byte(ipString), ip)
	if err != nil {
		return nil, err
	}

	return ip, nil
}
