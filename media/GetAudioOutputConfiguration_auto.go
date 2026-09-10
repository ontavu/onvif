// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

// Code generated : DO NOT EDIT.

package media

import (
	"context"
	"github.com/jfsmig/onvif/networking"
)

// Call_GetAudioOutputConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a GetAudioOutputConfigurationResponse.
func Call_GetAudioOutputConfiguration(ctx context.Context, dev *networking.Client, request GetAudioOutputConfiguration) (GetAudioOutputConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetAudioOutputConfigurationResponse GetAudioOutputConfigurationResponse
		}
	}
	reply := Envelope{}
	httpReply, err := dev.CallMethod(ctx, request)
	if httpReply != nil {
		defer httpReply.Body.Close()
	}
	if err != nil {
		return reply.Body.GetAudioOutputConfigurationResponse, err
	} else {
		err = networking.ReadAndParse(httpReply, &reply, "GetAudioOutputConfiguration")
		return reply.Body.GetAudioOutputConfigurationResponse, err
	}
}
