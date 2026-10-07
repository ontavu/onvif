// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

// Code generated : DO NOT EDIT.

package device

import (
	"context"
	"github.com/ontavu/onvif/v2/networking"
)

// Call_GetDNS forwards the call to dev.CallMethod() then parses the payload of the reply as a GetDNSResponse.
func Call_GetDNS(ctx context.Context, dev *networking.Client, request GetDNS) (GetDNSResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetDNSResponse GetDNSResponse
		}
	}
	reply := Envelope{}
	httpReply, err := dev.CallMethod(ctx, request)
	if httpReply != nil {
		defer httpReply.Body.Close()
	}
	if err != nil {
		return reply.Body.GetDNSResponse, err
	} else {
		err = networking.ReadAndParse(httpReply, &reply, "GetDNS")
		return reply.Body.GetDNSResponse, err
	}
}
