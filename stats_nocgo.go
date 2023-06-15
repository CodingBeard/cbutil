//go:build !cgo

package cbutil

func GetTotalMemory() int64 {
	return 0
}
