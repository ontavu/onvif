// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

// Code generated : DO NOT EDIT.

package device

import (
	"context"
	"github.com/jfsmig/onvif/networking"
)

// Call_CreateDot1XConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a CreateDot1XConfigurationResponse.
func Call_CreateDot1XConfiguration(ctx context.Context, dev *networking.Client, request CreateDot1XConfiguration) (CreateDot1XConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			CreateDot1XConfigurationResponse CreateDot1XConfigurationResponse
		}
	}
	reply := Envelope{}
	httpReply, err := dev.CallMethod(ctx, request)
	if httpReply != nil {
		defer httpReply.Body.Close()
	}
	if err != nil {
		return reply.Body.CreateDot1XConfigurationResponse, err
	} else {
		err = networking.ReadAndParse(httpReply, &reply, "CreateDot1XConfiguration")
		return reply.Body.CreateDot1XConfigurationResponse, err
	}
}
