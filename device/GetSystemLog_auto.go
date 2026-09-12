// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

// Code generated : DO NOT EDIT.

package device

import (
	"context"
	"github.com/jfsmig/onvif/v2/networking"
)

// Call_GetSystemLog forwards the call to dev.CallMethod() then parses the payload of the reply as a GetSystemLogResponse.
func Call_GetSystemLog(ctx context.Context, dev *networking.Client, request GetSystemLog) (GetSystemLogResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetSystemLogResponse GetSystemLogResponse
		}
	}
	reply := Envelope{}
	httpReply, err := dev.CallMethod(ctx, request)
	if httpReply != nil {
		defer httpReply.Body.Close()
	}
	if err != nil {
		return reply.Body.GetSystemLogResponse, err
	} else {
		err = networking.ReadAndParse(httpReply, &reply, "GetSystemLog")
		return reply.Body.GetSystemLogResponse, err
	}
}
