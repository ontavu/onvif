// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

// Code generated : DO NOT EDIT.

package device

import (
	"context"
	"github.com/ontavu/onvif/v2/networking"
)

// Call_SetHostname forwards the call to dev.CallMethod() then parses the payload of the reply as a SetHostnameResponse.
func Call_SetHostname(ctx context.Context, dev *networking.Client, request SetHostname) (SetHostnameResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetHostnameResponse SetHostnameResponse
		}
	}
	reply := Envelope{}
	httpReply, err := dev.CallMethod(ctx, request)
	if httpReply != nil {
		defer httpReply.Body.Close()
	}
	if err != nil {
		return reply.Body.SetHostnameResponse, err
	} else {
		err = networking.ReadAndParse(httpReply, &reply, "SetHostname")
		return reply.Body.SetHostnameResponse, err
	}
}
