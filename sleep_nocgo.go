//go:build !cgo

package cbutil

import "time"

func Sleep(sleepfor time.Duration) {
	time.Sleep(sleepfor)
}
