// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

// Code generated : DO NOT EDIT.

package event

import (
	"context"
	"github.com/ontavu/onvif/v2/networking"
)

// Call_Unsubscribe forwards the call to dev.CallMethod() then parses the payload of the reply as a UnsubscribeResponse.
func Call_Unsubscribe(ctx context.Context, dev *networking.Client, request Unsubscribe) (UnsubscribeResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			UnsubscribeResponse UnsubscribeResponse
		}
	}
	reply := Envelope{}
	httpReply, err := dev.CallMethod(ctx, request)
	if httpReply != nil {
		defer httpReply.Body.Close()
	}
	if err != nil {
		return reply.Body.UnsubscribeResponse, err
	} else {
		err = networking.ReadAndParse(httpReply, &reply, "Unsubscribe")
		return reply.Body.UnsubscribeResponse, err
	}
}
