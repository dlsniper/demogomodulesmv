package demogomodulesmv

import "github.com/dlsniper/demogomodulesmv/v2/bad"

var Message = "Hello World"

func init() {
	Message = bad.Message()
}
