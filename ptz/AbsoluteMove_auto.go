// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

// Code generated : DO NOT EDIT.

package ptz

import (
	"context"
	"github.com/jfsmig/onvif/networking"
)

// Call_AbsoluteMove forwards the call to dev.CallMethod() then parses the payload of the reply as a AbsoluteMoveResponse.
func Call_AbsoluteMove(ctx context.Context, dev *networking.Client, request AbsoluteMove) (AbsoluteMoveResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			AbsoluteMoveResponse AbsoluteMoveResponse
		}
	}
	reply := Envelope{}
	httpReply, err := dev.CallMethod(ctx, request)
	if httpReply != nil {
		defer httpReply.Body.Close()
	}
	if err != nil {
		return reply.Body.AbsoluteMoveResponse, err
	} else {
		err = networking.ReadAndParse(httpReply, &reply, "AbsoluteMove")
		return reply.Body.AbsoluteMoveResponse, err
	}
}
