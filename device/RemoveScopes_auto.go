// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

// Code generated : DO NOT EDIT.

package device

import (
	"context"
	"github.com/ontavu/onvif/v2/networking"
)

// Call_RemoveScopes forwards the call to dev.CallMethod() then parses the payload of the reply as a RemoveScopesResponse.
func Call_RemoveScopes(ctx context.Context, dev *networking.Client, request RemoveScopes) (RemoveScopesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			RemoveScopesResponse RemoveScopesResponse
		}
	}
	reply := Envelope{}
	httpReply, err := dev.CallMethod(ctx, request)
	if httpReply != nil {
		defer httpReply.Body.Close()
	}
	if err != nil {
		return reply.Body.RemoveScopesResponse, err
	} else {
		err = networking.ReadAndParse(httpReply, &reply, "RemoveScopes")
		return reply.Body.RemoveScopesResponse, err
	}
}
