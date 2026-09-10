// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

// Code generated : DO NOT EDIT.

package device

import (
	"context"
	"github.com/jfsmig/onvif/networking"
)

// Call_DeleteGeoLocation forwards the call to dev.CallMethod() then parses the payload of the reply as a DeleteGeoLocationResponse.
func Call_DeleteGeoLocation(ctx context.Context, dev *networking.Client, request DeleteGeoLocation) (DeleteGeoLocationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			DeleteGeoLocationResponse DeleteGeoLocationResponse
		}
	}
	reply := Envelope{}
	httpReply, err := dev.CallMethod(ctx, request)
	if httpReply != nil {
		defer httpReply.Body.Close()
	}
	if err != nil {
		return reply.Body.DeleteGeoLocationResponse, err
	} else {
		err = networking.ReadAndParse(httpReply, &reply, "DeleteGeoLocation")
		return reply.Body.DeleteGeoLocationResponse, err
	}
}
