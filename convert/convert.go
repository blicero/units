// /home/krylon/go/src/github.com/blicero/units/convert/convert.go
// -*- mode: go; coding: utf-8; -*-
// Created on 23. 12. 2022 by Benjamin Walkenhorst
// (c) 2022 Benjamin Walkenhorst
// Time-stamp: <2022-12-29 20:35:58 krylon>

// Package convert provides types and functions to convert measurements
// between various units.
package convert

import (
	"github.com/blicero/units/convert/conversion"
)

// Convert takes a value and the desired type of conversion, it returns the
func Convert(value float64, kind conversion.Conversion) (float64, error) {
	switch kind {
	}
}

func f2c(temp float64) (float64, error) {
	return (temp - 32) * (5.0 / 9.0), nil
} // func f2c(temp float64) (float64, error)

func c2f(temp float64) (float64, error) {
	return (temp * 1.8) + 32, nil
} // func c2f(temp float64) (float64, error)
