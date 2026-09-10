// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

// Code generated : DO NOT EDIT.

package media

import (
	"context"
	"github.com/jfsmig/onvif/networking"
)

// Call_AddAudioEncoderConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a AddAudioEncoderConfigurationResponse.
func Call_AddAudioEncoderConfiguration(ctx context.Context, dev *networking.Client, request AddAudioEncoderConfiguration) (AddAudioEncoderConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			AddAudioEncoderConfigurationResponse AddAudioEncoderConfigurationResponse
		}
	}
	reply := Envelope{}
	httpReply, err := dev.CallMethod(ctx, request)
	if httpReply != nil {
		defer httpReply.Body.Close()
	}
	if err != nil {
		return reply.Body.AddAudioEncoderConfigurationResponse, err
	} else {
		err = networking.ReadAndParse(httpReply, &reply, "AddAudioEncoderConfiguration")
		return reply.Body.AddAudioEncoderConfigurationResponse, err
	}
}
