package ipify

const baseURLv6 = "https://api6.ipify.org"

func GetIPv6() (*IP, error) {
	return getIP(baseURLv6)
}

func GetIPv6AsString() (string, error) {
	return getIPAsString(baseURLv6)
}

func GetIPv6AsJSON() (string, error) {
	return getIPAsJSON(baseURLv6)
}

func GetIPv6AsJSONP(callback ...string) (string, error) {
	return getIPAsJSONP(baseURLv6, callback...)
}
