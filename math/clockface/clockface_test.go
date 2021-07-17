package clockface

import (
	"math"
	"testing"
	"time"
)

func Test_secondsInRadians(t *testing.T) {
	tests := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 0, 30), math.Pi},
		{simpleTime(0, 0, 0), 0},
		{simpleTime(0, 0, 45), math.Pi / 2 * 3},
		{simpleTime(0, 0, 7), math.Pi / 30 * 7},
	}
	for _, test := range tests {
		t.Run(testName(test.time), func(t *testing.T) {
			got := secondsInRadians(test.time)
			if got != test.angle {
				t.Fatalf("want %v radians, got %v", test.angle, got)
			}
		})
	}
}

func simpleTime(h, m, s int) time.Time {
	return time.Date(312, time.October, 28, h, m, s, 0, time.UTC)
}

func testName(t time.Time) string {
	return t.Format("15:04:05")
}
