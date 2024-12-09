//go:build !asan

package asanutils

// see also asan.go

// CheckLeaks triggers the `lsan` Leaks Checker, but only when compiled with `asan` support
func CheckLeaks() {
	// noop
}
