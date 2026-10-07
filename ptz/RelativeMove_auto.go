// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

// Code generated : DO NOT EDIT.

package ptz

import (
	"context"
	"github.com/ontavu/onvif/v2/networking"
)

// Call_RelativeMove forwards the call to dev.CallMethod() then parses the payload of the reply as a RelativeMoveResponse.
func Call_RelativeMove(ctx context.Context, dev *networking.Client, request RelativeMove) (RelativeMoveResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			RelativeMoveResponse RelativeMoveResponse
		}
	}
	reply := Envelope{}
	httpReply, err := dev.CallMethod(ctx, request)
	if httpReply != nil {
		defer httpReply.Body.Close()
	}
	if err != nil {
		return reply.Body.RelativeMoveResponse, err
	} else {
		err = networking.ReadAndParse(httpReply, &reply, "RelativeMove")
		return reply.Body.RelativeMoveResponse, err
	}
}
