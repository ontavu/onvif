// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

// Code generated : DO NOT EDIT.

package device

import (
	"context"
	"github.com/ontavu/onvif/v2/networking"
)

// Call_SetNTP forwards the call to dev.CallMethod() then parses the payload of the reply as a SetNTPResponse.
func Call_SetNTP(ctx context.Context, dev *networking.Client, request SetNTP) (SetNTPResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetNTPResponse SetNTPResponse
		}
	}
	reply := Envelope{}
	httpReply, err := dev.CallMethod(ctx, request)
	if httpReply != nil {
		defer httpReply.Body.Close()
	}
	if err != nil {
		return reply.Body.SetNTPResponse, err
	} else {
		err = networking.ReadAndParse(httpReply, &reply, "SetNTP")
		return reply.Body.SetNTPResponse, err
	}
}
