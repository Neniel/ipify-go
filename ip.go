package ipify

type IP struct {
	IP string `json:"ip"`
}

func NewIP() *IP {
	return &IP{}
}
