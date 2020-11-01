// Copyright 2020 Steven Giacomelli. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package scurvy contains common utility functions.
package scurvy

import (
	"os"
	"strings"

	"github.com/mgutz/ansi"
)

type colorFunc func(string) string

type colorScheme struct {
	Magenta colorFunc
	Cyan    colorFunc
	Red     colorFunc
	Yellow  colorFunc
	Blue    colorFunc
	Green   colorFunc
	Gray    colorFunc
	Bold    colorFunc
}

func makeColorFunc(color string, colorEnabled bool) colorFunc {
	if colorEnabled {
		cf := ansi.ColorFunc(color)
		return func(v string) string {
			return cf(v)
		}
	}

	return func(v string) string {
		return v
	}
}

func newColorScheme(colorEnabled bool) *colorScheme {
	return &colorScheme{
		Magenta: makeColorFunc("magenta", colorEnabled),
		Cyan:    makeColorFunc("cyan", colorEnabled),
		Red:     makeColorFunc("red", colorEnabled),
		Yellow:  makeColorFunc("yellow", colorEnabled),
		Blue:    makeColorFunc("blue", colorEnabled),
		Green:   makeColorFunc("green", colorEnabled),
		Gray:    makeColorFunc("black+h", colorEnabled),
		Bold:    makeColorFunc("default+b", colorEnabled),
	}
}

func is256ColorSupported() bool {
	term := os.Getenv("TERM")
	colorterm := os.Getenv("COLORTERM")

	return strings.Contains(term, "256") ||
		strings.Contains(term, "24bit") ||
		strings.Contains(term, "truecolor") ||
		strings.Contains(colorterm, "256") ||
		strings.Contains(colorterm, "24bit") ||
		strings.Contains(colorterm, "truecolor")
}
