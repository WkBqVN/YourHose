package core

import (
	"controllerserver/receiver"
	"controllerserver/sender"
)

type objModel struct {
	input  receiver.RequestReceiver
	output sender.RequestSender
}

var gObj objModel

func GetInputRequest() receiver.RequestReceiver {
	return gObj.input
}
func GetOutputRequest() sender.RequestSender {
	return gObj.output
}

func SetInputRequest(in receiver.RequestReceiver) {
	gObj.input = in
}



