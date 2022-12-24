package iso3166

import (
	"reflect"
	"testing"
)

func TestRawAlpha3Match(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"GBR Test", args{s: "gb"}, GBR},
		{"GBR Num Test", args{s: "826"}, ""},
		{"USA upper Test", args{s: "USA"}, ""},
		{"Nothing Test", args{s: ""}, ""},
		{"Åland Islands Test", args{s: "Aland Islands"}, ""},
		{"Åland Islands Test", args{s: "Åland Islands"}, ""},
		{"Bahreïn Test", args{s: "Bahreïn"}, ""},
		{"Bahreïn Test (FR)", args{s: "Bahrein"}, ""},
		{"Bonaire, Sint Eustatius and Saba", args{s: "Bonaire, Sint Eustatius and Saba"}, ""},
		{"Bonaire, Sint Eustatius and Saba (FR)", args{s: "Bonaire, Sint Eustatius and Saba"}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RawAlpha3Match(tt.args.s); got != tt.want {
				t.Errorf("RawAlpha3Match() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRawAlpha2Match(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"GBR Test", args{s: "gb"}, GB},
		{"GBR Num Test", args{s: "826"}, ""},
		{"USA upper Test", args{s: "USA"}, ""},
		{"Nothing Test", args{s: ""}, ""},
		{"Åland Islands Test", args{s: "Aland Islands"}, ""},
		{"Åland Islands Test", args{s: "Åland Islands"}, ""},
		{"Bahreïn Test", args{s: "Bahreïn"}, ""},
		{"Bahreïn Test (FR)", args{s: "Bahrein"}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RawAlpha2Match(tt.args.s); got != tt.want {
				t.Errorf("RawAlpha2Match() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRawStructMatch(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want *ISOEntry
	}{
		{"GBR Test", args{s: "gb"}, &GBRStruct},
		{"GBR Test", args{s: "826"}, nil},
		{"USA upper Test", args{s: "USA"}, nil},
		{"Nothing Test", args{s: ""}, nil},
		{"Åland Islands Test", args{s: "Aland Islands"}, nil},
		{"Åland Islands Test", args{s: "Åland Islands"}, nil},
		{"Bahreïn Test", args{s: "Bahreïn"}, nil},
		{"Bahreïn Test (FR)", args{s: "Bahrein"}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RawStructMatch(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RawStructMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
