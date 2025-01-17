// Copyright 2022 Juan Pablo Tosso and the OWASP Coraza contributors
// SPDX-License-Identifier: Apache-2.0

package transformations

import "testing"

func TestLength(t *testing.T) {
	tests := []struct {
		input  string
		length string
	}{
		{
			input:  "hello",
			length: "5",
		},
		{
			input:  "",
			length: "0",
		},
		{
			input:  "ハローワールド",
			length: "21",
		},
	}

	for _, tc := range tests {
		tt := tc
		t.Run(tt.input, func(t *testing.T) {
			have, err := length(tt.input)
			if err != nil {
				t.Error(err)
			}
			if have != tt.length {
				t.Errorf("Expected %s, have %s", tt.length, have)
			}
		})
	}
}
