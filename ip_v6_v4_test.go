package ipify

import (
	"testing"
)

func TestGetIP64(t *testing.T) {
	got, err := GetIP64()
	if err != nil {
		t.Errorf("GetIP64() error = %v, wantErr %v", err, nil)
		return
	}
	if got == nil {
		t.Errorf("GetIP64() got = %v, want got != nil", got)
	}

	if len(got.IP) == 0 {
		t.Errorf("GetIP64() got = len(got.IP) == %v, want len(got.IP) > 0", len(got.IP))
	}
}

func TestGetIP64AsJSON(t *testing.T) {
	got, err := GetIP64AsJSON()
	if err != nil {
		t.Errorf("GetIP64AsJSON() error = %v, wantErr %v", err, nil)
		return
	}
	if len(got) == 0 {
		t.Errorf("GetIP64AsJSON() got = len(got) == %v, want len(got) > 0", len(got))
	}
}

func TestGetIP64AsJSONP(t *testing.T) {
	got, err := GetIP64AsJSONP()
	if err != nil {
		t.Errorf("GetIP64AsJSONP() error = %v, wantErr %v", err, nil)
		return
	}
	if len(got) == 0 {
		t.Errorf("GetIP64AsJSONP() got = len(got) == %v, want len(got) > 0", len(got))
	}
}

func TestGetIP64AsString(t *testing.T) {
	got, err := GetIP64AsString()
	if err != nil {
		t.Errorf("GetIP64AsString() error = %v, wantErr %v", err, nil)
		return
	}
	if len(got) == 0 {
		t.Errorf("GetIP64AsString() got = len(got) == %v, want len(got) > 0", len(got))
	}
}
