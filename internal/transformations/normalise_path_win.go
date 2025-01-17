// Copyright 2022 Juan Pablo Tosso and the OWASP Coraza contributors
// SPDX-License-Identifier: Apache-2.0

package transformations

import (
	"strings"
)

func normalisePathWin(data string) (string, error) {
	leng := len(data)
	if leng < 1 {
		return data, nil
	}
	data = strings.ReplaceAll(data, "\\", "/")
	return normalisePath(data)
}
