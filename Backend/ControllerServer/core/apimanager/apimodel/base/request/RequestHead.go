package request

type requestHead struct {
	requestFrom      string
	requestSourceIP  string
	requestInputType string
	requestMethod    string
}

func (request *requestHead) GetRequestHead() *requestHead {
	return request
}
func (requestHead *requestHead) GetRequestMethod() string {
	return requestHead.requestMethod
}
func (requestHead *requestHead) GetRequestHeadSourceIP() string {
	return requestHead.requestSourceIP
}

func (requestHead *requestHead) GetRequestHeadInputType() string {
	return requestHead.requestInputType
}
func (requestHead *requestHead) GetRequestHeadFrom() string {
	return requestHead.requestSourceIP
}
func (requestHead *requestHead) SetRequestHeadIP(ip string) {
	requestHead.requestSourceIP = ip
}
func (requestHead *requestHead) SetRequestHeadInputType(typeRequest string) {
	requestHead.requestInputType = typeRequest
}

func (requestHead *requestHead) SetRequestHeadFrom(from string) {
	requestHead.requestFrom = from
}
func (requestHead *requestHead) SetRequestMethod(method string) {
	requestHead.requestMethod = method
}
func (request *requestHead) ToString() string {
	returnString := ""
	returnString += "From: " + request.requestFrom + "\n"
	returnString += "Source IP: " + request.requestFrom + "\n"
	returnString += "Method: " + request.requestFrom + "\n"
	returnString += "Type: " + request.requestFrom + "\n"
	return returnString
}
