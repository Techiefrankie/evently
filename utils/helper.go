package utils

import "evently/api"

func GetResponse(msg string, code int) api.Response {
	return api.Response{
		Message:    msg,
		StatusCode: code,
	}
}
