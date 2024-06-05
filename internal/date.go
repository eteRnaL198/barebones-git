package internal

import "time"

func GetJst() (string, error) {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return "", err
	}
	return time.Now().In(jst).Format("Mon Jan 2 15:04:05 2006 -0700"), nil
}
