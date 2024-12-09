# ASAN Utils for go

When compiling cgo to use in a go application, then a sanitizer is a good way to find common problems.

Go inludes support for [AddressSanitizer (asan)](https://github.com/google/sanitizers/wiki/addresssanitizer).

```bash
CC=clang CGO_ENABLED=1 CGO_LDFLAGS='-fsanitize=address' CGO_CFLAGS='-O0 -g3 -fsanitize=address' go build -asan <path>
```

This module adds the functions that are needed to trigger certain `asan` functionality on demand. Since they are only available when asan support is enabled, this module provides a noop fallback function if asan is not enabled.