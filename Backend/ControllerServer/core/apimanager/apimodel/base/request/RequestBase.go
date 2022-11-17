package request

type RequestBase struct {
	head requestHead
	body requestBody
	tail requestTail
}

func GetNewRequestBase() RequestBase {
	return RequestBase{}
}

func (request *RequestBase) GetRequestBaseHead() requestHead {
	return request.head
}
func (request *RequestBase) GetRequestBaseBody() requestBody {
	return request.body
}
func (request *RequestBase) GetRequestBaseTail() requestTail {
	return request.tail
}
func (request *RequestBase) SetRequestBaseHead(requestHead requestHead) {
	request.head = requestHead
}
func (request *RequestBase) SetRequestBaseBody(requestBody requestBody) {
	request.body = requestBody
}
func (request *RequestBase) SetRequestBaseTail(requestTail requestTail) {
	request.tail = requestTail
}

func (request *RequestBase) SetRequestBase(requestNew RequestBase) {
	*request = requestNew
}

func (request *RequestBase) ToString() string {
	returnString := ""
	returnString += "Head: " + request.head.ToString() + "\n"
	returnString += "Body: " + request.body.ToString() + "\n"
	returnString += "Tail: " + request.tail.ToString() + "\n"
	return returnString
}

func (request *RequestBase) HeadToString() string {
	return request.head.ToString()
}
func (request *RequestBase) BodyToString() string {
	return request.tail.ToString()
}
func (request *RequestBase) TailToString() string {
	return request.body.ToString()
}
