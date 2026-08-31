// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

package utils

import (
	"errors"
	"fmt"
	"testing"
)

// The sentinels are wrapped at their call sites — networking.ReadAndParse attaches the
// HTTP status — so the contract that matters is that errors.Is still matches through the
// wrapping, not that the caller can compare with ==.
func TestSentinelSurvivesWrapping(t *testing.T) {
	for _, tc := range []struct {
		name string
		err  error
	}{
		{"ErrHTTP", ErrHTTP},
		{"ErrNotOnvif", ErrNotOnvif},
	} {
		t.Run(tc.name, func(t *testing.T) {
			wrapped := fmt.Errorf("%w: 401 Unauthorized", tc.err)
			if !errors.Is(wrapped, tc.err) {
				t.Fatalf("errors.Is lost the sentinel through one wrap: %v", wrapped)
			}
			if !errors.Is(fmt.Errorf("outer: %w", wrapped), tc.err) {
				t.Fatal("errors.Is lost the sentinel through two wraps")
			}
		})
	}
}

// The two sentinels must stay distinguishable: a constError compares by its string, so
// giving two of them the same text would silently make errors.Is match both.
func TestSentinelsAreDistinct(t *testing.T) {
	if errors.Is(ErrHTTP, ErrNotOnvif) {
		t.Fatal("ErrHTTP and ErrNotOnvif are indistinguishable")
	}
}
