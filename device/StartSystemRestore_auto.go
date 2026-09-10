// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

// Code generated : DO NOT EDIT.

package device

import (
	"context"
	"github.com/jfsmig/onvif/networking"
)

// Call_StartSystemRestore forwards the call to dev.CallMethod() then parses the payload of the reply as a StartSystemRestoreResponse.
func Call_StartSystemRestore(ctx context.Context, dev *networking.Client, request StartSystemRestore) (StartSystemRestoreResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			StartSystemRestoreResponse StartSystemRestoreResponse
		}
	}
	reply := Envelope{}
	httpReply, err := dev.CallMethod(ctx, request)
	if httpReply != nil {
		defer httpReply.Body.Close()
	}
	if err != nil {
		return reply.Body.StartSystemRestoreResponse, err
	} else {
		err = networking.ReadAndParse(httpReply, &reply, "StartSystemRestore")
		return reply.Body.StartSystemRestoreResponse, err
	}
}
