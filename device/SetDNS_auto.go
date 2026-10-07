// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

// Code generated : DO NOT EDIT.

package device

import (
	"context"
	"github.com/ontavu/onvif/v2/networking"
)

// Call_SetDNS forwards the call to dev.CallMethod() then parses the payload of the reply as a SetDNSResponse.
func Call_SetDNS(ctx context.Context, dev *networking.Client, request SetDNS) (SetDNSResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetDNSResponse SetDNSResponse
		}
	}
	reply := Envelope{}
	httpReply, err := dev.CallMethod(ctx, request)
	if httpReply != nil {
		defer httpReply.Body.Close()
	}
	if err != nil {
		return reply.Body.SetDNSResponse, err
	} else {
		err = networking.ReadAndParse(httpReply, &reply, "SetDNS")
		return reply.Body.SetDNSResponse, err
	}
}
