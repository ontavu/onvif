// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

// Code generated : DO NOT EDIT.

package ptz

import (
	"context"
	"github.com/ontavu/onvif/v2/networking"
)

// Call_GotoPreset forwards the call to dev.CallMethod() then parses the payload of the reply as a GotoPresetResponse.
func Call_GotoPreset(ctx context.Context, dev *networking.Client, request GotoPreset) (GotoPresetResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GotoPresetResponse GotoPresetResponse
		}
	}
	reply := Envelope{}
	httpReply, err := dev.CallMethod(ctx, request)
	if httpReply != nil {
		defer httpReply.Body.Close()
	}
	if err != nil {
		return reply.Body.GotoPresetResponse, err
	} else {
		err = networking.ReadAndParse(httpReply, &reply, "GotoPreset")
		return reply.Body.GotoPresetResponse, err
	}
}
