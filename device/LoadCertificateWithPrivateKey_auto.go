// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

// Code generated : DO NOT EDIT.

package device

import (
	"context"
	"github.com/jfsmig/onvif/networking"
)

// Call_LoadCertificateWithPrivateKey forwards the call to dev.CallMethod() then parses the payload of the reply as a LoadCertificateWithPrivateKeyResponse.
func Call_LoadCertificateWithPrivateKey(ctx context.Context, dev *networking.Client, request LoadCertificateWithPrivateKey) (LoadCertificateWithPrivateKeyResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			LoadCertificateWithPrivateKeyResponse LoadCertificateWithPrivateKeyResponse
		}
	}
	reply := Envelope{}
	httpReply, err := dev.CallMethod(ctx, request)
	if httpReply != nil {
		defer httpReply.Body.Close()
	}
	if err != nil {
		return reply.Body.LoadCertificateWithPrivateKeyResponse, err
	} else {
		err = networking.ReadAndParse(httpReply, &reply, "LoadCertificateWithPrivateKey")
		return reply.Body.LoadCertificateWithPrivateKeyResponse, err
	}
}
