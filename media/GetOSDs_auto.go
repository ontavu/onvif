// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

// Code generated : DO NOT EDIT.

package media

import (
	"context"
	"github.com/jfsmig/onvif/networking"
)

// Call_GetOSDs forwards the call to dev.CallMethod() then parses the payload of the reply as a GetOSDsResponse.
func Call_GetOSDs(ctx context.Context, dev *networking.Client, request GetOSDs) (GetOSDsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetOSDsResponse GetOSDsResponse
		}
	}
	reply := Envelope{}
	httpReply, err := dev.CallMethod(ctx, request)
	if httpReply != nil {
		defer httpReply.Body.Close()
	}
	if err != nil {
		return reply.Body.GetOSDsResponse, err
	} else {
		err = networking.ReadAndParse(httpReply, &reply, "GetOSDs")
		return reply.Body.GetOSDsResponse, err
	}
}
