package utils

import (
	"strconv"
	"time"
)

func Sleep(duration time.Duration) {
	time.Sleep(duration)
}

func Test(base string) func(int) string {
	return func(n int) string {
		for i := 0; i < n; i++ {
			base += strconv.Itoa(i)
		}
		return base
	}
}
