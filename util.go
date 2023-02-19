package ipify

import (
	"encoding/json"
	"fmt"
	"io"
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
