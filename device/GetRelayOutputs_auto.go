// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

// Code generated : DO NOT EDIT.

package device

import (
	"context"
	"github.com/jfsmig/onvif/networking"
)

// Call_GetRelayOutputs forwards the call to dev.CallMethod() then parses the payload of the reply as a GetRelayOutputsResponse.
func Call_GetRelayOutputs(ctx context.Context, dev *networking.Client, request GetRelayOutputs) (GetRelayOutputsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetRelayOutputsResponse GetRelayOutputsResponse
		}
	}
	reply := Envelope{}
	httpReply, err := dev.CallMethod(ctx, request)
	if httpReply != nil {
		defer httpReply.Body.Close()
	}
	if err != nil {
		return reply.Body.GetRelayOutputsResponse, err
	} else {
		err = networking.ReadAndParse(httpReply, &reply, "GetRelayOutputs")
		return reply.Body.GetRelayOutputsResponse, err
	}
}
