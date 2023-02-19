package model

type As struct {
	Asn    int    `json:"asn,omitempty"`
	Name   string `json:"name,omitempty"`
	Route  string `json:"route,omitempty"`
	Domain string `json:"domain,omitempty"`
	Type   string `json:"type,omitempty"`
}

func NewAs() *As {
	return &As{}
}
