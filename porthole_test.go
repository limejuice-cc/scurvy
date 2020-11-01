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
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"encoding/json"

	"github.com/stretchr/testify/assert"
)

func decodeLogEntryMessage(entry string) string {
	var f interface{}
	json.Unmarshal([]byte(entry), &f)
	m := f.(map[string]interface{})
	return m["message"].(string)
}

func TestPortHole(t *testing.T) {
	out := strings.Builder{}
	p := NewPortHole(ioutil.Discard, &out, nil, "info", true)

	p.Trace("Trace")
	assert.Equal(t, "", out.String())
	out.Reset()

	p.Debug("Debug")
	assert.Equal(t, "", out.String())
	out.Reset()

	p.Info("Info")
	assert.Equal(t, "Info", decodeLogEntryMessage(out.String()))
	out.Reset()

	p.Warn("Warn")
	assert.Equal(t, "Warn", decodeLogEntryMessage(out.String()))
	out.Reset()

	p.Error("Error")
	assert.Equal(t, "Error", decodeLogEntryMessage(out.String()))
	out.Reset()

	p.Fatal("Fatal")
	assert.Equal(t, "Fatal", decodeLogEntryMessage(out.String()))
	out.Reset()

	p.Panic("Panic")
	assert.Equal(t, "Panic", decodeLogEntryMessage(out.String()))

	assert.Equal(t, false, p.IsConsoleTTY())
	assert.Equal(t, false, p.IsConsoleColored())
	assert.Equal(t, nil, p.ConsoleWrite("test"))
	assert.Equal(t, nil, p.Config())
}

func TestPortHoleColors(t *testing.T) {
	// Test colored functions simply return text
	// Colors disabled because out is not a terminal

	p := NewPortHole(ioutil.Discard, ioutil.Discard, nil, "info", true)
	assert.Equal(t, "test", p.Magenta("test"))
	assert.Equal(t, "test", p.Cyan("test"))
	assert.Equal(t, "test", p.Red("test"))
	assert.Equal(t, "test", p.Yellow("test"))
	assert.Equal(t, "test", p.Blue("test"))
	assert.Equal(t, "test", p.Green("test"))
	assert.Equal(t, "test", p.Gray("test"))
	assert.Equal(t, "test", p.Bold("test"))
}

func TestBadLogLevel(t *testing.T) {
	p := NewPortHole(os.Stdout, os.Stderr, nil, "bad", true)
	assert.Equal(t, "warn", p.GetLogLevel())
}
