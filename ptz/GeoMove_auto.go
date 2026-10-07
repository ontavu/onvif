// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

// Code generated : DO NOT EDIT.

package ptz

import (
	"context"
	"github.com/ontavu/onvif/v2/networking"
)

// Call_GeoMove forwards the call to dev.CallMethod() then parses the payload of the reply as a GeoMoveResponse.
func Call_GeoMove(ctx context.Context, dev *networking.Client, request GeoMove) (GeoMoveResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GeoMoveResponse GeoMoveResponse
		}
	}
	reply := Envelope{}
	httpReply, err := dev.CallMethod(ctx, request)
	if httpReply != nil {
		defer httpReply.Body.Close()
	}
	if err != nil {
		return reply.Body.GeoMoveResponse, err
	} else {
		err = networking.ReadAndParse(httpReply, &reply, "GeoMove")
		return reply.Body.GeoMoveResponse, err
	}
}
