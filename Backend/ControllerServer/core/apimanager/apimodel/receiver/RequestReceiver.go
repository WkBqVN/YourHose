package receiver

import (
	"controllerserver/core/apimanager/apimodel/base/request"
)

type RequestReceiver struct {
	requestObj request.RequestBase
	timeIn     string
	isValid    bool
}

func GetNewRequestReceiver() RequestReceiver {
	return RequestReceiver{}
}

func (requestReceiver *RequestReceiver) GetRequestObj() *request.RequestBase {
	return &requestReceiver.requestObj
}

func (requestReceiver *RequestReceiver) GetTimeIn() string {
	return requestReceiver.timeIn
}

func (requestReceiver *RequestReceiver) GetValid() bool {
	return requestReceiver.isValid
}

func (requestReceiver *RequestReceiver) SetRequestObj(request request.RequestBase) {
	requestReceiver.requestObj = request
}
func (requestReceiver *RequestReceiver) SetTimeIn(time string) {
	requestReceiver.timeIn = time
}
func (requestReceiver *RequestReceiver) SetValid(isValid bool) {
	requestReceiver.isValid = isValid
}
