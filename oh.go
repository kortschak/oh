// Package oh provides an unrecoverable throw function.
package oh

import _ "unsafe"

// The name of the function provided by this package was intentionally
// chosen to be offensive to reduce the likelihood that someone would
// get its use through code review. It is not intended to be used in
// production code.

// Fuck is a linkname alias to runtime.throw.
//
//go:linkname Fuck runtime.throw
func Fuck(s string)
