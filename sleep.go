//go:build cgo

package cbutil

// #include <unistd.h>
// int usleep(useconds_t usec);
// unsigned int sleep(unsigned int seconds);
import "C"
import "time"

func Sleep(sleepfor time.Duration) {
	if sleepfor < 0 {
		return
	}
	if sleepfor < time.Second {
		C.usleep(C.uint(sleepfor / 1000))
	} else {
		C.sleep(C.uint(sleepfor / 1000000000))
	}
}
