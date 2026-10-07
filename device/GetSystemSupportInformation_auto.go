// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

// Code generated : DO NOT EDIT.

package device

import (
	"context"
	"github.com/ontavu/onvif/v2/networking"
)

// Call_GetSystemSupportInformation forwards the call to dev.CallMethod() then parses the payload of the reply as a GetSystemSupportInformationResponse.
func Call_GetSystemSupportInformation(ctx context.Context, dev *networking.Client, request GetSystemSupportInformation) (GetSystemSupportInformationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetSystemSupportInformationResponse GetSystemSupportInformationResponse
		}
	}
	reply := Envelope{}
	httpReply, err := dev.CallMethod(ctx, request)
	if httpReply != nil {
		defer httpReply.Body.Close()
	}
	if err != nil {
		return reply.Body.GetSystemSupportInformationResponse, err
	} else {
		err = networking.ReadAndParse(httpReply, &reply, "GetSystemSupportInformation")
		return reply.Body.GetSystemSupportInformationResponse, err
	}
}
