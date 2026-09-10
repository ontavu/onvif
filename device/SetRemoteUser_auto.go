// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

// Code generated : DO NOT EDIT.

package device

import (
	"context"
	"github.com/jfsmig/onvif/networking"
)

// Call_SetRemoteUser forwards the call to dev.CallMethod() then parses the payload of the reply as a SetRemoteUserResponse.
func Call_SetRemoteUser(ctx context.Context, dev *networking.Client, request SetRemoteUser) (SetRemoteUserResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetRemoteUserResponse SetRemoteUserResponse
		}
	}
	reply := Envelope{}
	httpReply, err := dev.CallMethod(ctx, request)
	if httpReply != nil {
		defer httpReply.Body.Close()
	}
	if err != nil {
		return reply.Body.SetRemoteUserResponse, err
	} else {
		err = networking.ReadAndParse(httpReply, &reply, "SetRemoteUser")
		return reply.Body.SetRemoteUserResponse, err
	}
}
