// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

// Code generated : DO NOT EDIT.

package ptz

import (
	"context"
	"github.com/jfsmig/onvif/networking"
)

// Call_SetHomePosition forwards the call to dev.CallMethod() then parses the payload of the reply as a SetHomePositionResponse.
func Call_SetHomePosition(ctx context.Context, dev *networking.Client, request SetHomePosition) (SetHomePositionResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetHomePositionResponse SetHomePositionResponse
		}
	}
	reply := Envelope{}
	httpReply, err := dev.CallMethod(ctx, request)
	if httpReply != nil {
		defer httpReply.Body.Close()
	}
	if err != nil {
		return reply.Body.SetHomePositionResponse, err
	} else {
		err = networking.ReadAndParse(httpReply, &reply, "SetHomePosition")
		return reply.Body.SetHomePositionResponse, err
	}
}
