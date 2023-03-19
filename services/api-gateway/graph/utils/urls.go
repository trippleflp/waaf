package utils

import "os"

var FunctionGroupServiceUrl = func() string {
	url, exist := os.LookupEnv("FUNCTIONGROUP_URL")
	if !exist {
		return "http://localhost:10001"
	}
	return url
}()

var AuthenticationServiceUrl = func() string {
	url, exist := os.LookupEnv("AUTHENTICATION_URL")
	if !exist {
		return "http://localhost:10002"
	}
	return url
}()
