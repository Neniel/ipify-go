package ipify

const baseURL64 = "https://api64.ipify.org"

func GetIP64() (*IP, error) {
	return getIP(baseURL64)
}

func GetIP64AsString() (string, error) {
	return getIPAsString(baseURL64)
}

func GetIP64AsJSON() (string, error) {
	return getIPAsJSON(baseURL64)
}

func GetIP64AsJSONP(callback ...string) (string, error) {
	return getIPAsJSONP(baseURL64, callback...)
}
