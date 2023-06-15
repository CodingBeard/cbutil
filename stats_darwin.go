//go:build cgo

package cbutil

// #include <unistd.h>
import "C"

func GetTotalMemory() int64 {
	return int64(C.sysconf(C._SC_PHYS_PAGES) * C.sysconf(C._SC_PAGE_SIZE))
}
