// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

// Code generated : DO NOT EDIT.

package media

import (
	"context"
	"github.com/jfsmig/onvif/networking"
)

// Call_GetCompatibleVideoSourceConfigurations forwards the call to dev.CallMethod() then parses the payload of the reply as a GetCompatibleVideoSourceConfigurationsResponse.
func Call_GetCompatibleVideoSourceConfigurations(ctx context.Context, dev *networking.Client, request GetCompatibleVideoSourceConfigurations) (GetCompatibleVideoSourceConfigurationsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetCompatibleVideoSourceConfigurationsResponse GetCompatibleVideoSourceConfigurationsResponse
		}
	}
	reply := Envelope{}
	httpReply, err := dev.CallMethod(ctx, request)
	if httpReply != nil {
		defer httpReply.Body.Close()
	}
	if err != nil {
		return reply.Body.GetCompatibleVideoSourceConfigurationsResponse, err
	} else {
		err = networking.ReadAndParse(httpReply, &reply, "GetCompatibleVideoSourceConfigurations")
		return reply.Body.GetCompatibleVideoSourceConfigurationsResponse, err
	}
}
