// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

// Code generated : DO NOT EDIT.

package media

import (
	"context"
	"github.com/jfsmig/onvif/networking"
)

// Call_RemoveAudioEncoderConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a RemoveAudioEncoderConfigurationResponse.
func Call_RemoveAudioEncoderConfiguration(ctx context.Context, dev *networking.Client, request RemoveAudioEncoderConfiguration) (RemoveAudioEncoderConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			RemoveAudioEncoderConfigurationResponse RemoveAudioEncoderConfigurationResponse
		}
	}
	reply := Envelope{}
	httpReply, err := dev.CallMethod(ctx, request)
	if httpReply != nil {
		defer httpReply.Body.Close()
	}
	if err != nil {
		return reply.Body.RemoveAudioEncoderConfigurationResponse, err
	} else {
		err = networking.ReadAndParse(httpReply, &reply, "RemoveAudioEncoderConfiguration")
		return reply.Body.RemoveAudioEncoderConfigurationResponse, err
	}
}
