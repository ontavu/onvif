// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

// Code generated : DO NOT EDIT.

package ptz

import (
	"context"
	"github.com/jfsmig/onvif/networking"
)

// Call_Stop forwards the call to dev.CallMethod() then parses the payload of the reply as a StopResponse.
func Call_Stop(ctx context.Context, dev *networking.Client, request Stop) (StopResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			StopResponse StopResponse
		}
	}
	reply := Envelope{}
	httpReply, err := dev.CallMethod(ctx, request)
	if httpReply != nil {
		defer httpReply.Body.Close()
	}
	if err != nil {
		return reply.Body.StopResponse, err
	} else {
		err = networking.ReadAndParse(httpReply, &reply, "Stop")
		return reply.Body.StopResponse, err
	}
}
