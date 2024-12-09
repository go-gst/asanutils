//go:build asan

package asanutils

// #include <stdio.h>
// #include <stdlib.h>
//
// void __lsan_do_leak_check(void);
import "C"
import "runtime"

// CheckLeaks triggers the `lsan` Leaks Checker, but only when compiled with `asan` support
func CheckLeaks() {
	// finalizers are funky, so we need to call the GC twice.
	// The reason is that a finalizer revives the object and cleans it in the next GC cycle.
	runtime.GC()
	runtime.GC()

	runtime.GC() // one more for good measure :)
	C.__lsan_do_leak_check()
}
