// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

// Code generated : DO NOT EDIT.

package media

import (
	"context"
	"github.com/ontavu/onvif/v2/networking"
)

// Call_SetVideoEncoderConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a SetVideoEncoderConfigurationResponse.
func Call_SetVideoEncoderConfiguration(ctx context.Context, dev *networking.Client, request SetVideoEncoderConfiguration) (SetVideoEncoderConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetVideoEncoderConfigurationResponse SetVideoEncoderConfigurationResponse
		}
	}
	reply := Envelope{}
	httpReply, err := dev.CallMethod(ctx, request)
	if httpReply != nil {
		defer httpReply.Body.Close()
	}
	if err != nil {
		return reply.Body.SetVideoEncoderConfigurationResponse, err
	} else {
		err = networking.ReadAndParse(httpReply, &reply, "SetVideoEncoderConfiguration")
		return reply.Body.SetVideoEncoderConfigurationResponse, err
	}
}
