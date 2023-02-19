package ipify

import (
	"testing"
)

func TestGetIPv4(t *testing.T) {
	got, err := GetIPv4()
	if err != nil {
		t.Errorf("GetIPv4() error = %v, wantErr %v", err, nil)
		return
	}
	if got == nil {
		t.Errorf("GetIPv4() got = %v, want got != nil", got)
	}

	if len(got.IP) == 0 {
		t.Errorf("GetIPv4() got = len(got.IP) == %v, want len(got.IP) > 0", len(got.IP))
	}
}

func TestGetIPv4AsJSON(t *testing.T) {
	got, err := GetIPv4AsJSON()
	if err != nil {
		t.Errorf("GetIPv4AsJSON() error = %v, wantErr %v", err, nil)
		return
	}
	if len(got) == 0 {
		t.Errorf("GetIPv4AsJSON() got = len(got) == %v, want len(got) > 0", len(got))
	}
}

func TestGetIPv4AsJSONP(t *testing.T) {
	got, err := GetIPv4AsJSONP()
	if err != nil {
		t.Errorf("GetIPv4AsJSONP() error = %v, wantErr %v", err, nil)
		return
	}
	if len(got) == 0 {
		t.Errorf("GetIPv4AsJSONP() got = len(got) == %v, want len(got) > 0", len(got))
	}
}

func TestGetIPv4AsString(t *testing.T) {
	got, err := GetIPv4AsString()
	if err != nil {
		t.Errorf("GetIPv4AsString() error = %v, wantErr %v", err, nil)
		return
	}
	if len(got) == 0 {
		t.Errorf("GetIPv4AsString() got = len(got) == %v, want len(got) > 0", len(got))
	}
}
