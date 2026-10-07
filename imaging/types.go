// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// Portions of this file derive from the goonvif project, now use-go/onvif,
// originally distributed under the MIT License, see LICENSE.MIT,
// Copyright (c) 2018 Yakovlev Dmitry, Zhorzh Palanjyan, Crazybber.
//
// SPDX-License-Identifier: MIT

package imaging

import (
	"github.com/ontavu/onvif/v2/xsd"
	"github.com/ontavu/onvif/v2/xsd/onvif"
)

type GetServiceCapabilities struct {
	XMLName string `xml:"timg:GetServiceCapabilities"`
}

type GetImagingSettings struct {
	XMLName          string               `xml:"timg:GetImagingSettings"`
	VideoSourceToken onvif.ReferenceToken `xml:"timg:VideoSourceToken"`
}

type SetImagingSettings struct {
	XMLName          string                  `xml:"timg:SetImagingSettings"`
	VideoSourceToken onvif.ReferenceToken    `xml:"timg:VideoSourceToken"`
	ImagingSettings  onvif.ImagingSettings20 `xml:"timg:ImagingSettings"`
	ForcePersistence xsd.Boolean             `xml:"timg:ForcePersistence"`
}

type GetOptions struct {
	XMLName          string               `xml:"timg:GetOptions"`
	VideoSourceToken onvif.ReferenceToken `xml:"timg:VideoSourceToken"`
}

type Move struct {
	XMLName          string               `xml:"timg:Move"`
	VideoSourceToken onvif.ReferenceToken `xml:"timg:VideoSourceToken"`
	Focus            onvif.FocusMove      `xml:"timg:Focus"`
}

type GetMoveOptions struct {
	XMLName          string               `xml:"timg:GetMoveOptions"`
	VideoSourceToken onvif.ReferenceToken `xml:"timg:VideoSourceToken"`
}

type Stop struct {
	XMLName          string               `xml:"timg:Stop"`
	VideoSourceToken onvif.ReferenceToken `xml:"timg:VideoSourceToken"`
}

type GetStatus struct {
	XMLName          string               `xml:"timg:GetStatus"`
	VideoSourceToken onvif.ReferenceToken `xml:"timg:VideoSourceToken"`
}

type GetPresets struct {
	XMLName          string               `xml:"timg:GetPresets"`
	VideoSourceToken onvif.ReferenceToken `xml:"timg:VideoSourceToken"`
}

type GetCurrentPreset struct {
	XMLName          string               `xml:"timg:GetCurrentPreset"`
	VideoSourceToken onvif.ReferenceToken `xml:"timg:VideoSourceToken"`
}

type SetCurrentPreset struct {
	XMLName          string               `xml:"timg:SetCurrentPreset"`
	VideoSourceToken onvif.ReferenceToken `xml:"timg:VideoSourceToken"`
	PresetToken      onvif.ReferenceToken `xml:"timg:PresetToken"`
}
