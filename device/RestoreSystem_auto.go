// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

// Code generated : DO NOT EDIT.

package device

import (
	"context"
	"github.com/ontavu/onvif/v2/networking"
)

// Call_RestoreSystem forwards the call to dev.CallMethod() then parses the payload of the reply as a RestoreSystemResponse.
func Call_RestoreSystem(ctx context.Context, dev *networking.Client, request RestoreSystem) (RestoreSystemResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			RestoreSystemResponse RestoreSystemResponse
		}
	}
	reply := Envelope{}
	httpReply, err := dev.CallMethod(ctx, request)
	if httpReply != nil {
		defer httpReply.Body.Close()
	}
	if err != nil {
		return reply.Body.RestoreSystemResponse, err
	} else {
		err = networking.ReadAndParse(httpReply, &reply, "RestoreSystem")
		return reply.Body.RestoreSystemResponse, err
	}
}
