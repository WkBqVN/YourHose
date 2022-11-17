package response

type ResponseBase struct {
	head responseHead
	body responseBody
	tail responseTail
}

func GetNewResponseBase() interface{} {
	return ResponseBase{}
}
func (response *ResponseBase) GetResponseBaseHead() responseHead {
	return response.head
}
func (response *ResponseBase) GetResponseBaseBody() responseBody {
	return response.body
}
func (response *ResponseBase) GetResponseBaseTail() responseTail {
	return response.tail
}
func (response *ResponseBase) SetResponseBaseHead(responseHead responseHead) {
	response.head = responseHead
}
func (response *ResponseBase) SetResponseBaseBody(responseBody responseBody) {
	response.body = responseBody
}
func (response *ResponseBase) SetResponseBaseTail(responseTail responseTail) {
	response.tail = responseTail
}
func (response *ResponseBase) SetResponseBase(responseNew ResponseBase) {
	*response = responseNew
}

func (response *ResponseBase) ToString() string {
	returnString := ""
	returnString += response.head.ToString()
	returnString += response.body.ToString()
	returnString += response.tail.ToString()
	return returnString
}

func (response *ResponseBase) HeadToString() string {
	return response.head.ToString()
}
func (response *ResponseBase) BodyToString() string {
	return response.body.ToString()
}
func (response *ResponseBase) TailToString() string {
	return response.tail.ToString()
}
