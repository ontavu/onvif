// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

// Code generated : DO NOT EDIT.

package device

import (
	"context"
	"github.com/ontavu/onvif/v2/networking"
)

// Call_SetStorageConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a SetStorageConfigurationResponse.
func Call_SetStorageConfiguration(ctx context.Context, dev *networking.Client, request SetStorageConfiguration) (SetStorageConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetStorageConfigurationResponse SetStorageConfigurationResponse
		}
	}
	reply := Envelope{}
	httpReply, err := dev.CallMethod(ctx, request)
	if httpReply != nil {
		defer httpReply.Body.Close()
	}
	if err != nil {
		return reply.Body.SetStorageConfigurationResponse, err
	} else {
		err = networking.ReadAndParse(httpReply, &reply, "SetStorageConfiguration")
		return reply.Body.SetStorageConfigurationResponse, err
	}
}
