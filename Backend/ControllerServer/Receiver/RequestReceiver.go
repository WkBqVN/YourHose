package Receiver

type RequestReceiver struct {
	RequestIP string
	RequestType string
	RequestDes string
	RequestBody interface{}
}

func (request RequestReceiver) validateRequest() bool{
	isValid := true
	if request.RequestIP == ""{
		isValid = false
	}
	if request.RequestType == ""{
		isValid = false
	}
	return isValid
}


