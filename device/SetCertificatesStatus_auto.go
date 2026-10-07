// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

// Code generated : DO NOT EDIT.

package device

import (
	"context"
	"github.com/ontavu/onvif/v2/networking"
)

// Call_SetCertificatesStatus forwards the call to dev.CallMethod() then parses the payload of the reply as a SetCertificatesStatusResponse.
func Call_SetCertificatesStatus(ctx context.Context, dev *networking.Client, request SetCertificatesStatus) (SetCertificatesStatusResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetCertificatesStatusResponse SetCertificatesStatusResponse
		}
	}
	reply := Envelope{}
	httpReply, err := dev.CallMethod(ctx, request)
	if httpReply != nil {
		defer httpReply.Body.Close()
	}
	if err != nil {
		return reply.Body.SetCertificatesStatusResponse, err
	} else {
		err = networking.ReadAndParse(httpReply, &reply, "SetCertificatesStatus")
		return reply.Body.SetCertificatesStatusResponse, err
	}
}
