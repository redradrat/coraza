// Copyright 2022 Juan Pablo Tosso and the OWASP Coraza contributors
// SPDX-License-Identifier: Apache-2.0

package corazawaf

import (
	"fmt"
	"testing"

	"github.com/corazawaf/coraza/v3/loggers"
)

func TestLoggerLogLevels(t *testing.T) {
	waf := NewWAF()
	testCases := map[string]struct {
		logFunction                func(message string, args ...interface{})
		expectedLowestPrintedLevel int
	}{
		"Trace": {
			logFunction:                waf.Logger.Trace,
			expectedLowestPrintedLevel: 9,
		},
		"Debug": {
			logFunction:                waf.Logger.Debug,
			expectedLowestPrintedLevel: 4,
		},
		"Info": {
			logFunction:                waf.Logger.Info,
			expectedLowestPrintedLevel: 3,
		},
		"Warn": {
			logFunction:                waf.Logger.Warn,
			expectedLowestPrintedLevel: 2,
		},
		"Error": {
			logFunction:                waf.Logger.Error,
			expectedLowestPrintedLevel: 1,
		},
	}

	for name, tCase := range testCases {
		t.Run(name, func(t *testing.T) {
			for settedLevel := 0; settedLevel <= 9; settedLevel++ {
				l := &inspectableLogger{}
				waf.Logger.SetOutput(l)
				waf.Logger.SetLevel(loggers.LogLevel(settedLevel))
				tCase.logFunction("this is a log")

				if settedLevel >= tCase.expectedLowestPrintedLevel && len(l.entries) != 1 {
					t.Fatalf("Missing expected log. Level: %d, Function: %s", settedLevel, name)
				}
				if settedLevel < tCase.expectedLowestPrintedLevel && len(l.entries) == 1 {
					t.Fatalf("Unexpected log. Level: %d, Function: %s", settedLevel, name)
				}
			}
		})
	}

	t.Run("Invalid", func(t *testing.T) {
		l := &inspectableLogger{}
		waf.Logger.SetOutput(l)
		waf.Logger.SetLevel(loggers.LogLevelError)
		waf.Logger.SetLevel(loggers.LogLevel(11))
		waf.Logger.Warn("first log")
		waf.Logger.Debug("second log")

		if want, have := 2, len(l.entries); want != have {
			fmt.Println(l.entries)
			t.Fatalf("Unexpected number of logs. want: %d, have: %d", want, have)

		}

		if want, have := "[INFO] Invalid log level, defaulting to INFO\n", l.entries[0]; want != have {
			t.Errorf("Unexpected log message. want: %q, have: %q", want, have)
		}

		if want, have := "[WARN] first log\n", l.entries[1]; want != have {
			t.Errorf("Unexpected log message. want: %q, have: %q", want, have)
		}
	})
}

func TestLoggerLevelDefaultsToInfo(t *testing.T) {
	waf := NewWAF()
	waf.Logger.SetLevel(loggers.LogLevel(10))
	if waf.Logger.(*stdDebugLogger).Level != loggers.LogLevelInfo {
		t.Fatalf("Unexpected log level: %d. It should default to Info (%s)", waf.Logger.(*stdDebugLogger).Level, loggers.LogLevelInfo)
	}
}
