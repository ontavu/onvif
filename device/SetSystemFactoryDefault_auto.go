// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

// Code generated : DO NOT EDIT.

package device

import (
	"context"
	"github.com/jfsmig/onvif/networking"
)

// Call_SetSystemFactoryDefault forwards the call to dev.CallMethod() then parses the payload of the reply as a SetSystemFactoryDefaultResponse.
func Call_SetSystemFactoryDefault(ctx context.Context, dev *networking.Client, request SetSystemFactoryDefault) (SetSystemFactoryDefaultResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetSystemFactoryDefaultResponse SetSystemFactoryDefaultResponse
		}
	}
	reply := Envelope{}
	httpReply, err := dev.CallMethod(ctx, request)
	if httpReply != nil {
		defer httpReply.Body.Close()
	}
	if err != nil {
		return reply.Body.SetSystemFactoryDefaultResponse, err
	} else {
		err = networking.ReadAndParse(httpReply, &reply, "SetSystemFactoryDefault")
		return reply.Body.SetSystemFactoryDefaultResponse, err
	}
}
