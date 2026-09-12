// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

package sdk

import (
	"context"
	"sync"

	"github.com/jfsmig/onvif/event"
)

type Event struct {
	// A pointer, so that a failed GetServiceCapabilities is null in a dump rather than a
	// struct of false bools, which reads as a camera that supports nothing. Same shape as
	// DeviceDescriptor.Capabilities, and the same reason.
	Capabilities *event.Capabilities
	Properties   event.GetEventPropertiesResponse
}

// FetchEvent issues its two operations concurrently.
//
// Each closure writes a different field of out, which is what makes the fan-out safe
// without a mutex, and it is the same argument the fan-outs in profiles.go rest on. Not
// errgroup: a camera without the events service errors on both calls, and
// first-error-cancels would turn a partial result into an empty one.
func (p *ProfileS) FetchEvent(ctx context.Context) Event {
	out := Event{}

	var wg sync.WaitGroup

	wg.Go(func() {
		if capa, err := event.Call_GetServiceCapabilities(ctx, p.client, event.GetServiceCapabilities{}); err == nil {
			out.Capabilities = &capa.Capabilities
		} else {
			Logger.Trace().Err(err).Str("rpc", "GetServiceCapabilities").Msg("event")
		}
	})

	wg.Go(func() {
		if props, err := event.Call_GetEventProperties(ctx, p.client, event.GetEventProperties{}); err == nil {
			out.Properties = props
		} else {
			Logger.Trace().Err(err).Str("rpc", "GetEventProperties").Msg("event")
		}
	})

	wg.Wait()
	return out
}
