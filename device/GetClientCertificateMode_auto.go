// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

// Code generated : DO NOT EDIT.

package device

import (
	"context"
	"github.com/jfsmig/onvif/networking"
)

// Call_GetClientCertificateMode forwards the call to dev.CallMethod() then parses the payload of the reply as a GetClientCertificateModeResponse.
func Call_GetClientCertificateMode(ctx context.Context, dev *networking.Client, request GetClientCertificateMode) (GetClientCertificateModeResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetClientCertificateModeResponse GetClientCertificateModeResponse
		}
	}
	reply := Envelope{}
	httpReply, err := dev.CallMethod(ctx, request)
	if httpReply != nil {
		defer httpReply.Body.Close()
	}
	if err != nil {
		return reply.Body.GetClientCertificateModeResponse, err
	} else {
		err = networking.ReadAndParse(httpReply, &reply, "GetClientCertificateMode")
		return reply.Body.GetClientCertificateModeResponse, err
	}
}
