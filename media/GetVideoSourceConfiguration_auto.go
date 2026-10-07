// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

// Code generated : DO NOT EDIT.

package media

import (
	"context"
	"github.com/ontavu/onvif/v2/networking"
)

// Call_GetVideoSourceConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a GetVideoSourceConfigurationResponse.
func Call_GetVideoSourceConfiguration(ctx context.Context, dev *networking.Client, request GetVideoSourceConfiguration) (GetVideoSourceConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetVideoSourceConfigurationResponse GetVideoSourceConfigurationResponse
		}
	}
	reply := Envelope{}
	httpReply, err := dev.CallMethod(ctx, request)
	if httpReply != nil {
		defer httpReply.Body.Close()
	}
	if err != nil {
		return reply.Body.GetVideoSourceConfigurationResponse, err
	} else {
		err = networking.ReadAndParse(httpReply, &reply, "GetVideoSourceConfiguration")
		return reply.Body.GetVideoSourceConfigurationResponse, err
	}
}
