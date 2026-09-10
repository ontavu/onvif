// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

// Code generated : DO NOT EDIT.

package media

import (
	"context"
	"github.com/jfsmig/onvif/networking"
)

// Call_GetOSD forwards the call to dev.CallMethod() then parses the payload of the reply as a GetOSDResponse.
func Call_GetOSD(ctx context.Context, dev *networking.Client, request GetOSD) (GetOSDResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetOSDResponse GetOSDResponse
		}
	}
	reply := Envelope{}
	httpReply, err := dev.CallMethod(ctx, request)
	if httpReply != nil {
		defer httpReply.Body.Close()
	}
	if err != nil {
		return reply.Body.GetOSDResponse, err
	} else {
		err = networking.ReadAndParse(httpReply, &reply, "GetOSD")
		return reply.Body.GetOSDResponse, err
	}
}
