// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

// Code generated : DO NOT EDIT.

package event

import (
	"context"
	"github.com/jfsmig/onvif/networking"
)

// Call_SetSynchronizationPoint forwards the call to dev.CallMethod() then parses the payload of the reply as a SetSynchronizationPointResponse.
func Call_SetSynchronizationPoint(ctx context.Context, dev *networking.Client, request SetSynchronizationPoint) (SetSynchronizationPointResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetSynchronizationPointResponse SetSynchronizationPointResponse
		}
	}
	reply := Envelope{}
	httpReply, err := dev.CallMethod(ctx, request)
	if httpReply != nil {
		defer httpReply.Body.Close()
	}
	if err != nil {
		return reply.Body.SetSynchronizationPointResponse, err
	} else {
		err = networking.ReadAndParse(httpReply, &reply, "SetSynchronizationPoint")
		return reply.Body.SetSynchronizationPointResponse, err
	}
}
