// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

// Code generated : DO NOT EDIT.

package device

import (
	"context"
	"github.com/jfsmig/onvif/networking"
)

// Call_GetCertificates forwards the call to dev.CallMethod() then parses the payload of the reply as a GetCertificatesResponse.
func Call_GetCertificates(ctx context.Context, dev *networking.Client, request GetCertificates) (GetCertificatesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetCertificatesResponse GetCertificatesResponse
		}
	}
	reply := Envelope{}
	httpReply, err := dev.CallMethod(ctx, request)
	if httpReply != nil {
		defer httpReply.Body.Close()
	}
	if err != nil {
		return reply.Body.GetCertificatesResponse, err
	} else {
		err = networking.ReadAndParse(httpReply, &reply, "GetCertificates")
		return reply.Body.GetCertificatesResponse, err
	}
}
