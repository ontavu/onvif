// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

// Code generated : DO NOT EDIT.

package device

import (
	"context"
	"github.com/ontavu/onvif/v2/networking"
)

// Call_GetNetworkInterfaces forwards the call to dev.CallMethod() then parses the payload of the reply as a GetNetworkInterfacesResponse.
func Call_GetNetworkInterfaces(ctx context.Context, dev *networking.Client, request GetNetworkInterfaces) (GetNetworkInterfacesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetNetworkInterfacesResponse GetNetworkInterfacesResponse
		}
	}
	reply := Envelope{}
	httpReply, err := dev.CallMethod(ctx, request)
	if httpReply != nil {
		defer httpReply.Body.Close()
	}
	if err != nil {
		return reply.Body.GetNetworkInterfacesResponse, err
	} else {
		err = networking.ReadAndParse(httpReply, &reply, "GetNetworkInterfaces")
		return reply.Body.GetNetworkInterfacesResponse, err
	}
}
