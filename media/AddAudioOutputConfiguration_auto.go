// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

// Code generated : DO NOT EDIT.

package media

import (
	"context"
	"github.com/ontavu/onvif/v2/networking"
)

// Call_AddAudioOutputConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a AddAudioOutputConfigurationResponse.
func Call_AddAudioOutputConfiguration(ctx context.Context, dev *networking.Client, request AddAudioOutputConfiguration) (AddAudioOutputConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			AddAudioOutputConfigurationResponse AddAudioOutputConfigurationResponse
		}
	}
	reply := Envelope{}
	httpReply, err := dev.CallMethod(ctx, request)
	if httpReply != nil {
		defer httpReply.Body.Close()
	}
	if err != nil {
		return reply.Body.AddAudioOutputConfigurationResponse, err
	} else {
		err = networking.ReadAndParse(httpReply, &reply, "AddAudioOutputConfiguration")
		return reply.Body.AddAudioOutputConfigurationResponse, err
	}
}
