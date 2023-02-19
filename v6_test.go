package ipify

import (
	"testing"
)

func TestGetIPv6(t *testing.T) {
	got, err := GetIPv6()
	if err != nil {
		t.Errorf("GetIPv6() error = %v, wantErr %v", err, nil)
		return
	}
	if got == nil {
		t.Errorf("GetIPv6() got = %v, want got != nil", got)
	}

	if len(got.IP) == 0 {
		t.Errorf("GetIPv6() got = len(got.IP) == %v, want len(got.IP) > 0", len(got.IP))
	}
}

func TestGetIPv6AsJSON(t *testing.T) {
	got, err := GetIPv6AsJSON()
	if err != nil {
		t.Errorf("GetIPv6AsJSON() error = %v, wantErr %v", err, nil)
		return
	}
	if len(got) == 0 {
		t.Errorf("GetIPv6AsJSON() got = len(got) == %v, want len(got) > 0", len(got))
	}
}

func TestGetIPv6AsJSONP(t *testing.T) {
	got, err := GetIPv6AsJSONP()
	if err != nil {
		t.Errorf("GetIPv6AsJSONP() error = %v, wantErr %v", err, nil)
		return
	}
	if len(got) == 0 {
		t.Errorf("GetIPv6AsJSONP() got = len(got) == %v, want len(got) > 0", len(got))
	}
}

func TestGetIPv6AsString(t *testing.T) {
	got, err := GetIPv6AsString()
	if err != nil {
		t.Errorf("GetIPv6AsString() error = %v, wantErr %v", err, nil)
		return
	}
	if len(got) == 0 {
		t.Errorf("GetIPv6AsString() got = len(got) == %v, want len(got) > 0", len(got))
	}
}
