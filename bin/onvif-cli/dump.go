// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

package main

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"sync"

	"github.com/jfsmig/onvif/networking"
	"github.com/jfsmig/onvif/sdk"
)

type OnvifFullOutput struct {
	Descriptor     sdk.DeviceDescriptor
	DeviceSystem   sdk.DeviceSystem
	DeviceSecurity sdk.DeviceSecurity
	DeviceNetwork  sdk.DeviceNetwork
	Media          sdk.Media
	Ptz            sdk.Ptz
	Profiles       sdk.MediaProfiles
	Events         sdk.Event
}

type OnvifDeviceOutput struct {
	Descriptor     sdk.DeviceDescriptor
	DeviceSystem   sdk.DeviceSystem
	DeviceSecurity sdk.DeviceSecurity
	DeviceNetwork  sdk.DeviceNetwork
}

func dumpAll(ctx context.Context, params networking.ClientInfo) error {
	return dumpSomething(ctx, params, func(app sdk.Appliance, s *sdk.ProfileS) interface{} {
		out := OnvifFullOutput{}
		var wg sync.WaitGroup
		wg.Go(func() { out.Descriptor = s.FetchDeviceDescriptor(ctx) })
		wg.Go(func() { out.DeviceNetwork = s.FetchDeviceNetwork(ctx) })
		wg.Go(func() { out.DeviceSystem = s.FetchDeviceSystem(ctx) })
		wg.Go(func() { out.DeviceSecurity = s.FetchDeviceSecurity(ctx) })
		wg.Go(func() { out.Media = s.FetchMedia(ctx) })
		wg.Go(func() { out.Ptz = s.FetchPTZ(ctx) })
		wg.Go(func() { out.Events = s.FetchEvent(ctx) })
		wg.Go(func() { out.Profiles = s.FetchMediaProfiles(ctx) })
		wg.Wait()
		return out
	})
}

func dumpMedia(ctx context.Context, params networking.ClientInfo) error {
	return dumpSomething(ctx, params, func(app sdk.Appliance, s *sdk.ProfileS) interface{} {
		return s.FetchMedia(ctx)
	})
}

func dumpDescriptor(ctx context.Context, params networking.ClientInfo) error {
	type Output struct {
		Services   map[string]string
		UUID       string
		Descriptor sdk.DeviceDescriptor
	}
	return dumpSomething(ctx, params, func(app sdk.Appliance, s *sdk.ProfileS) interface{} {
		out := Output{}
		out.UUID = app.GetUUID()
		out.Services = app.GetServices()
		out.Descriptor = s.FetchDeviceDescriptor(ctx)
		return out
	})
}

func dumpPTZ(ctx context.Context, params networking.ClientInfo) error {
	return dumpSomething(ctx, params, func(app sdk.Appliance, s *sdk.ProfileS) interface{} {
		return s.FetchPTZ(ctx)
	})
}

func dumpEvents(ctx context.Context, params networking.ClientInfo) error {
	return dumpSomething(ctx, params, func(app sdk.Appliance, s *sdk.ProfileS) interface{} {
		return s.FetchEvent(ctx)
	})
}

func dumpDevice(ctx context.Context, params networking.ClientInfo) error {
	return dumpSomething(ctx, params, func(app sdk.Appliance, s *sdk.ProfileS) interface{} {
		out := OnvifDeviceOutput{}
		var wg sync.WaitGroup
		wg.Go(func() { out.Descriptor = s.FetchDeviceDescriptor(ctx) })
		wg.Go(func() { out.DeviceNetwork = s.FetchDeviceNetwork(ctx) })
		wg.Go(func() { out.DeviceSystem = s.FetchDeviceSystem(ctx) })
		wg.Go(func() { out.DeviceSecurity = s.FetchDeviceSecurity(ctx) })
		wg.Wait()
		return out
	})
}

func dumpProfiles(ctx context.Context, params networking.ClientInfo) error {
	return dumpSomething(ctx, params, func(app sdk.Appliance, s *sdk.ProfileS) interface{} {
		return s.FetchMediaProfiles(ctx)
	})
}

// ErrNoProfileS reports an appliance that does not advertise what ONVIF Profile S requires.
//
// Reported rather than shrugged off: every dump here is built from Profile S operations, so
// without that client there is nothing to print, and emitting an object full of zero values
// would look like a camera that answered with nothing rather than one that was never asked.
var ErrNoProfileS = errors.New("the appliance advertises no ONVIF Profile S services (device, media)")

func dumpSomething(ctx context.Context, params networking.ClientInfo, generate func(app sdk.Appliance, s *sdk.ProfileS) interface{}) error {
	sdkDev, err := sdk.NewDevice(ctx, params, auth, &httpClient)
	if err != nil {
		return err
	}
	profileS, ok := sdkDev.ProfileS()
	if !ok {
		return ErrNoProfileS
	}
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(generate(sdkDev, profileS))
}
