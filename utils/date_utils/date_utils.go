package date_utils

import "time"

const (
	apiDateLayout   =  "2006-1-2"
)
func GetNowString() string {
	now := time.Now().UTC()
	return now.Format(apiDateLayout)
}