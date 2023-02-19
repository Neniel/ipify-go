package ipify

import (
	"github.com/Neniel/ipify-go/model"
	"testing"
)

func Test_getIP(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		want    *model.IP
		wantErr bool
	}{
		{
			name: "Should fail on invalid url",
			args: args{
				url: "",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Should work for v4",
			args: args{
				url: baseUrlIPv4,
			},
			wantErr: false,
		},
		{
			name: "Should work for v6 or v4",
			args: args{
				url: baseUrlV6orV4,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := getIP(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("getIP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_getIPAsJSON(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Should fail on invalid url",
			args: args{
				url: "",
			},
			wantErr: true,
		},
		{
			name: "Should work for v4",
			args: args{
				url: baseUrlIPv4,
			},
			wantErr: false,
		},
		{
			name: "Should work for v6 or v4",
			args: args{
				url: baseUrlV6orV4,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := getIPAsJSON(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("getIPAsJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_getIPAsJSONP(t *testing.T) {
	type args struct {
		url      string
		callback []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Should fail for v4 because of multiple callbacks",
			args: args{
				url:      baseUrlIPv4,
				callback: []string{"callbackFunction1", "callbackFunction2"},
			},
			wantErr: true,
		},
		{
			name: "Should fail for v6 or v4 because of multiple callbacks",
			args: args{
				url:      baseUrlV6orV4,
				callback: []string{"callbackFunction1", "callbackFunction2"},
			},
			wantErr: true,
		},
		{
			name: "Should work for v4 without callback",
			args: args{
				url: baseUrlIPv4,
			},
			wantErr: false,
		},
		{
			name: "Should work for v6 or v4 without callback",
			args: args{
				url: baseUrlV6orV4,
			},
			wantErr: false,
		},
		{
			name: "Should work for v4 with callback",
			args: args{
				url:      baseUrlIPv4,
				callback: []string{"callbackFunction1"},
			},
			wantErr: false,
		},
		{
			name: "Should work for v6 or v4 with callback",
			args: args{
				url:      baseUrlV6orV4,
				callback: []string{"callbackFunction1"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := getIPAsJSONP(tt.args.url, tt.args.callback...)
			if (err != nil) != tt.wantErr {
				t.Errorf("getIPAsJSONP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_getIPAsString(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Should fail on invalid url",
			args: args{
				url: "",
			},
			wantErr: true,
		},
		{
			name: "Should work for v4",
			args: args{
				url: baseUrlIPv4,
			},
			wantErr: false,
		},
		{
			name: "Should work for v6 or v4",
			args: args{
				url: baseUrlV6orV4,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := getIPAsString(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("getIPAsString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
