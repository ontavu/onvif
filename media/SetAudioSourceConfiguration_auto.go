// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

// Code generated : DO NOT EDIT.

package media

import (
	"context"
	"github.com/jfsmig/onvif/networking"
)

// Call_SetAudioSourceConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a SetAudioSourceConfigurationResponse.
func Call_SetAudioSourceConfiguration(ctx context.Context, dev *networking.Client, request SetAudioSourceConfiguration) (SetAudioSourceConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetAudioSourceConfigurationResponse SetAudioSourceConfigurationResponse
		}
	}
	reply := Envelope{}
	httpReply, err := dev.CallMethod(ctx, request)
	if httpReply != nil {
		defer httpReply.Body.Close()
	}
	if err != nil {
		return reply.Body.SetAudioSourceConfigurationResponse, err
	} else {
		err = networking.ReadAndParse(httpReply, &reply, "SetAudioSourceConfiguration")
		return reply.Body.SetAudioSourceConfigurationResponse, err
	}
}
