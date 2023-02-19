package geo

import (
	"reflect"
	"testing"

	"github.com/Neniel/ipify-go/model"
)

func TestGetGeo(t *testing.T) {
	type args struct {
		apiKey    string
		entities  []string
		optParams *model.OptParams
	}
	tests := []struct {
		name  string
		args  args
		want  *model.Geo
		want1 *model.GeoError
	}{
		{
			name: "Make request with no API key should fail",
			args: args{
				apiKey:    "",
				entities:  nil,
				optParams: nil,
			},
			want:  nil,
			want1: model.NewGeoError(403, "Access restricted. Check credits balance or enter the correct API key."),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := GetGeo(tt.args.apiKey, tt.args.entities, tt.args.optParams)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetGeo() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GetGeo() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_buildUrl(t *testing.T) {
	type args struct {
		apiKey    string
		entities  []string
		optParams *model.OptParams
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Should get URL with country entity",
			args: args{
				apiKey:    "MY_API_KEY",
				entities:  []string{},
				optParams: nil,
			},
			want: "https://geo.ipify.org/api/v2/country?apiKey=MY_API_KEY",
		},
		{
			name: "Should get URL with country entity by ipAddress",
			args: args{
				apiKey:    "MY_API_KEY",
				entities:  []string{},
				optParams: model.NewOptParams().Add(model.ParamIpAddress, "127.0.0.1"),
			},
			want: "https://geo.ipify.org/api/v2/country?apiKey=MY_API_KEY&ipAddress=127.0.0.1",
		},
		{
			name: "Should get URL with country entity by domain",
			args: args{
				apiKey:    "MY_API_KEY",
				entities:  []string{},
				optParams: model.NewOptParams().Add(model.ParamDomain, "test.net"),
			},
			want: "https://geo.ipify.org/api/v2/country?apiKey=MY_API_KEY&domain=test.net",
		},
		{
			name: "Should get URL with country entity by email",
			args: args{
				apiKey:    "MY_API_KEY",
				entities:  []string{},
				optParams: model.NewOptParams().Add(model.ParamEmail, "test@test.net"),
			},
			want: "https://geo.ipify.org/api/v2/country?apiKey=MY_API_KEY&email=test@test.net",
		},
		{
			name: "Should get URL with country entity by email because it is the param with highest priority among ipAddress, domain and email",
			args: args{
				apiKey:    "MY_API_KEY",
				entities:  []string{},
				optParams: model.NewOptParams().Add(model.ParamEmail, "test@test.net").Add(model.ParamDomain, "test.net").Add(model.ParamIpAddress, "127.0.0.1"),
			},
			want: "https://geo.ipify.org/api/v2/country?apiKey=MY_API_KEY&email=test@test.net",
		},
		{
			name: "Should get URL with country entity by email because it is the param with highest priority among ipAddress and email",
			args: args{
				apiKey:    "MY_API_KEY",
				entities:  []string{},
				optParams: model.NewOptParams().Add(model.ParamEmail, "test@test.net").Add(model.ParamIpAddress, "127.0.0.1"),
			},
			want: "https://geo.ipify.org/api/v2/country?apiKey=MY_API_KEY&email=test@test.net",
		},
		{
			name: "Should get URL with country entity by email because it is the param with highest priority among email and domain",
			args: args{
				apiKey:    "MY_API_KEY",
				entities:  []string{},
				optParams: model.NewOptParams().Add(model.ParamEmail, "test@test.net").Add(model.ParamDomain, "test.net"),
			},
			want: "https://geo.ipify.org/api/v2/country?apiKey=MY_API_KEY&email=test@test.net",
		},
		{
			name: "Should get URL with country entity",
			args: args{
				apiKey:    "MY_API_KEY",
				entities:  []string{EntityCountry},
				optParams: nil,
			},
			want: "https://geo.ipify.org/api/v2/country?apiKey=MY_API_KEY",
		},
		{
			name: "Should get URL with country and city entities",
			args: args{
				apiKey:    "MY_API_KEY",
				entities:  []string{EntityCountry, EntityCity},
				optParams: nil,
			},
			want: "https://geo.ipify.org/api/v2/country,city?apiKey=MY_API_KEY",
		},
		{
			name: "Should get URL with country, city and vpn entities",
			args: args{
				apiKey:    "MY_API_KEY",
				entities:  []string{EntityCountry, EntityCity, EntityVpn},
				optParams: nil,
			},
			want: "https://geo.ipify.org/api/v2/country,city,vpn?apiKey=MY_API_KEY",
		},
		{
			name: "Should get URL with country, city and vpn entities by ipAddress",
			args: args{
				apiKey:    "MY_API_KEY",
				entities:  []string{EntityCountry, EntityCity, EntityVpn},
				optParams: model.NewOptParams().Add(model.ParamIpAddress, "127.0.0.1"),
			},
			want: "https://geo.ipify.org/api/v2/country,city,vpn?apiKey=MY_API_KEY&ipAddress=127.0.0.1",
		},
		{
			name: "Should get URL with country, city and vpn entities by domain",
			args: args{
				apiKey:    "MY_API_KEY",
				entities:  []string{EntityCountry, EntityCity, EntityVpn},
				optParams: model.NewOptParams().Add(model.ParamDomain, "test.net"),
			},
			want: "https://geo.ipify.org/api/v2/country,city,vpn?apiKey=MY_API_KEY&domain=test.net",
		},
		{
			name: "Should get URL with country, city and vpn entities by email",
			args: args{
				apiKey:    "MY_API_KEY",
				entities:  []string{EntityCountry, EntityCity, EntityVpn},
				optParams: model.NewOptParams().Add(model.ParamEmail, "test@test.net"),
			},
			want: "https://geo.ipify.org/api/v2/country,city,vpn?apiKey=MY_API_KEY&email=test@test.net",
		},
		{
			name: "Should get URL with country, city and vpn entities by email because it has highest priority",
			args: args{
				apiKey:    "MY_API_KEY",
				entities:  []string{EntityCountry, EntityCity, EntityVpn},
				optParams: model.NewOptParams().Add(model.ParamIpAddress, "127.0.0.1").Add(model.ParamEmail, "test@test.net"),
			},
			want: "https://geo.ipify.org/api/v2/country,city,vpn?apiKey=MY_API_KEY&email=test@test.net",
		},
		{
			name: "Should get URL with country, city and vpn entities by domain because it has highest priority",
			args: args{
				apiKey:    "MY_API_KEY",
				entities:  []string{EntityCountry, EntityCity, EntityVpn},
				optParams: model.NewOptParams().Add(model.ParamIpAddress, "127.0.0.1").Add(model.ParamDomain, "test.net"),
			},
			want: "https://geo.ipify.org/api/v2/country,city,vpn?apiKey=MY_API_KEY&domain=test.net",
		},
		{
			name: "Should get URL with country, city and vpn entities by email because it has highest priority",
			args: args{
				apiKey:    "MY_API_KEY",
				entities:  []string{EntityCountry, EntityCity, EntityVpn},
				optParams: model.NewOptParams().Add(model.ParamEmail, "test@test.net").Add(model.ParamDomain, "test.net"),
			},
			want: "https://geo.ipify.org/api/v2/country,city,vpn?apiKey=MY_API_KEY&email=test@test.net",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildUrl(tt.args.apiKey, tt.args.entities, tt.args.optParams); got != tt.want {
				t.Errorf("buildUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}
