// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

// Code generated : DO NOT EDIT.

package media

import (
	"context"
	"github.com/jfsmig/onvif/networking"
)

// Call_GetAudioOutputConfigurationOptions forwards the call to dev.CallMethod() then parses the payload of the reply as a GetAudioOutputConfigurationOptionsResponse.
func Call_GetAudioOutputConfigurationOptions(ctx context.Context, dev *networking.Client, request GetAudioOutputConfigurationOptions) (GetAudioOutputConfigurationOptionsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetAudioOutputConfigurationOptionsResponse GetAudioOutputConfigurationOptionsResponse
		}
	}
	reply := Envelope{}
	httpReply, err := dev.CallMethod(ctx, request)
	if httpReply != nil {
		defer httpReply.Body.Close()
	}
	if err != nil {
		return reply.Body.GetAudioOutputConfigurationOptionsResponse, err
	} else {
		err = networking.ReadAndParse(httpReply, &reply, "GetAudioOutputConfigurationOptions")
		return reply.Body.GetAudioOutputConfigurationOptionsResponse, err
	}
}
