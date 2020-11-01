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
	"fmt"
	"io"
	"os"

	"github.com/mattn/go-isatty"
	"github.com/rs/zerolog"
)

type defaultPortHole struct {
	colors        *colorScheme
	colorsEnabled bool
	isTTY         bool
	output        io.Writer
	logger        zerolog.Logger
	config        Config
}

func (p *defaultPortHole) Trace(format string, values ...interface{}) {
	p.logger.WithLevel(zerolog.TraceLevel).Msgf(format, values...)
}

func (p *defaultPortHole) Debug(format string, values ...interface{}) {
	p.logger.WithLevel(zerolog.DebugLevel).Msgf(format, values...)
}

func (p *defaultPortHole) Info(format string, values ...interface{}) {
	p.logger.WithLevel(zerolog.InfoLevel).Msgf(format, values...)
}

func (p *defaultPortHole) Warn(format string, values ...interface{}) {
	p.logger.WithLevel(zerolog.WarnLevel).Msgf(format, values...)
}

func (p *defaultPortHole) Error(format string, values ...interface{}) {
	p.logger.WithLevel(zerolog.ErrorLevel).Msgf(format, values...)
}

func (p *defaultPortHole) Fatal(format string, values ...interface{}) {
	p.logger.WithLevel(zerolog.FatalLevel).Msgf(format, values...)
}

func (p *defaultPortHole) Panic(format string, values ...interface{}) {
	p.logger.WithLevel(zerolog.PanicLevel).Msgf(format, values...)
}

func (p *defaultPortHole) GetLogLevel() string {
	return p.logger.GetLevel().String()
}

func (p *defaultPortHole) Magenta(value string) string {
	return p.colors.Magenta(value)
}

func (p *defaultPortHole) Cyan(value string) string {
	return p.colors.Cyan(value)
}

func (p *defaultPortHole) Red(value string) string {
	return p.colors.Red(value)
}

func (p *defaultPortHole) Yellow(value string) string {
	return p.colors.Yellow(value)
}

func (p *defaultPortHole) Blue(value string) string {
	return p.colors.Blue(value)
}

func (p *defaultPortHole) Green(value string) string {
	return p.colors.Green(value)
}

func (p *defaultPortHole) Gray(value string) string {
	return p.colors.Gray(value)
}

func (p *defaultPortHole) Bold(value string) string {
	return p.colors.Bold(value)
}

func (p *defaultPortHole) IsConsoleTTY() bool {
	return p.isTTY
}

func (p *defaultPortHole) IsConsoleColored() bool {
	return p.colorsEnabled
}

func (p *defaultPortHole) ConsoleWrite(format string, values ...interface{}) error {
	_, err := fmt.Fprintf(p.output, format, values...)
	return err
}

func (p *defaultPortHole) Config() Config {
	return p.config
}

func isTerminal(w interface{}) bool {
	if f, ok := w.(*os.File); ok {
		return isatty.IsTerminal(f.Fd())
	}
	return false
}
