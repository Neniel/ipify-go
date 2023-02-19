package ipify

import "github.com/Neniel/ipify-go/model"

const baseUrlIPv4 = "https://api.ipify.org"

func GetIPv4() (*model.IP, error) {
	return getIP(baseUrlIPv4)
}

func GetIPv4AsString() (string, error) {
	return getIPAsString(baseUrlIPv4)
}

func GetIPv4AsJSON() (string, error) {
	return getIPAsJSON(baseUrlIPv4)
}

func GetIPv4AsJSONP(callback ...string) (string, error) {
	return getIPAsJSONP(baseUrlIPv4, callback...)
}
