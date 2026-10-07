// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

// Code generated : DO NOT EDIT.

package media

import (
	"context"
	"github.com/ontavu/onvif/v2/networking"
)

// Call_RemoveAudioOutputConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a RemoveAudioOutputConfigurationResponse.
func Call_RemoveAudioOutputConfiguration(ctx context.Context, dev *networking.Client, request RemoveAudioOutputConfiguration) (RemoveAudioOutputConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			RemoveAudioOutputConfigurationResponse RemoveAudioOutputConfigurationResponse
		}
	}
	reply := Envelope{}
	httpReply, err := dev.CallMethod(ctx, request)
	if httpReply != nil {
		defer httpReply.Body.Close()
	}
	if err != nil {
		return reply.Body.RemoveAudioOutputConfigurationResponse, err
	} else {
		err = networking.ReadAndParse(httpReply, &reply, "RemoveAudioOutputConfiguration")
		return reply.Body.RemoveAudioOutputConfigurationResponse, err
	}
}
