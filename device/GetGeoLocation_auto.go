// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

// Code generated : DO NOT EDIT.

package device

import (
	"context"
	"github.com/ontavu/onvif/v2/networking"
)

// Call_GetGeoLocation forwards the call to dev.CallMethod() then parses the payload of the reply as a GetGeoLocationResponse.
func Call_GetGeoLocation(ctx context.Context, dev *networking.Client, request GetGeoLocation) (GetGeoLocationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetGeoLocationResponse GetGeoLocationResponse
		}
	}
	reply := Envelope{}
	httpReply, err := dev.CallMethod(ctx, request)
	if httpReply != nil {
		defer httpReply.Body.Close()
	}
	if err != nil {
		return reply.Body.GetGeoLocationResponse, err
	} else {
		err = networking.ReadAndParse(httpReply, &reply, "GetGeoLocation")
		return reply.Body.GetGeoLocationResponse, err
	}
}
