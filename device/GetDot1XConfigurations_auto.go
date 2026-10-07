// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

// Code generated : DO NOT EDIT.

package device

import (
	"context"
	"github.com/ontavu/onvif/v2/networking"
)

// Call_GetDot1XConfigurations forwards the call to dev.CallMethod() then parses the payload of the reply as a GetDot1XConfigurationsResponse.
func Call_GetDot1XConfigurations(ctx context.Context, dev *networking.Client, request GetDot1XConfigurations) (GetDot1XConfigurationsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetDot1XConfigurationsResponse GetDot1XConfigurationsResponse
		}
	}
	reply := Envelope{}
	httpReply, err := dev.CallMethod(ctx, request)
	if httpReply != nil {
		defer httpReply.Body.Close()
	}
	if err != nil {
		return reply.Body.GetDot1XConfigurationsResponse, err
	} else {
		err = networking.ReadAndParse(httpReply, &reply, "GetDot1XConfigurations")
		return reply.Body.GetDot1XConfigurationsResponse, err
	}
}
