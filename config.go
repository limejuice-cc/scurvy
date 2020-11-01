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

	cmap "github.com/orcaman/concurrent-map"
	"github.com/spf13/cast"
)

type defaultConfig struct {
	values cmap.ConcurrentMap
}

func formatKey(ns, key string) string {
	return fmt.Sprintf("%s.%s", ns, key)
}

func (c *defaultConfig) IsSet(ns, key string) bool {
	_, ok := c.values.Get(formatKey(ns, key))
	return ok
}

func (c *defaultConfig) SetValue(ns, key string, value interface{}) error {
	c.values.Set(formatKey(ns, key), value)
	return nil
}

func (c *defaultConfig) GetBool(ns, key string) (bool, error) {
	val, _ := c.values.Get(formatKey(ns, key))
	return cast.ToBoolE(val)
}

func (c *defaultConfig) GetInt(ns, key string) (int, error) {
	val, _ := c.values.Get(formatKey(ns, key))
	return cast.ToIntE(val)
}

func (c *defaultConfig) GetString(ns, key string) (string, error) {
	val, _ := c.values.Get(formatKey(ns, key))
	return cast.ToStringE(val)
}

func (c *defaultConfig) GetStringSlice(ns, key string) ([]string, error) {
	val, _ := c.values.Get(formatKey(ns, key))
	return cast.ToStringSliceE(val)
}

func (c *defaultConfig) GetStringMap(ns, key string) (map[string]string, error) {
	val, _ := c.values.Get(formatKey(ns, key))
	return cast.ToStringMapStringE(val)
}
