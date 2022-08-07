package iso3166

import (
	"reflect"
	"testing"
)

var (
	Nil       *string
	NilStruct *ISOEntry

	GBR       string = isocodes["GB"].alph3
	GB        string = isocodes["GB"].alph2
	GBRStruct        = isocodes["GB"]

	USA       string = isocodes["US"].alph3
	US        string = isocodes["US"].alph2
	USAStruct        = isocodes["US"]

	ALA       string = isocodes["AX"].alph3
	AX        string = isocodes["AX"].alph2
	ALAStruct        = isocodes["AX"]

	BHR       string = isocodes["BH"].alph3
	BH        string = isocodes["BH"].alph2
	BHRStruct        = isocodes["BH"]

	BQ        string = isocodes["BQ"].alph2
	BES       string = isocodes["BQ"].alph3
	BESStruct        = isocodes["BQ"]
)

func TestAlpha3Match(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want *string
	}{
		{"GBR Test", args{s: "gb"}, &GBR},
		{"GBR Num Test", args{s: "826"}, &GBR},
		{"USA upper Test", args{s: "USA"}, &USA},
		{"Nothing Test", args{s: ""}, Nil},
		{"Åland Islands Test", args{s: "Aland Islands"}, &ALA},
		{"Åland Islands Test", args{s: "Åland Islands"}, &ALA},
		{"Bahreïn Test", args{s: "Bahreïn"}, &BHR},
		{"Bahreïn Test (FR)", args{s: "Bahrein"}, &BHR},
		{"Bonaire, Sint Eustatius and Saba", args{s: "Bonaire, Sint Eustatius and Saba"}, &BES},
		{"Bonaire, Sint Eustatius and Saba (FR)", args{s: "Bonaire, Sint Eustatius and Saba"}, &BES},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Alpha3Match(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Alpha3Match() = %#v, want %#v", *got, *tt.want)
			}
		})
	}
}

func TestAlpha2Match(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want *string
	}{
		{"GBR Test", args{s: "gb"}, &GB},
		{"GBR Num Test", args{s: "826"}, &GB},
		{"USA upper Test", args{s: "USA"}, &US},
		{"Nothing Test", args{s: ""}, Nil},
		{"Åland Islands Test", args{s: "Aland Islands"}, &AX},
		{"Åland Islands Test", args{s: "Åland Islands"}, &AX},
		{"Bahreïn Test", args{s: "Bahreïn"}, &BH},
		{"Bahreïn Test (FR)", args{s: "Bahrein"}, &BH},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Alpha2Match(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Alpha3Match() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStructMatch(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want *ISOEntry
	}{
		{"GBR Test", args{s: "gb"}, &GBRStruct},
		{"GBR Test", args{s: "826"}, &GBRStruct},
		{"USA upper Test", args{s: "USA"}, &USAStruct},
		{"Nothing Test", args{s: ""}, NilStruct},
		{"Åland Islands Test", args{s: "Aland Islands"}, &ALAStruct},
		{"Åland Islands Test", args{s: "Åland Islands"}, &ALAStruct},
		{"Bahreïn Test", args{s: "Bahreïn"}, &BHRStruct},
		{"Bahreïn Test (FR)", args{s: "Bahrein"}, &BHRStruct},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StructMatch(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StructMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
