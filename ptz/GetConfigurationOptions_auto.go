// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

// Code generated : DO NOT EDIT.

package ptz

import (
	"context"
	"github.com/ontavu/onvif/v2/networking"
)

// Call_GetConfigurationOptions forwards the call to dev.CallMethod() then parses the payload of the reply as a GetConfigurationOptionsResponse.
func Call_GetConfigurationOptions(ctx context.Context, dev *networking.Client, request GetConfigurationOptions) (GetConfigurationOptionsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetConfigurationOptionsResponse GetConfigurationOptionsResponse
		}
	}
	reply := Envelope{}
	httpReply, err := dev.CallMethod(ctx, request)
	if httpReply != nil {
		defer httpReply.Body.Close()
	}
	if err != nil {
		return reply.Body.GetConfigurationOptionsResponse, err
	} else {
		err = networking.ReadAndParse(httpReply, &reply, "GetConfigurationOptions")
		return reply.Body.GetConfigurationOptionsResponse, err
	}
}
