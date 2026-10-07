// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

// Code generated : DO NOT EDIT.

package media

import (
	"context"
	"github.com/ontavu/onvif/v2/networking"
)

// Call_GetAudioSourceConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a GetAudioSourceConfigurationResponse.
func Call_GetAudioSourceConfiguration(ctx context.Context, dev *networking.Client, request GetAudioSourceConfiguration) (GetAudioSourceConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetAudioSourceConfigurationResponse GetAudioSourceConfigurationResponse
		}
	}
	reply := Envelope{}
	httpReply, err := dev.CallMethod(ctx, request)
	if httpReply != nil {
		defer httpReply.Body.Close()
	}
	if err != nil {
		return reply.Body.GetAudioSourceConfigurationResponse, err
	} else {
		err = networking.ReadAndParse(httpReply, &reply, "GetAudioSourceConfiguration")
		return reply.Body.GetAudioSourceConfigurationResponse, err
	}
}
