// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

// Code generated : DO NOT EDIT.

package device

import (
	"context"
	"github.com/jfsmig/onvif/networking"
)

// Call_SetZeroConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a SetZeroConfigurationResponse.
func Call_SetZeroConfiguration(ctx context.Context, dev *networking.Client, request SetZeroConfiguration) (SetZeroConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetZeroConfigurationResponse SetZeroConfigurationResponse
		}
	}
	reply := Envelope{}
	httpReply, err := dev.CallMethod(ctx, request)
	if httpReply != nil {
		defer httpReply.Body.Close()
	}
	if err != nil {
		return reply.Body.SetZeroConfigurationResponse, err
	} else {
		err = networking.ReadAndParse(httpReply, &reply, "SetZeroConfiguration")
		return reply.Body.SetZeroConfigurationResponse, err
	}
}
