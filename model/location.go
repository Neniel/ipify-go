package model

type Location struct {
	Country    string   `json:"country,omitempty"`
	Region     string   `json:"region,omitempty"`
	City       *string  `json:"city,omitempty"`
	Lat        *float64 `json:"lat,omitempty"`
	Lng        *float64 `json:"lng,omitempty"`
	PostalCode *string  `json:"postalCode,omitempty"`
	Timezone   string   `json:"timezone,omitempty"`
	GeonameId  *int     `json:"geonameId,omitempty"`
}

func NewLocation() *Location {
	return &Location{}
}
