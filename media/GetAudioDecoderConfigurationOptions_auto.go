// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

// Code generated : DO NOT EDIT.

package media

import (
	"context"
	"github.com/ontavu/onvif/v2/networking"
)

// Call_GetAudioDecoderConfigurationOptions forwards the call to dev.CallMethod() then parses the payload of the reply as a GetAudioDecoderConfigurationOptionsResponse.
func Call_GetAudioDecoderConfigurationOptions(ctx context.Context, dev *networking.Client, request GetAudioDecoderConfigurationOptions) (GetAudioDecoderConfigurationOptionsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetAudioDecoderConfigurationOptionsResponse GetAudioDecoderConfigurationOptionsResponse
		}
	}
	reply := Envelope{}
	httpReply, err := dev.CallMethod(ctx, request)
	if httpReply != nil {
		defer httpReply.Body.Close()
	}
	if err != nil {
		return reply.Body.GetAudioDecoderConfigurationOptionsResponse, err
	} else {
		err = networking.ReadAndParse(httpReply, &reply, "GetAudioDecoderConfigurationOptions")
		return reply.Body.GetAudioDecoderConfigurationOptionsResponse, err
	}
}
