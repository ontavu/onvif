// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

// Code generated : DO NOT EDIT.

package event

import (
	"context"
	"github.com/ontavu/onvif/v2/networking"
)

// Call_GetEventProperties forwards the call to dev.CallMethod() then parses the payload of the reply as a GetEventPropertiesResponse.
func Call_GetEventProperties(ctx context.Context, dev *networking.Client, request GetEventProperties) (GetEventPropertiesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetEventPropertiesResponse GetEventPropertiesResponse
		}
	}
	reply := Envelope{}
	httpReply, err := dev.CallMethod(ctx, request)
	if httpReply != nil {
		defer httpReply.Body.Close()
	}
	if err != nil {
		return reply.Body.GetEventPropertiesResponse, err
	} else {
		err = networking.ReadAndParse(httpReply, &reply, "GetEventProperties")
		return reply.Body.GetEventPropertiesResponse, err
	}
}
