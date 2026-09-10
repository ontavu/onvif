// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

// Code generated : DO NOT EDIT.

package media

import (
	"context"
	"github.com/jfsmig/onvif/networking"
)

// Call_GetVideoSourceConfigurations forwards the call to dev.CallMethod() then parses the payload of the reply as a GetVideoSourceConfigurationsResponse.
func Call_GetVideoSourceConfigurations(ctx context.Context, dev *networking.Client, request GetVideoSourceConfigurations) (GetVideoSourceConfigurationsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetVideoSourceConfigurationsResponse GetVideoSourceConfigurationsResponse
		}
	}
	reply := Envelope{}
	httpReply, err := dev.CallMethod(ctx, request)
	if httpReply != nil {
		defer httpReply.Body.Close()
	}
	if err != nil {
		return reply.Body.GetVideoSourceConfigurationsResponse, err
	} else {
		err = networking.ReadAndParse(httpReply, &reply, "GetVideoSourceConfigurations")
		return reply.Body.GetVideoSourceConfigurationsResponse, err
	}
}
