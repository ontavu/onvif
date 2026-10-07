// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

// Code generated : DO NOT EDIT.

package device

import (
	"context"
	"github.com/ontavu/onvif/v2/networking"
)

// Call_GetServices forwards the call to dev.CallMethod() then parses the payload of the reply as a GetServicesResponse.
func Call_GetServices(ctx context.Context, dev *networking.Client, request GetServices) (GetServicesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetServicesResponse GetServicesResponse
		}
	}
	reply := Envelope{}
	httpReply, err := dev.CallMethod(ctx, request)
	if httpReply != nil {
		defer httpReply.Body.Close()
	}
	if err != nil {
		return reply.Body.GetServicesResponse, err
	} else {
		err = networking.ReadAndParse(httpReply, &reply, "GetServices")
		return reply.Body.GetServicesResponse, err
	}
}
