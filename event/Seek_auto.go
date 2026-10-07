// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

// Code generated : DO NOT EDIT.

package event

import (
	"context"
	"github.com/ontavu/onvif/v2/networking"
)

// Call_Seek forwards the call to dev.CallMethod() then parses the payload of the reply as a SeekResponse.
func Call_Seek(ctx context.Context, dev *networking.Client, request Seek) (SeekResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SeekResponse SeekResponse
		}
	}
	reply := Envelope{}
	httpReply, err := dev.CallMethod(ctx, request)
	if httpReply != nil {
		defer httpReply.Body.Close()
	}
	if err != nil {
		return reply.Body.SeekResponse, err
	} else {
		err = networking.ReadAndParse(httpReply, &reply, "Seek")
		return reply.Body.SeekResponse, err
	}
}
