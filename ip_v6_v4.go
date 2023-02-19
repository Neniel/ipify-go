package ipify

import "github.com/Neniel/ipify-go/model"

const baseUrlV6orV4 = "https://api64.ipify.org"

func GetIP64() (*model.IP, error) {
	return getIP(baseUrlV6orV4)
}

func GetIP64AsString() (string, error) {
	return getIPAsString(baseUrlV6orV4)
}

func GetIP64AsJSON() (string, error) {
	return getIPAsJSON(baseUrlV6orV4)
}

func GetIP64AsJSONP(callback ...string) (string, error) {
	return getIPAsJSONP(baseUrlV6orV4, callback...)
}
