// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

// Code generated : DO NOT EDIT.

package media

import (
	"context"
	"github.com/ontavu/onvif/v2/networking"
)

// Call_GetAudioSourceConfigurations forwards the call to dev.CallMethod() then parses the payload of the reply as a GetAudioSourceConfigurationsResponse.
func Call_GetAudioSourceConfigurations(ctx context.Context, dev *networking.Client, request GetAudioSourceConfigurations) (GetAudioSourceConfigurationsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetAudioSourceConfigurationsResponse GetAudioSourceConfigurationsResponse
		}
	}
	reply := Envelope{}
	httpReply, err := dev.CallMethod(ctx, request)
	if httpReply != nil {
		defer httpReply.Body.Close()
	}
	if err != nil {
		return reply.Body.GetAudioSourceConfigurationsResponse, err
	} else {
		err = networking.ReadAndParse(httpReply, &reply, "GetAudioSourceConfigurations")
		return reply.Body.GetAudioSourceConfigurationsResponse, err
	}
}
