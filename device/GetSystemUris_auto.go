// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

// Code generated : DO NOT EDIT.

package device

import (
	"context"
	"github.com/ontavu/onvif/v2/networking"
)

// Call_GetSystemUris forwards the call to dev.CallMethod() then parses the payload of the reply as a GetSystemUrisResponse.
func Call_GetSystemUris(ctx context.Context, dev *networking.Client, request GetSystemUris) (GetSystemUrisResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetSystemUrisResponse GetSystemUrisResponse
		}
	}
	reply := Envelope{}
	httpReply, err := dev.CallMethod(ctx, request)
	if httpReply != nil {
		defer httpReply.Body.Close()
	}
	if err != nil {
		return reply.Body.GetSystemUrisResponse, err
	} else {
		err = networking.ReadAndParse(httpReply, &reply, "GetSystemUris")
		return reply.Body.GetSystemUrisResponse, err
	}
}
