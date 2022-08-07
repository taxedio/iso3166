<img src="assets\taxediologolandscape.jpg" alt="drawing" width="200"/>

<h1 align="center">
  ISO 3166 Country Codes
</h1>

<h3 align="center">
  <!-- <a href="https://pkg.go.dev/github.com/KalbiProject/Kalbi">Documentation</a> •  -->
  <a href="https://taxed.io">taxed.io</a>
</h3>

[![Go Report Card](https://goreportcard.com/badge/github.com/taxedio/iso3166)](https://goreportcard.com/report/github.com/taxedio/iso3166)
[![GitHub license](https://img.shields.io/github/license/Naereen/StrapDown.js.svg)](https://github.com/taxedio/iso3166/LICENCE)  
[![GitHub contributors](https://img.shields.io/github/contributors/taxedio/iso3166)](https://github.com/taxedio/iso3166/graphs/contributors)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/taxedio/iso3166)

# Introduction

This package provides the latest ISO-3166 codes (<a href="https://www.iso.org/iso-3166-country-codes.html"><b>iso.org</b></a>)

# Basic Example

```GO
package main

import (
	"fmt"

	"github.com/taxedio/iso3166"
)

var (
	enName  string = "Albania"
	frName  string = "Albanie (l')"
	alph2   string = "AL"
	alph3   string = "ALB"
	numCode string = "008"
)

func main() {
	// English Names, codes can be found via English name (enName in Struct)
	fmt.Printf(`iso3166.Alpha3Match("Albania") - %v`, *iso3166.Alpha3Match(enName))
	fmt.Printf(`iso3166.Alpha2Match("Albania") - %v`, *iso3166.Alpha2Match(enName))

  // English Names with different casing and spaces will still return requested codes
	fmt.Printf(`iso3166.Alpha3Match(" albania   ") - %v`, *iso3166.Alpha3Match(" albania   "))
	fmt.Printf(`iso3166.Alpha2Match(" albania   ") - %v`, *iso3166.Alpha2Match(" albania   "))

  // French Names, codes can be found via English name (frName in Struct)
	fmt.Printf(`iso3166.Alpha3Match("Albanie (l')") - %v`, *iso3166.Alpha3Match(frName))
	fmt.Printf(`iso3166.Alpha2Match("Albanie (l')") - %v`, *iso3166.Alpha2Match(frName))

  // Alpha Codes, both can return the other:
	fmt.Printf(`iso3166.Alpha3Match("AL") - %v`, *iso3166.Alpha3Match(alph2))
	fmt.Printf(`iso3166.Alpha2Match("ALB") - %v`, *iso3166.Alpha2Match(alph3))

  // Numeric Codes, can be used to search for codes and Structs:
	fmt.Printf(`iso3166.Alpha3Match("AL") - %v`, *iso3166.Alpha3Match(numCode))
	fmt.Printf(`iso3166.Alpha2Match("ALB") - %v`, *iso3166.Alpha2Match(numCode))

  // Diacritics, (IN DEVELOPMENT). If A is used instead of Å then the values will still match:
	fmt.Printf(`iso3166.Alpha2Match("Åland Islands") - %v`, *iso3166.Alpha3Match("Åland Islands"))
	fmt.Printf(`iso3166.Alpha2Match("Åland Islands") - %v`, *iso3166.Alpha2Match("Aland Islands"))
	fmt.Printf(`iso3166.Alpha3Match("Aland Islands") - %v`, *iso3166.Alpha3Match("Aland Islands"))
	fmt.Printf(`iso3166.Alpha3Match("Aland Islands") - %v`, *iso3166.Alpha2Match("Aland Islands"))

  // Returning Structs, all of the rules and examples above work using the StructMatch
	fmt.Printf(`iso3166.StructMatch("Albania") - %v`, *iso3166.StructMatch(enName))
	fmt.Printf(`iso3166.StructMatch("  albania  ") - %v`, *iso3166.StructMatch(" albania  "))
	fmt.Printf(`iso3166.StructMatch("Albanie (l')") - %v`, *iso3166.StructMatch(frName))
	fmt.Printf(`iso3166.StructMatch("AL") - %v`, *iso3166.StructMatch(alph2))
	fmt.Printf(`iso3166.StructMatch("ALB") - %v`, *iso3166.StructMatch(alph3))
	fmt.Printf(`iso3166.StructMatch("008") - %v`, *iso3166.StructMatch(numCode))
	fmt.Printf(`iso3166.StructMatch("Aland Islands") - %v`, *iso3166.StructMatch("Aland Islands"))
	fmt.Printf(`iso3166.StructMatch("aland Islands  ") - %v`, *iso3166.StructMatch("aland Islands  "))
}
```

**console**:

```stdout
// English Names, codes can be found via English name (enName in Struct)
iso3166.Alpha3Match("Albania") - ALB
iso3166.Alpha2Match("Albania") - AL

// English Names with different casing and spaces will still return requested codes
iso3166.Alpha3Match(" albania   ") - ALB
iso3166.Alpha2Match(" albania   ") - AL

// French Names, codes can be found via English name (frName in Struct)
iso3166.Alpha3Match("Albanie (l')") - ALB
iso3166.Alpha2Match("Albanie (l')") - AL

// Alpha Codes, both can return the other:
iso3166.Alpha3Match("AL") - ALB
iso3166.Alpha2Match("ALB") - AL

// Numeric Codes, can be used to search for codes and Structs:
iso3166.Alpha3Match("AL") - ALB
iso3166.Alpha2Match("ALB") - AL

// Diacritics, (IN DEVELOPMENT). If A is used instead of Å then the values will still match:
iso3166.Alpha2Match("Åland Islands") - ALA
iso3166.Alpha2Match("Åland Islands") - AX
iso3166.Alpha3Match("Aland Islands") - ALA
iso3166.Alpha3Match("Aland Islands") - AX

// Returning Structs, all of the rules and examples above work using the StructMatch
iso3166.StructMatch("Albania") - {Albania Albanie (l') AL ALB 008}
iso3166.StructMatch("  albania  ") - {Albania Albanie (l') AL ALB 008}
iso3166.StructMatch("Albanie (l')") - {Albania Albanie (l') AL ALB 008}
iso3166.StructMatch("AL") - {Albania Albanie (l') AL ALB 008}
iso3166.StructMatch("ALB") - {Albania Albanie (l') AL ALB 008}
iso3166.StructMatch("008") - {Albania Albanie (l') AL ALB 008}
iso3166.StructMatch("Aland Islands") - {Åland Islands Åland(les Îles) AX ALA 248}
iso3166.StructMatch("aland Islands  ") - {Åland Islands Åland(les Îles) AX ALA 248}
```
