package model

type Geo struct {
	Ip                      string            `json:"ip,omitempty"`
	Location                *Location         `json:"location,omitempty"`
	Domains                 []string          `json:"domains,omitempty"`
	AutonomousSystem        *AutonomousSystem `json:"as,omitempty"`
	InternetServiceProvider string            `json:"isp,omitempty"`
	Proxy                   *Proxy            `json:"proxy,omitempty"`
}

func NewGeo() *Geo {
	return &Geo{}
}
