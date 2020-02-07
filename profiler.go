package cbutil

import (
	"fmt"
	"time"
)

type timestamp struct {
	step      string
	timestamp int64
}

var timestamps []timestamp

func ProfileStep(enabled bool, step string) {
	if !enabled {
		return
	}
	timestamps = append(timestamps, timestamp{
		step:      step,
		timestamp: time.Now().UnixNano(),
	})
}

func ProfileReset(enabled bool) {
	if !enabled {
		return
	}

	timestamps = timestamps[:0]
}

func ProfileGetResults(enabled bool) []string {
	if !enabled {
		return []string{}
	}
	var results []string

	lastTimestamp := int64(0)
	for _, timestamp := range timestamps {
		results = append(results, fmt.Sprintf("%fms +%fms: %s", float64(timestamp.timestamp)/float64(1000000), float64(timestamp.timestamp-lastTimestamp)/float64(1000000), timestamp.step))
		lastTimestamp = timestamp.timestamp
	}

	return results
}
