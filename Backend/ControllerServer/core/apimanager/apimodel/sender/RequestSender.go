package sender

import "controllerserver/core/apimanager/apimodel/base/response"

type RequestSender struct {
	requestObj response.ResponseBase
	timeIn     int
	timeOut    int
	message    string
}
