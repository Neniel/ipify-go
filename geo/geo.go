package geo

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/Neniel/ipify-go/model"
)

const (
	baseUrl = "https://geo.ipify.org/api/v2/%s?%s=%s"

	EntityCountry = "country"
	EntityCity    = "city"
	EntityVpn     = "vpn"

	ParamApiKey = "apiKey"
)

// GetGeo allows you to locate and identify website visitors by IP address.
//
// # Required parameters
//
//   - apiKey: Required. Get your personal API KEY on [My subscriptions]
//
// # Optional parameters
//
//   - entities: a slice of entities you would like to query: possible values are [EntityCountry], [EntityCity] and [EntityVpn]
//   - optParams: when provided, it will try to put the one with the highest priority on the request URL. For example:
//   - ipAddress: Optional. IPv4 or IPv6 to search location by. If the parameter is not specified, then it defaults to client request's public IP address.
//   - domain: Optional. Domain name to search location by.If the parameter is not specified, then 'ipAddress' will be used.
//   - email: Optional. Email address or domain name to search location by it's MX servers. If the parameter is not specified, then 'ipAddress' will be used.
//   - escapedUnicode: Optional. 1 - allows you to receive unescaped Unicode characters in the API response (default is to escape as \uXXXX). Acceptable values: 0 | 1. Default: 0
//
// # Return values
//
//   - If success: It will return a pointer to [model.Geo] and nil for [model.GeoError].
//   - If error:  It will return nil for [model.Geo] and a pointer to [model.GeoError].
//
// [My subscriptions]: https://geo.ipify.org/subscriptions
func GetGeo(apiKey string, entities []string, optParams *model.OptParams) (*model.Geo, *model.GeoError) {
	url := buildUrl(apiKey, entities, optParams)

	response, err := http.Get(url)
	defer func(body io.ReadCloser) {
		err := body.Close()
		if err != nil {
			log.Printf("Error when closing body: %v\n", err.Error())
		}
	}(response.Body)

	if err != nil {
		return nil, model.NewGeoError(-1, err.Error())
	}

	if response.StatusCode == http.StatusOK {
		geo := model.NewGeo()
		err = json.NewDecoder(response.Body).Decode(geo)
		if err != nil {
			return nil, model.NewGeoError(-1, err.Error())
		}
		return geo, nil
	}

	geoError := model.NewDefaultGeoError()

	err = json.NewDecoder(response.Body).Decode(geoError)
	if err != nil {
		return nil, model.NewGeoError(-1, err.Error())
	}

	return nil, geoError
}

func buildUrl(apiKey string, entities []string, optParams *model.OptParams) string {
	if len(entities) == 0 {
		entities = append(entities, EntityCountry)
	}
	strEntities := strings.Join(entities, ",")

	url := fmt.Sprintf(baseUrl, strEntities, ParamApiKey, apiKey)

	if optParams != nil {
		param := optParams.GetHighPriorityParam()
		if param != nil {
			url = fmt.Sprintf(url+"&%s=%s", param.Key, param.Value)
		}

		if esc := optParams.GetParam(model.ParamEscapedUnicode); esc != nil {
			url = fmt.Sprintf(url+"&%s=%s", esc.Key, esc.Value)
		}
	}

	return url
}
