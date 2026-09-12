// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

// Code generated : DO NOT EDIT.

package media

import (
	"context"
	"github.com/jfsmig/onvif/v2/networking"
)

// Call_GetOSDOptions forwards the call to dev.CallMethod() then parses the payload of the reply as a GetOSDOptionsResponse.
func Call_GetOSDOptions(ctx context.Context, dev *networking.Client, request GetOSDOptions) (GetOSDOptionsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetOSDOptionsResponse GetOSDOptionsResponse
		}
	}
	reply := Envelope{}
	httpReply, err := dev.CallMethod(ctx, request)
	if httpReply != nil {
		defer httpReply.Body.Close()
	}
	if err != nil {
		return reply.Body.GetOSDOptionsResponse, err
	} else {
		err = networking.ReadAndParse(httpReply, &reply, "GetOSDOptions")
		return reply.Body.GetOSDOptionsResponse, err
	}
}
