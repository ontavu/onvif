// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

// Code generated : DO NOT EDIT.

package device

import (
	"context"
	"github.com/ontavu/onvif/v2/networking"
)

// Call_GetStorageConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a GetStorageConfigurationResponse.
func Call_GetStorageConfiguration(ctx context.Context, dev *networking.Client, request GetStorageConfiguration) (GetStorageConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetStorageConfigurationResponse GetStorageConfigurationResponse
		}
	}
	reply := Envelope{}
	httpReply, err := dev.CallMethod(ctx, request)
	if httpReply != nil {
		defer httpReply.Body.Close()
	}
	if err != nil {
		return reply.Body.GetStorageConfigurationResponse, err
	} else {
		err = networking.ReadAndParse(httpReply, &reply, "GetStorageConfiguration")
		return reply.Body.GetStorageConfigurationResponse, err
	}
}
