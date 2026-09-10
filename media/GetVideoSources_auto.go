// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

// Code generated : DO NOT EDIT.

package media

import (
	"context"
	"github.com/jfsmig/onvif/networking"
)

// Call_GetVideoSources forwards the call to dev.CallMethod() then parses the payload of the reply as a GetVideoSourcesResponse.
func Call_GetVideoSources(ctx context.Context, dev *networking.Client, request GetVideoSources) (GetVideoSourcesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetVideoSourcesResponse GetVideoSourcesResponse
		}
	}
	reply := Envelope{}
	httpReply, err := dev.CallMethod(ctx, request)
	if httpReply != nil {
		defer httpReply.Body.Close()
	}
	if err != nil {
		return reply.Body.GetVideoSourcesResponse, err
	} else {
		err = networking.ReadAndParse(httpReply, &reply, "GetVideoSources")
		return reply.Body.GetVideoSourcesResponse, err
	}
}
