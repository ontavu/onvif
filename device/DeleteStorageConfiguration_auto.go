// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

// Code generated : DO NOT EDIT.

package device

import (
	"context"
	"github.com/jfsmig/onvif/networking"
)

// Call_DeleteStorageConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a DeleteStorageConfigurationResponse.
func Call_DeleteStorageConfiguration(ctx context.Context, dev *networking.Client, request DeleteStorageConfiguration) (DeleteStorageConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			DeleteStorageConfigurationResponse DeleteStorageConfigurationResponse
		}
	}
	reply := Envelope{}
	httpReply, err := dev.CallMethod(ctx, request)
	if httpReply != nil {
		defer httpReply.Body.Close()
	}
	if err != nil {
		return reply.Body.DeleteStorageConfigurationResponse, err
	} else {
		err = networking.ReadAndParse(httpReply, &reply, "DeleteStorageConfiguration")
		return reply.Body.DeleteStorageConfigurationResponse, err
	}
}
