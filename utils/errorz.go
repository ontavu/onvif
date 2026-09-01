// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

package utils

// constError is a string that satisfies error. Being a defined string type, it can be
// declared const: a sentinel that no importer can reassign, and that stays comparable so
// errors.Is keeps working through a wrapping chain.
type constError string

func (e constError) Error() string { return string(e) }

const (
	// ErrHTTP reports a non-200 reply from the device. Callers wrap it with the status,
	// so match it with errors.Is rather than ==.
	ErrHTTP = constError("http request error")

	// ErrNotOnvif reports a host that answered but does not speak ONVIF.
	ErrNotOnvif = constError("not an ONVIF device")

	// ErrNoService reports that the appliance advertises no endpoint for the service a
	// request was addressed to. It is a property of the device, not a failure of the
	// exchange: nothing was sent. Callers need to tell it apart from a rejected request
	// because ONVIF makes whole services conditional — a camera without PTZ answers
	// everything else perfectly well — so this is the error a conditional capability
	// yields, and it must not be reported as a fault.
	ErrNoService = constError("no endpoint for the service")
)
