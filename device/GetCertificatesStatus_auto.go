// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

// Code generated : DO NOT EDIT.

package device

import (
	"context"
	"github.com/ontavu/onvif/v2/networking"
)

// Call_GetCertificatesStatus forwards the call to dev.CallMethod() then parses the payload of the reply as a GetCertificatesStatusResponse.
func Call_GetCertificatesStatus(ctx context.Context, dev *networking.Client, request GetCertificatesStatus) (GetCertificatesStatusResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetCertificatesStatusResponse GetCertificatesStatusResponse
		}
	}
	reply := Envelope{}
	httpReply, err := dev.CallMethod(ctx, request)
	if httpReply != nil {
		defer httpReply.Body.Close()
	}
	if err != nil {
		return reply.Body.GetCertificatesStatusResponse, err
	} else {
		err = networking.ReadAndParse(httpReply, &reply, "GetCertificatesStatus")
		return reply.Body.GetCertificatesStatusResponse, err
	}
}
