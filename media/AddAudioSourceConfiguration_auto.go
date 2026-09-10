// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

// Code generated : DO NOT EDIT.

package media

import (
	"context"
	"github.com/jfsmig/onvif/networking"
)

// Call_AddAudioSourceConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a AddAudioSourceConfigurationResponse.
func Call_AddAudioSourceConfiguration(ctx context.Context, dev *networking.Client, request AddAudioSourceConfiguration) (AddAudioSourceConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			AddAudioSourceConfigurationResponse AddAudioSourceConfigurationResponse
		}
	}
	reply := Envelope{}
	httpReply, err := dev.CallMethod(ctx, request)
	if httpReply != nil {
		defer httpReply.Body.Close()
	}
	if err != nil {
		return reply.Body.AddAudioSourceConfigurationResponse, err
	} else {
		err = networking.ReadAndParse(httpReply, &reply, "AddAudioSourceConfiguration")
		return reply.Body.AddAudioSourceConfigurationResponse, err
	}
}
