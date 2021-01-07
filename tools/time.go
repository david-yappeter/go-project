package tools

import "time"

//TimeUTC Get Time UTC Now
func TimeUTC() string {
	return time.Now().UTC().Format("2006-01-02 15:04:05")
}
