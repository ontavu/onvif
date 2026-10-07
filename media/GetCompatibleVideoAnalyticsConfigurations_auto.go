// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

// Code generated : DO NOT EDIT.

package media

import (
	"context"
	"github.com/ontavu/onvif/v2/networking"
)

// Call_GetCompatibleVideoAnalyticsConfigurations forwards the call to dev.CallMethod() then parses the payload of the reply as a GetCompatibleVideoAnalyticsConfigurationsResponse.
func Call_GetCompatibleVideoAnalyticsConfigurations(ctx context.Context, dev *networking.Client, request GetCompatibleVideoAnalyticsConfigurations) (GetCompatibleVideoAnalyticsConfigurationsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetCompatibleVideoAnalyticsConfigurationsResponse GetCompatibleVideoAnalyticsConfigurationsResponse
		}
	}
	reply := Envelope{}
	httpReply, err := dev.CallMethod(ctx, request)
	if httpReply != nil {
		defer httpReply.Body.Close()
	}
	if err != nil {
		return reply.Body.GetCompatibleVideoAnalyticsConfigurationsResponse, err
	} else {
		err = networking.ReadAndParse(httpReply, &reply, "GetCompatibleVideoAnalyticsConfigurations")
		return reply.Body.GetCompatibleVideoAnalyticsConfigurationsResponse, err
	}
}
