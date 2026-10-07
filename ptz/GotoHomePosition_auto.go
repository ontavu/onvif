// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

// Code generated : DO NOT EDIT.

package ptz

import (
	"context"
	"github.com/ontavu/onvif/v2/networking"
)

// Call_GotoHomePosition forwards the call to dev.CallMethod() then parses the payload of the reply as a GotoHomePositionResponse.
func Call_GotoHomePosition(ctx context.Context, dev *networking.Client, request GotoHomePosition) (GotoHomePositionResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GotoHomePositionResponse GotoHomePositionResponse
		}
	}
	reply := Envelope{}
	httpReply, err := dev.CallMethod(ctx, request)
	if httpReply != nil {
		defer httpReply.Body.Close()
	}
	if err != nil {
		return reply.Body.GotoHomePositionResponse, err
	} else {
		err = networking.ReadAndParse(httpReply, &reply, "GotoHomePosition")
		return reply.Body.GotoHomePositionResponse, err
	}
}
