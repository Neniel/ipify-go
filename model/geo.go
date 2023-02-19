package model

type Geo struct {
	Ip       string   `json:"ip,omitempty"`
	Location Location `json:"location,omitempty"`
	Domains  []string `json:"domains,omitempty"`
	As       As       `json:"as,omitempty"`
	Isp      string   `json:"isp,omitempty"`
	Proxy    Proxy    `json:"proxy,omitempty"`
}

func NewGeo() *Geo {
	return &Geo{}
}
