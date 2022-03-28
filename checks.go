package main

import (
	"fmt"
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func convertDia(s string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	s, _, _ = transform.String(t, s)
	return s
}

// Check map for Alpha3 Match and return ISO
// example, "gb" will return "GBR"
func Alpha3Match(s string) *string {
	var (
		returnStr string
	)
	s = strings.ToLower(strings.TrimSpace(s))

	for _, val := range isocodes {
		if s == strings.ToLower(val.alph2) || s == strings.ToLower(val.alph3) {
			returnStr = fmt.Sprintf("%v", val.alph3)
			return &returnStr
		}
		if s == strings.ToLower(strings.TrimSpace(convertDia(val.enName))) || s == strings.ToLower(strings.TrimSpace(convertDia(val.frName))) {
			returnStr = fmt.Sprintf("%v", val.alph3)
			return &returnStr
		}
		if s == strings.ToLower(strings.TrimSpace(val.enName)) || s == strings.ToLower(strings.TrimSpace(val.frName)) {
			returnStr = fmt.Sprintf("%v", val.alph3)
			return &returnStr
		}
		if s == strings.ToLower(strings.TrimSpace(val.numCode)) {
			returnStr = fmt.Sprintf("%v", val.alph3)
			return &returnStr
		}
	}
	return nil
}

// Check map for Alpha2 Match and return ISO
// example, "gb" will return "GB"
func Alpha2Match(s string) *string {
	var (
		returnStr string
	)
	s = strings.ToLower(strings.TrimSpace(s))
	for _, val := range isocodes {
		if s == strings.ToLower(val.alph2) || s == strings.ToLower(val.alph3) {
			returnStr = fmt.Sprintf("%v", val.alph2)
			return &returnStr
		}
		if s == strings.ToLower(strings.TrimSpace(convertDia(val.enName))) || s == strings.ToLower(strings.TrimSpace(convertDia(val.frName))) {
			returnStr = fmt.Sprintf("%v", val.alph2)
			return &returnStr
		}
		if s == strings.ToLower(strings.TrimSpace(val.enName)) || s == strings.ToLower(strings.TrimSpace(val.frName)) {
			returnStr = fmt.Sprintf("%v", val.alph2)
			return &returnStr
		}
		if s == strings.ToLower(strings.TrimSpace(val.numCode)) {
			returnStr = fmt.Sprintf("%v", val.alph2)
			return &returnStr
		}
	}
	return nil
}

// Check map for Alpha3 Match and return ISO
// example, "gb" will return GBR Struct 
func StructMatch(s string) *IsoEntry {
	var (
		entry IsoEntry
	)
	s = strings.ToLower(strings.TrimSpace(s))
	for _, val := range isocodes {
		if s == strings.ToLower(val.alph2) || s == strings.ToLower(val.alph3) {
			entry = IsoEntry{
				enName:  val.enName,
				frName:  val.frName,
				alph2:   val.alph2,
				alph3:   val.alph3,
				numCode: val.numCode,
			}
			return &entry
		}
		if s == strings.ToLower(strings.TrimSpace(convertDia(val.enName))) || s == strings.ToLower(strings.TrimSpace(convertDia(val.frName))) {
			entry = IsoEntry{
				enName:  val.enName,
				frName:  val.frName,
				alph2:   val.alph2,
				alph3:   val.alph3,
				numCode: val.numCode,
			}
			return &entry
		}
		if s == strings.ToLower(strings.TrimSpace(val.enName)) || s == strings.ToLower(strings.TrimSpace(val.frName)) {
			entry = IsoEntry{
				enName:  val.enName,
				frName:  val.frName,
				alph2:   val.alph2,
				alph3:   val.alph3,
				numCode: val.numCode,
			}
			return &entry
		}
		if s == strings.ToLower(strings.TrimSpace(val.numCode)) {
			entry = IsoEntry{
				enName:  val.enName,
				frName:  val.frName,
				alph2:   val.alph2,
				alph3:   val.alph3,
				numCode: val.numCode,
			}
			return &entry
		}

	}
	return nil
}
