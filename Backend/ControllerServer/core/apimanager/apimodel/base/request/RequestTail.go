package request

type requestTail struct {
	requestTo         string
	requestDesIP      string
	requestOutputType string
}

func GetNewRequestTail() requestTail {
	return requestTail{}
}
func (requestTail *requestTail) GetRequestTailTo() string {
	return requestTail.requestTo
}

func (requestTail *requestTail) GetRequestTailDesIP() string {
	return requestTail.requestDesIP
}
func (requestTail *requestTail) GetRequestTailOutpputType() string {
	return requestTail.requestOutputType
}

func (requestTail *requestTail) SetRequestTailTo(to string) {
	requestTail.requestTo = to
}
func (requestTail *requestTail) SetRequestTailDesIP(ip string) {
	requestTail.requestDesIP = ip
}
func (requestTail *requestTail) SetRequestTailOutputType(typeRequest string) {
	requestTail.requestOutputType = typeRequest
}

func (request *requestTail) ToString() string {
	returnString := ""
	returnString += "To: " + request.requestTo + "\n"
	returnString += "Des IP: " + request.requestDesIP + "\n"
	returnString += "Output Type: " + request.requestOutputType + "\n"
	return returnString
}
