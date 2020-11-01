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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultConfig(t *testing.T) {
	c := NewConfig()

	assert.Equal(t, false, c.IsSet("global", "verbose"))
	c.SetValue("global", "verbose", true)
	assert.Equal(t, true, c.IsSet("global", "verbose"))

	if val, err := c.GetBool("global", "verbose"); err != nil {
		assert.FailNow(t, err.Error())
	} else {
		assert.Equal(t, true, val)
	}

	c.SetValue("global", "intval", 55)
	if val, err := c.GetInt("global", "intval"); err != nil {
		assert.FailNow(t, err.Error())
	} else {
		assert.Equal(t, 55, val)
	}

	c.SetValue("global", "strval", "test")
	if val, err := c.GetString("global", "strval"); err != nil {
		assert.FailNow(t, err.Error())
	} else {
		assert.Equal(t, "test", val)
	}

	c.SetValue("global", "strslice", []string{"a", "b"})
	if val, err := c.GetStringSlice("global", "strslice"); err != nil {
		assert.FailNow(t, err.Error())
	} else {
		assert.Equal(t, []string{"a", "b"}, val)
	}

	c.SetValue("global", "strmap", map[string]string{"a": "t", "b": "t"})
	if val, err := c.GetStringMap("global", "strmap"); err != nil {
		assert.FailNow(t, err.Error())
	} else {
		assert.Equal(t, map[string]string{"a": "t", "b": "t"}, val)
	}
}

func TestColorsEnabled(t *testing.T) {
	os.Setenv("TERM", "")
	os.Setenv("COLORTERM", "")
	assert.Equal(t, false, is256ColorSupported())
}
