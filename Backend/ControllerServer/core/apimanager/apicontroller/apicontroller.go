package apicontroller

import (
	"controllerserver/core/apimanager/apimodel/receiver"
	"controllerserver/core/apimanager/apimodel/sender"
)

type Controller struct {
	receirver  receiver.RequestReceiver
	controller int
	sender     sender.RequestSender
	dataMap    map[string]string
}

func getController() Controller {
	return Controller{}
}
