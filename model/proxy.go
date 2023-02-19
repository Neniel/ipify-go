package model

type Proxy struct {
	Proxy bool `json:"proxy,omitempty"`
	Vpn   bool `json:"vpn,omitempty"`
	Tor   bool `json:"tor,omitempty"`
}

func NewProxy() *Proxy {
	return &Proxy{}
}
