package ipify

import "github.com/Neniel/ipify-go/model"

const baseUrlIPv6 = "https://api6.ipify.org"

func GetIPv6() (*model.IP, error) {
	return getIP(baseUrlIPv6)
}

func GetIPv6AsString() (string, error) {
	return getIPAsString(baseUrlIPv6)
}

func GetIPv6AsJSON() (string, error) {
	return getIPAsJSON(baseUrlIPv6)
}

func GetIPv6AsJSONP(callback ...string) (string, error) {
	return getIPAsJSONP(baseUrlIPv6, callback...)
}
