package ipify

import "github.com/Neniel/ipify-go/model"

const baseUrlIPv4 = "https://api.ipify.org"

// GetIPv4 returns a pointer to [model.IP] containing the IP v4 address. If an error happens, the pointer to [model.IP] will be null and the error
// will provide useful information.
func GetIPv4() (*model.IP, error) {
	return getIP(baseUrlIPv4)
}

// GetIPv4AsString returns a string representing the IP v4 address. If an error happens, the string will be empty and the error
// will provide useful information.
func GetIPv4AsString() (string, error) {
	return getIPAsString(baseUrlIPv4)
}

// GetIPv4AsJSON returns a JSON string representing the IP v4 address. If an error happens, the string will be empty and the error
// will provide useful information.
func GetIPv4AsJSON() (string, error) {
	return getIPAsJSON(baseUrlIPv4)
}

// GetIPv4AsJSONP returns a JSON string representing the IP v4 address. If an error happens, the string will be empty and the error
// will provide useful information. It can take an optional callback function name. If no callback function name is provided,
// it will return the JSONP string with the default `callback` function.
func GetIPv4AsJSONP(callback ...string) (string, error) {
	return getIPAsJSONP(baseUrlIPv4, callback...)
}
