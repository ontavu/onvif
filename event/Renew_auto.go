// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

// Code generated : DO NOT EDIT.

package event

import (
	"context"
	"github.com/jfsmig/onvif/v2/networking"
)

// Call_Renew forwards the call to dev.CallMethod() then parses the payload of the reply as a RenewResponse.
func Call_Renew(ctx context.Context, dev *networking.Client, request Renew) (RenewResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			RenewResponse RenewResponse
		}
	}
	reply := Envelope{}
	httpReply, err := dev.CallMethod(ctx, request)
	if httpReply != nil {
		defer httpReply.Body.Close()
	}
	if err != nil {
		return reply.Body.RenewResponse, err
	} else {
		err = networking.ReadAndParse(httpReply, &reply, "Renew")
		return reply.Body.RenewResponse, err
	}
}
