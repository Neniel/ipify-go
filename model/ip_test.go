package model

import (
	"reflect"
	"testing"
)

func TestNewIP(t *testing.T) {
	tests := []struct {
		name string
		want *IP
	}{
		{
			name: "Should get a default IP object",
			want: &IP{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewIP(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIP() = %v, want %v", got, tt.want)
			}
		})
	}
}
