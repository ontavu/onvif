// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

// Code generated : DO NOT EDIT.

package ptz

import (
	"context"
	"github.com/ontavu/onvif/v2/networking"
)

// Call_SetConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a SetConfigurationResponse.
func Call_SetConfiguration(ctx context.Context, dev *networking.Client, request SetConfiguration) (SetConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetConfigurationResponse SetConfigurationResponse
		}
	}
	reply := Envelope{}
	httpReply, err := dev.CallMethod(ctx, request)
	if httpReply != nil {
		defer httpReply.Body.Close()
	}
	if err != nil {
		return reply.Body.SetConfigurationResponse, err
	} else {
		err = networking.ReadAndParse(httpReply, &reply, "SetConfiguration")
		return reply.Body.SetConfigurationResponse, err
	}
}
