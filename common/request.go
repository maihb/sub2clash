package common

import (
	"resty.dev/v3"
)

func Request(retryTimes int) *resty.Client {
	client := resty.New()
	client.
		SetRetryCount(retryTimes)
	return client
}
