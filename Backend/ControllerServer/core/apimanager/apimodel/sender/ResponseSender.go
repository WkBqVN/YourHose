package sender

import (
	"controllerserver/core/apimanager/apimodel/base/response"
)

type ResponseSender struct {
	requestObj response.ResponseBase
	timeIn     string
	timeOut    string
	message    string
}

func GetNewResponseSender() ResponseSender {
	return ResponseSender{}
}

func (requestSender *ResponseSender) GetRequestObj() *response.ResponseBase {
	return &requestSender.requestObj
}

func (requestSender *ResponseSender) GetTimeIn() string {
	return requestSender.timeIn
}

func (requestSender *ResponseSender) GetTimeOut() string {
	return requestSender.timeOut
}
func (requestSender *ResponseSender) GetMessage() string {
	return requestSender.message
}
func (requestSender *ResponseSender) SetRequestObj(response *response.ResponseBase) {
	requestSender.requestObj = *response
}

func (requestSender *ResponseSender) SetTimeIn(time string) {
	requestSender.timeIn = time
}

func (requestSender *ResponseSender) SetTimeOut(time string) {
	requestSender.timeOut = time
}
func (requestSender *ResponseSender) SetMessage(message string) {
	requestSender.message = message
}
