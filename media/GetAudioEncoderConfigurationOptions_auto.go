// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

// Code generated : DO NOT EDIT.

package media

import (
	"context"
	"github.com/jfsmig/onvif/networking"
)

// Call_GetAudioEncoderConfigurationOptions forwards the call to dev.CallMethod() then parses the payload of the reply as a GetAudioEncoderConfigurationOptionsResponse.
func Call_GetAudioEncoderConfigurationOptions(ctx context.Context, dev *networking.Client, request GetAudioEncoderConfigurationOptions) (GetAudioEncoderConfigurationOptionsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetAudioEncoderConfigurationOptionsResponse GetAudioEncoderConfigurationOptionsResponse
		}
	}
	reply := Envelope{}
	httpReply, err := dev.CallMethod(ctx, request)
	if httpReply != nil {
		defer httpReply.Body.Close()
	}
	if err != nil {
		return reply.Body.GetAudioEncoderConfigurationOptionsResponse, err
	} else {
		err = networking.ReadAndParse(httpReply, &reply, "GetAudioEncoderConfigurationOptions")
		return reply.Body.GetAudioEncoderConfigurationOptionsResponse, err
	}
}
