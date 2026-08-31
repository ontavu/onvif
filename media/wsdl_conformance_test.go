// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

package media

import (
	"encoding/xml"
	"strings"
	"testing"
)

// The XMLName tag names the operation element that goes on the wire. It read
// trt:GetDeviceInformation — a copy-paste slip, with the sibling fields correctly trt:
// prefixed — so calling SetMetadataConfiguration actually asked the device for its
// information. Verified against docs/wsdl/media.wsdl.
func TestSetMetadataConfigurationNamesItsOwnOperation(t *testing.T) {
	b, err := xml.Marshal(SetMetadataConfiguration{})
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	got := string(b)
	if strings.Contains(got, "GetDeviceInformation") {
		t.Fatalf("SetMetadataConfiguration still serialises as GetDeviceInformation: %s", got)
	}
	if !strings.HasPrefix(got, "<trt:SetMetadataConfiguration>") {
		t.Fatalf("operation element is not trt:SetMetadataConfiguration: %s", got)
	}
}
