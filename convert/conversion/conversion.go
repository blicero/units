// /home/krylon/go/src/github.com/blicero/units/convert/conversion/conversion.go
// -*- mode: go; coding: utf-8; -*-
// Created on 23. 12. 2022 by Benjamin Walkenhorst
// (c) 2022 Benjamin Walkenhorst
// Time-stamp: <2022-12-28 18:59:19 krylon>

// Package conversion provides constants to indicate which unit is to be
// converted into which.
package conversion

//go:generate stringer -type=Conversion

// Conversion indicates what kind of unit is to be converted into which one.
type Conversion uint8

// These constants represent the various types of conversion we are going
// to perform.
// F2C - ° Fahrenheit to ° Celsius
// C2F - ° Celsius to ° Fahrenheit
const (
	F2C Conversion = iota
	C2F
)
