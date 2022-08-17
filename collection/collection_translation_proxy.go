// Copyright 2022 Juan Pablo Tosso
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package collection

import (
	"regexp"

	"github.com/corazawaf/coraza/v3/types"
	"github.com/corazawaf/coraza/v3/types/variables"
)

// TranslationProxy are used to store VARIABLE data
// for transactions, this data structured is designed
// to store slices of data for keys
// Important: CollectionTranslationProxys ARE NOT concurrent safe
type TranslationProxy struct {
	data     []*Map
	name     string
	variable variables.RuleVariable
}

// FindRegex returns a slice of MatchData for the regex
func (c *TranslationProxy) FindRegex(key *regexp.Regexp) []types.MatchData {
	res := []types.MatchData{}
	keys := []string{}
	for _, c := range c.data {
		keys = append(keys, c.keysRx(key)...)
	}
	for _, k := range keys {
		res = append(res, types.MatchData{
			VariableName: c.name,
			Variable:     c.variable,
			Value:        k,
		})
	}
	return res
}

// FindString returns a slice of MatchData for the string
func (c *TranslationProxy) FindString(key string) []types.MatchData {
	for _, c := range c.data {
		if len(c.Get(key)) > 0 {
			return []types.MatchData{
				{
					VariableName: c.name,
					Variable:     c.variable,
					Value:        key,
				},
			}
		}
	}
	return nil
}

// FindAll returns all keys from Proxy Collections
func (c *TranslationProxy) FindAll() []types.MatchData {
	keys := []string{}
	for _, c := range c.data {
		keys = append(keys, c.keys()...)
	}
	res := []types.MatchData{}
	for _, k := range keys {
		res = append(res, types.MatchData{
			VariableName: c.name,
			Variable:     c.variable,
			Value:        k,
		})
	}
	return res
}

// Data returns the keys of all Proxy collections
func (c *TranslationProxy) Data() []string {
	res := []string{}
	for _, c := range c.data {
		res = append(res, c.keys()...)
	}
	return res
}

// Name returns the name for the current CollectionTranslationProxy
func (c *TranslationProxy) Name() string {
	return c.name
}

// Reset the current CollectionTranslationProxy
func (c *TranslationProxy) Reset() {
	// do nothing
}

// Get the value for the index
func (c *TranslationProxy) Get(index int) string {
	if index < len(c.data) {
		if v := c.data[index].Get(""); len(v) > 0 {
			return v[0]
		}
	}
	return ""
}

var _ Collection = &TranslationProxy{}

// NewTranslationProxy creates a translation proxy
// Translation proxies are used to merge variable keys from multiple collections
func NewTranslationProxy(variable variables.RuleVariable, data ...*Map) *TranslationProxy {
	return &TranslationProxy{
		name:     variable.Name(),
		variable: variable,
		data:     data,
	}
}