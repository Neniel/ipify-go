package model

type GeoError struct {
	Code     int    `json:"code"`
	Messages string `json:"messages"`
	Another  string `json:"another"`
}

func NewDefaultGeoError() *GeoError {
	return &GeoError{}
}

func NewGeoError(code int, messages string) *GeoError {
	return &GeoError{
		Code:     code,
		Messages: messages,
	}
}
