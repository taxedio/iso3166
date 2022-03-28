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

fmt.Println(iso3166.Alpha3Match("afghanistan"))
fmt.Println(iso3166.Alpha2Match("afghanistan"))

```
**stdout**:
```stdout
AFG
AF
```
