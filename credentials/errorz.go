// Copyright (C) 2022-2026 Jean-Francois SMIGIELSKI
//
// SPDX-License-Identifier: MIT

package credentials

// constError is the idiom of utils/errorz.go: a defined string type, so the sentinels
// below can be const -- no importer can reassign them -- while staying comparable, so
// errors.Is keeps working through a wrapping chain.
type constError string

func (e constError) Error() string { return string(e) }

const (
	// ErrMalformedFile reports a credentials file that could not be turned into entries:
	// invalid JSON, an empty file, an empty camera identifier, or an entry with no user.
	// It is wrapped with the file name and never with any of the file's content -- see
	// jsonFault for why the decoder's own message does not qualify.
	ErrMalformedFile = constError("malformed credentials file")

	// ErrDuplicateEntry reports one camera claimed twice in the same directory. An error
	// rather than a first-wins rule because the two entries may disagree, and settling
	// that by the order fs.ReadDir happens to return would make the effective password a
	// function of the file names.
	ErrDuplicateEntry = constError("duplicate camera identifier")
)
