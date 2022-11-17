package response

type responseTail struct {
	responseFrom       string
	responseSouceIP    string
	responseOutputType string
}

func GetNewResponseTail() responseTail {
	return responseTail{}
}
func (responseTail *responseTail) GetResponseTailFrom() string {
	return responseTail.responseFrom
}
func (responseTail *responseTail) GetResponseTailSourceIP() string {
	return responseTail.responseSouceIP
}
func (responseTail *responseTail) GetResponseTailOutpputType() string {
	return responseTail.responseOutputType
}
func (responseTail *responseTail) SetResponseTailFrom(from string) {
	responseTail.responseFrom = from
}
func (responseTail *responseTail) SetResponseTailDesIP(ip string) {
	responseTail.responseSouceIP = ip
}
func (responseTail *responseTail) SetResponseTailOutputType(typeRequest string) {
	responseTail.responseOutputType = typeRequest
}

func (responseTail *responseTail) SetResponseTail(responseNew responseTail) {
	*responseTail = responseNew
}

func (responseTail *responseTail) ToString() string {
	returnString := ""
	returnString += "From: " + responseTail.responseFrom + "\n"
	returnString += "Source IP: " + responseTail.responseSouceIP + "\n"
	returnString += "Output Type: " + responseTail.responseOutputType + "\n"
	return returnString
}
