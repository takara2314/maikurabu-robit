package main

import (
	"time"
)

func init() {
	// Using Asia/Tokyo timezone
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		loc = time.FixedZone("Asia/Tokyo", 9*60*60)
	}
	time.Local = loc
}
