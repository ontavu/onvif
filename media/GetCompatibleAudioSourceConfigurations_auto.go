// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

// Code generated : DO NOT EDIT.

package media

import (
	"context"
	"github.com/ontavu/onvif/v2/networking"
)

// Call_GetCompatibleAudioSourceConfigurations forwards the call to dev.CallMethod() then parses the payload of the reply as a GetCompatibleAudioSourceConfigurationsResponse.
func Call_GetCompatibleAudioSourceConfigurations(ctx context.Context, dev *networking.Client, request GetCompatibleAudioSourceConfigurations) (GetCompatibleAudioSourceConfigurationsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetCompatibleAudioSourceConfigurationsResponse GetCompatibleAudioSourceConfigurationsResponse
		}
	}
	reply := Envelope{}
	httpReply, err := dev.CallMethod(ctx, request)
	if httpReply != nil {
		defer httpReply.Body.Close()
	}
	if err != nil {
		return reply.Body.GetCompatibleAudioSourceConfigurationsResponse, err
	} else {
		err = networking.ReadAndParse(httpReply, &reply, "GetCompatibleAudioSourceConfigurations")
		return reply.Body.GetCompatibleAudioSourceConfigurationsResponse, err
	}
}
