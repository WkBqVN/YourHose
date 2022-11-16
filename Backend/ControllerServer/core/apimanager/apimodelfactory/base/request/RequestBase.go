package base

type RequestBase struct {
	head RequestHead
	body RequestBody
	tail RequestTail
}

type IRequestBase interface {
	getRequestBaseHead() RequestHead
	getRequestBaseBody() RequestBody
	getRequestBaseTail() RequestTail
	setRequestBaseHead(requestHead RequestHead)
	setRequestBaseBody(requestBody RequestBody)
	setRequestBaseTail(RequestTail RequestTail)
}

func (requestBase *RequestBase) getRequestBaseHead() RequestHead {
	return requestBase.head
}
func (requestBase *RequestBase) getRequestBaseBody() RequestBody {
	return requestBase.body
}
func (requestBase *RequestBase) getRequestBaseTail() RequestTail {
	return requestBase.tail
}
func (requestBase *RequestBase) setRequestBaseHead(requestHead RequestHead) {
	requestBase.head = requestHead
}
func (requestBase *RequestBase) setRequestBaseBody(requestBody RequestBody) {
	requestBase.body = requestBody
}
func (requestBase *RequestBase) setRequestBaseTail(requestTail RequestTail) {
	requestBase.tail = requestTail
}
