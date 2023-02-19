package ipify

const baseURLv4 = "https://api.ipify.org"

func GetIPv4() (*IP, error) {
	return getIP(baseURLv4)
}

func GetIPv4AsString() (string, error) {
	return getIPAsString(baseURLv4)
}

func GetIPv4AsJSON() (string, error) {
	return getIPAsJSON(baseURLv4)
}

func GetIPv4AsJSONP(callback ...string) (string, error) {
	return getIPAsJSONP(baseURLv4, callback...)
}
