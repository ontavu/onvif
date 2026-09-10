// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

// Code generated : DO NOT EDIT.

package device

import (
	"context"
	"github.com/jfsmig/onvif/networking"
)

// Call_GetRemoteDiscoveryMode forwards the call to dev.CallMethod() then parses the payload of the reply as a GetRemoteDiscoveryModeResponse.
func Call_GetRemoteDiscoveryMode(ctx context.Context, dev *networking.Client, request GetRemoteDiscoveryMode) (GetRemoteDiscoveryModeResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetRemoteDiscoveryModeResponse GetRemoteDiscoveryModeResponse
		}
	}
	reply := Envelope{}
	httpReply, err := dev.CallMethod(ctx, request)
	if httpReply != nil {
		defer httpReply.Body.Close()
	}
	if err != nil {
		return reply.Body.GetRemoteDiscoveryModeResponse, err
	} else {
		err = networking.ReadAndParse(httpReply, &reply, "GetRemoteDiscoveryMode")
		return reply.Body.GetRemoteDiscoveryModeResponse, err
	}
}
