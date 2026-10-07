// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

// Code generated : DO NOT EDIT.

package device

import (
	"context"
	"github.com/ontavu/onvif/v2/networking"
)

// Call_SetSystemDateAndTime forwards the call to dev.CallMethod() then parses the payload of the reply as a SetSystemDateAndTimeResponse.
func Call_SetSystemDateAndTime(ctx context.Context, dev *networking.Client, request SetSystemDateAndTime) (SetSystemDateAndTimeResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetSystemDateAndTimeResponse SetSystemDateAndTimeResponse
		}
	}
	reply := Envelope{}
	httpReply, err := dev.CallMethod(ctx, request)
	if httpReply != nil {
		defer httpReply.Body.Close()
	}
	if err != nil {
		return reply.Body.SetSystemDateAndTimeResponse, err
	} else {
		err = networking.ReadAndParse(httpReply, &reply, "SetSystemDateAndTime")
		return reply.Body.SetSystemDateAndTimeResponse, err
	}
}
