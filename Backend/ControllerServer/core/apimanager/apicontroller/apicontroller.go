package apicontroller

import (
	"controllerserver/core/apimanager/apimodel/receiver"
	"controllerserver/core/apimanager/apimodel/sender"
)

type abc struct {
	receirver  receiver.RequestReceiver
	controller int
	sender     sender.RequestSender
	dataMap    map[string]string
}
