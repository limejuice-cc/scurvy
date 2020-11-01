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
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	escapeBlueStart = "\033[0;34m"
	escapeGreyStart = "\033[0;90m"
	escapeBoldStart = "\033[0;1;39m"
	escapeEnd       = "\033[0m"
)

func TestTerminalColors(t *testing.T) {
	color := newColorScheme(true)
	assert.Equal(t, fmt.Sprintf("%sBlue%s", escapeBlueStart, escapeEnd), color.Blue("Blue"))
	assert.Equal(t, fmt.Sprintf("%sGrey%s", escapeGreyStart, escapeEnd), color.Gray("Grey"))
	assert.Equal(t, fmt.Sprintf("%sBold%s", escapeBoldStart, escapeEnd), color.Bold("Bold"))

	noColor := newColorScheme(false)
	assert.Equal(t, "Blue", noColor.Blue("Blue"))
	assert.Equal(t, "Grey", noColor.Gray("Grey"))
	assert.Equal(t, "Bold", noColor.Bold("Bold"))
}
