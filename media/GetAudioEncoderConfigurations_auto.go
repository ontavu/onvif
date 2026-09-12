// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

// Code generated : DO NOT EDIT.

package media

import (
	"context"
	"github.com/jfsmig/onvif/v2/networking"
)

// Call_GetAudioEncoderConfigurations forwards the call to dev.CallMethod() then parses the payload of the reply as a GetAudioEncoderConfigurationsResponse.
func Call_GetAudioEncoderConfigurations(ctx context.Context, dev *networking.Client, request GetAudioEncoderConfigurations) (GetAudioEncoderConfigurationsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetAudioEncoderConfigurationsResponse GetAudioEncoderConfigurationsResponse
		}
	}
	reply := Envelope{}
	httpReply, err := dev.CallMethod(ctx, request)
	if httpReply != nil {
		defer httpReply.Body.Close()
	}
	if err != nil {
		return reply.Body.GetAudioEncoderConfigurationsResponse, err
	} else {
		err = networking.ReadAndParse(httpReply, &reply, "GetAudioEncoderConfigurations")
		return reply.Body.GetAudioEncoderConfigurationsResponse, err
	}
}
