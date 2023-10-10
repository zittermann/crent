package utils

import "time"

func GetDateTime() time.Time {
	return time.Now().Add(-time.Hour * 3) // Argentina date time
}
