package receiver

import (
	"controllerserver/core/apimanager/apimodel/base/request"
)

type RequestReceiver struct {
	requestObj request.RequestBase
	timeIn int
	isValid bool
}

