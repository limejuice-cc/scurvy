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
	"io"

	cmap "github.com/orcaman/concurrent-map"
	"github.com/rs/zerolog"
)

// Config encapsulates a config store.
type Config interface {
	IsSet(ns, key string) bool

	GetBool(ns, key string) (bool, error)
	GetInt(ns, key string) (int, error)
	GetString(ns, key string) (string, error)
	GetStringSlice(ns, key string) ([]string, error)
	GetStringMap(ns, key string) (map[string]string, error)

	SetValue(ns, key string, value interface{}) error
}

// NewConfig creates a new config store.
func NewConfig() Config {
	return &defaultConfig{values: cmap.New()}
}

// PortHole encapsulates an execution context.
type PortHole interface {
	Trace(string, ...interface{})
	Debug(string, ...interface{})
	Info(string, ...interface{})
	Warn(string, ...interface{})
	Error(string, ...interface{})
	Fatal(string, ...interface{})
	Panic(string, ...interface{})
	GetLogLevel() string

	Magenta(string) string
	Cyan(string) string
	Red(string) string
	Yellow(string) string
	Blue(string) string
	Green(string) string
	Gray(string) string
	Bold(string) string

	IsConsoleTTY() bool
	IsConsoleColored() bool
	ConsoleWrite(string, ...interface{}) error

	Config() Config
}

// NewPortHole creates a new PortHole.
func NewPortHole(consoleOut, errorOut io.Writer, config Config, logLevel string, enableColors bool) PortHole {
	lvl, err := zerolog.ParseLevel(logLevel)
	if err != nil {
		lvl = zerolog.WarnLevel
	}

	colorsEnabled := enableColors
	if !is256ColorSupported() {
		colorsEnabled = false
	}

	isTTY := isTerminal(consoleOut)
	if !isTTY {
		colorsEnabled = false
	}

	return &defaultPortHole{
		colors:        newColorScheme(colorsEnabled),
		colorsEnabled: colorsEnabled,
		isTTY:         isTTY,
		output:        consoleOut,
		logger:        zerolog.New(errorOut).Level(lvl).With().Timestamp().Logger(),
		config:        config,
	}
}
