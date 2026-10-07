// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

// Code generated : DO NOT EDIT.

package device

import (
	"context"
	"github.com/ontavu/onvif/v2/networking"
)

// Call_SetNetworkProtocols forwards the call to dev.CallMethod() then parses the payload of the reply as a SetNetworkProtocolsResponse.
func Call_SetNetworkProtocols(ctx context.Context, dev *networking.Client, request SetNetworkProtocols) (SetNetworkProtocolsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetNetworkProtocolsResponse SetNetworkProtocolsResponse
		}
	}
	reply := Envelope{}
	httpReply, err := dev.CallMethod(ctx, request)
	if httpReply != nil {
		defer httpReply.Body.Close()
	}
	if err != nil {
		return reply.Body.SetNetworkProtocolsResponse, err
	} else {
		err = networking.ReadAndParse(httpReply, &reply, "SetNetworkProtocols")
		return reply.Body.SetNetworkProtocolsResponse, err
	}
}
