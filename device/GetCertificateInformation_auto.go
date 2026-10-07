// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

// Code generated : DO NOT EDIT.

package device

import (
	"context"
	"github.com/ontavu/onvif/v2/networking"
)

// Call_GetCertificateInformation forwards the call to dev.CallMethod() then parses the payload of the reply as a GetCertificateInformationResponse.
func Call_GetCertificateInformation(ctx context.Context, dev *networking.Client, request GetCertificateInformation) (GetCertificateInformationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetCertificateInformationResponse GetCertificateInformationResponse
		}
	}
	reply := Envelope{}
	httpReply, err := dev.CallMethod(ctx, request)
	if httpReply != nil {
		defer httpReply.Body.Close()
	}
	if err != nil {
		return reply.Body.GetCertificateInformationResponse, err
	} else {
		err = networking.ReadAndParse(httpReply, &reply, "GetCertificateInformation")
		return reply.Body.GetCertificateInformationResponse, err
	}
}
