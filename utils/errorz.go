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
)
