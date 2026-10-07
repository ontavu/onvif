// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

// Code generated : DO NOT EDIT.

package device

import (
	"context"
	"github.com/ontavu/onvif/v2/networking"
)

// Call_DeleteCertificates forwards the call to dev.CallMethod() then parses the payload of the reply as a DeleteCertificatesResponse.
func Call_DeleteCertificates(ctx context.Context, dev *networking.Client, request DeleteCertificates) (DeleteCertificatesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			DeleteCertificatesResponse DeleteCertificatesResponse
		}
	}
	reply := Envelope{}
	httpReply, err := dev.CallMethod(ctx, request)
	if httpReply != nil {
		defer httpReply.Body.Close()
	}
	if err != nil {
		return reply.Body.DeleteCertificatesResponse, err
	} else {
		err = networking.ReadAndParse(httpReply, &reply, "DeleteCertificates")
		return reply.Body.DeleteCertificatesResponse, err
	}
}
