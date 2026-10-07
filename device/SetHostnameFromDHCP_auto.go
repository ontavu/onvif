// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

// Code generated : DO NOT EDIT.

package device

import (
	"context"
	"github.com/ontavu/onvif/v2/networking"
)

// Call_SetHostnameFromDHCP forwards the call to dev.CallMethod() then parses the payload of the reply as a SetHostnameFromDHCPResponse.
func Call_SetHostnameFromDHCP(ctx context.Context, dev *networking.Client, request SetHostnameFromDHCP) (SetHostnameFromDHCPResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetHostnameFromDHCPResponse SetHostnameFromDHCPResponse
		}
	}
	reply := Envelope{}
	httpReply, err := dev.CallMethod(ctx, request)
	if httpReply != nil {
		defer httpReply.Body.Close()
	}
	if err != nil {
		return reply.Body.SetHostnameFromDHCPResponse, err
	} else {
		err = networking.ReadAndParse(httpReply, &reply, "SetHostnameFromDHCP")
		return reply.Body.SetHostnameFromDHCPResponse, err
	}
}
