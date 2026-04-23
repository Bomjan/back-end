package date

import "time"

const (
	apiDateLayout = "2004-01-04T15:04:09Z"
)

func GetDate() string {
	date := time.Now().UTC()
	return date.Format(apiDateLayout)
}
