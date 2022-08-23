// Copyright 2022 Juan Pablo Tosso and the OWASP Coraza contributors
// SPDX-License-Identifier: Apache-2.0

//go:build tinygo
// +build tinygo

package operators

import (
	"github.com/corazawaf/coraza/v3"
)

type rbl struct{}

func (o *rbl) Init(_ coraza.RuleOperatorOptions) error { return nil }

func (o *rbl) Evaluate(_ *coraza.Transaction, _ string) bool { return true }