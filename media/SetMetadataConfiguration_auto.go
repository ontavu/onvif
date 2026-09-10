// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

// Code generated : DO NOT EDIT.

package media

import (
	"context"
	"github.com/jfsmig/onvif/networking"
)

// Call_SetMetadataConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a SetMetadataConfigurationResponse.
func Call_SetMetadataConfiguration(ctx context.Context, dev *networking.Client, request SetMetadataConfiguration) (SetMetadataConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetMetadataConfigurationResponse SetMetadataConfigurationResponse
		}
	}
	reply := Envelope{}
	httpReply, err := dev.CallMethod(ctx, request)
	if httpReply != nil {
		defer httpReply.Body.Close()
	}
	if err != nil {
		return reply.Body.SetMetadataConfigurationResponse, err
	} else {
		err = networking.ReadAndParse(httpReply, &reply, "SetMetadataConfiguration")
		return reply.Body.SetMetadataConfigurationResponse, err
	}
}
