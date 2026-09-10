// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

// Code generated : DO NOT EDIT.

package media

import (
	"context"
	"github.com/jfsmig/onvif/networking"
)

// Call_DeleteProfile forwards the call to dev.CallMethod() then parses the payload of the reply as a DeleteProfileResponse.
func Call_DeleteProfile(ctx context.Context, dev *networking.Client, request DeleteProfile) (DeleteProfileResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			DeleteProfileResponse DeleteProfileResponse
		}
	}
	reply := Envelope{}
	httpReply, err := dev.CallMethod(ctx, request)
	if httpReply != nil {
		defer httpReply.Body.Close()
	}
	if err != nil {
		return reply.Body.DeleteProfileResponse, err
	} else {
		err = networking.ReadAndParse(httpReply, &reply, "DeleteProfile")
		return reply.Body.DeleteProfileResponse, err
	}
}
