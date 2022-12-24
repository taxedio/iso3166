package iso3166

import (
	"strings"
)

// Check map for Alpha3 Match and returns Alpha3 code as string
// example, "gb" will return "GBR"
func RawAlpha3Match(s string) string {
	if c, exists := isocodes[strings.ToUpper(s)]; exists {
		return c.alph3
	}
	return ""
}

// Check map for Alpha3 Match and returns Alpha3 code as string
// example, "gb" will return "GB"
func RawAlpha2Match(s string) string {
	if c, exists := isocodes[strings.ToUpper(s)]; exists {
		return c.alph2
	}
	return ""
}

// Check map for Alpha3 Match and return pointer to ISO Struct
// example, "gb" will return GBR Struct
func RawStructMatch(s string) *ISOEntry {
	if c, exists := isocodes[strings.ToUpper(s)]; exists {
		return &c
	}
	return nil
}
