package response

type responseHead struct {
	responseTo      string
	responseDesIP   string
	responseOutType string
	responseMethod  string
}

func GetNewResponseHead() responseHead {
	return responseHead{}
}
func (responseHead *responseHead) GetResponseHeadDesIP() string {
	return responseHead.responseDesIP
}
func (responseHead *responseHead) GetResonseHeadInputType() string {
	return responseHead.responseOutType
}
func (responseHead *responseHead) GetResponseHeadTo() string {
	return responseHead.responseTo
}
func (responseHead *responseHead) GetResponseMethod() string {
	return responseHead.responseMethod
}
func (responseHead *responseHead) SetResponseHeadDesIP(ip string) {
	responseHead.responseDesIP = ip
}
func (responseHead *responseHead) SetResponseHeadOutputType(typeRequest string) {
	responseHead.responseOutType = typeRequest
}
func (responseHead *responseHead) SetResponseHeadTo(to string) {
	responseHead.responseTo = to
}
func (responseHead *responseHead) SetResponseMethod(method string) {
	responseHead.responseMethod = method
}
func (responseHead *responseHead) SetResponseHead(responseNew responseHead) {
	*responseHead = responseNew
}
func (responseHead *responseHead) ToString() string {
	returnString := ""
	returnString += "To: " + responseHead.responseTo + "\n"
	returnString += "Des IP: " + responseHead.responseDesIP + "\n"
	returnString += "Method: " + responseHead.responseMethod + "\n"
	returnString += "Output Type: " + responseHead.responseOutType + "\n"
	return returnString
}
