// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

package analytics

import (
	"encoding/xml"
	"strings"
	"testing"
)

// The XMLName carried the tev (events) prefix on an analytics operation, while the sibling
// fields correctly used tan. Verified against docs/wsdl/analytics.wsdl, whose
// targetNamespace is http://www.onvif.org/ver20/analytics/wsdl.
func TestCreateAnalyticsModulesUsesTheAnalyticsNamespace(t *testing.T) {
	b, err := xml.Marshal(CreateAnalyticsModules{})
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	got := string(b)
	if strings.Contains(got, "tev:") {
		t.Fatalf("analytics operation still carries the events prefix: %s", got)
	}
	if !strings.HasPrefix(got, "<tan:CreateAnalyticsModules>") {
		t.Fatalf("operation element is not tan:CreateAnalyticsModules: %s", got)
	}
}
