package utils

import "time"

// GetCurrentTimeInJakarta returns the current time in Asia/Jakarta timezone (GMT+7)
func GetCurrentTimeInJakarta() time.Time {
    loc, _ := time.LoadLocation("Asia/Jakarta")
    return time.Now().In(loc)
}