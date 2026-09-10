// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

// Code generated : DO NOT EDIT.

package media

import (
	"context"
	"github.com/jfsmig/onvif/networking"
)

// Call_RemoveVideoEncoderConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a RemoveVideoEncoderConfigurationResponse.
func Call_RemoveVideoEncoderConfiguration(ctx context.Context, dev *networking.Client, request RemoveVideoEncoderConfiguration) (RemoveVideoEncoderConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			RemoveVideoEncoderConfigurationResponse RemoveVideoEncoderConfigurationResponse
		}
	}
	reply := Envelope{}
	httpReply, err := dev.CallMethod(ctx, request)
	if httpReply != nil {
		defer httpReply.Body.Close()
	}
	if err != nil {
		return reply.Body.RemoveVideoEncoderConfigurationResponse, err
	} else {
		err = networking.ReadAndParse(httpReply, &reply, "RemoveVideoEncoderConfiguration")
		return reply.Body.RemoveVideoEncoderConfigurationResponse, err
	}
}
