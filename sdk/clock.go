// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

package sdk

import (
	"time"

	"github.com/jfsmig/onvif/v2/xsd/onvif"
)

// deviceClockOffset returns deviceUTC - local, or zero when the device reported no usable
// UTCDateTime.
//
// UTCDateTime and not LocalDateTime, for two reasons. The WS-Security Created stamp is
// formatted as an RFC3339 UTC instant, so UTC is what is wanted; and the only offset a
// device reports alongside LocalDateTime is TimeZone, a POSIX 1003.1 section 8.3 TZ string
// whose DST rules this package does not parse, so reconstructing UTC from the local time
// would be a guess. docs/wsdl/devicemgmt.wsdl:2688 requires UTCDateTime to be present.
//
// A zero Year means the optional element was absent or unparsed. The offset is then zero,
// which reproduces exactly the previous behaviour of stamping in local time, rather than an
// offset computed against year 0.
//
// There is deliberately no sanity clamp on the magnitude: a camera whose clock reads 1970
// wants its token stamped in 1970, and clamping would defeat the whole point.
func deviceClockOffset(local time.Time, sdt onvif.SystemDateTime) time.Duration {
	d := sdt.UTCDateTime
	if d.Date.Year == 0 {
		return 0
	}

	// xsd.Int is int32, hence the conversions.
	deviceUTC := time.Date(
		int(d.Date.Year), time.Month(d.Date.Month), int(d.Date.Day),
		int(d.Time.Hour), int(d.Time.Minute), int(d.Time.Second), 0, time.UTC)

	return deviceUTC.Sub(local.UTC())
}
