// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

// Code generated : DO NOT EDIT.

package device

import (
	"context"
	"github.com/jfsmig/onvif/networking"
)

// Call_SetIPAddressFilter forwards the call to dev.CallMethod() then parses the payload of the reply as a SetIPAddressFilterResponse.
func Call_SetIPAddressFilter(ctx context.Context, dev *networking.Client, request SetIPAddressFilter) (SetIPAddressFilterResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetIPAddressFilterResponse SetIPAddressFilterResponse
		}
	}
	reply := Envelope{}
	httpReply, err := dev.CallMethod(ctx, request)
	if httpReply != nil {
		defer httpReply.Body.Close()
	}
	if err != nil {
		return reply.Body.SetIPAddressFilterResponse, err
	} else {
		err = networking.ReadAndParse(httpReply, &reply, "SetIPAddressFilter")
		return reply.Body.SetIPAddressFilterResponse, err
	}
}
