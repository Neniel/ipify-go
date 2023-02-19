package geo

import (
	"encoding/json"
	"fmt"
	"github.com/Neniel/ipify-go/model"
	"net/http"
	"strings"
)

const (
	baseUrl = "https://geo.ipify.org/api/v2/%s?%s=%s"

	EntityCountry = "country"
	EntityCity    = "city"
	EntityVpn     = "vpn"

	ParamApiKey = "apiKey"
)

func GetGeo(apiKey string, entities []string, optParams *model.OptParams) (*model.Geo, *model.GeoError) {
	url := buildUrl(apiKey, entities, optParams)

	response, err := http.Get(url)
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

	url := fmt.Sprintf(baseUrl, ParamApiKey, apiKey, strEntities)

	if optParams != nil {
		param := optParams.GetHighPriorityParam()
		if param != nil {
			url = fmt.Sprintf(url+"%s=%s", param.Key, param.Value)
		}

		if esc := optParams.GetParam(model.ParamEscapedUnicode); esc != nil {
			url = fmt.Sprintf(url+"%s=%s", esc.Key, esc.Value)
		}
	}

	return url
}
