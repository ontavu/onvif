// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

// Code generated : DO NOT EDIT.

package ptz

import (
	"context"
	"github.com/ontavu/onvif/v2/networking"
)

// Call_CreatePresetTour forwards the call to dev.CallMethod() then parses the payload of the reply as a CreatePresetTourResponse.
func Call_CreatePresetTour(ctx context.Context, dev *networking.Client, request CreatePresetTour) (CreatePresetTourResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			CreatePresetTourResponse CreatePresetTourResponse
		}
	}
	reply := Envelope{}
	httpReply, err := dev.CallMethod(ctx, request)
	if httpReply != nil {
		defer httpReply.Body.Close()
	}
	if err != nil {
		return reply.Body.CreatePresetTourResponse, err
	} else {
		err = networking.ReadAndParse(httpReply, &reply, "CreatePresetTour")
		return reply.Body.CreatePresetTourResponse, err
	}
}
