//go:build asan

package asanutils

// #include <stdio.h>
// #include <stdlib.h>
//
// void __lsan_do_leak_check(void);
import "C"

// CheckLeaks triggers the `lsan` Leaks Checker, but only when compiled with `asan` support
func CheckLeaks() {
	C.__lsan_do_leak_check()
}
