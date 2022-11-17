package request

type IRequestBase interface {
	GetNewRequestBase() interface{}
	GetRequestBaseHead() interface{}
	GetRequestBaseBody() interface{}
	GetRequestBaseTail() interface{}
	SetRequestBaseHead(requestHead interface{})
	SetRequestBaseBody(requestBody interface{})
	SetRequestBaseTail(RequestTail interface{})
	SetRequestBase(request *interface{})
	IRequestToData
}
type IRequestHead interface {
	GetNewRequestHead() interface{}
	GetRequestHeadSourceIP() string
	GetRequestHeadInputType() string
	GetRequestHeadFrom() string
	GetRequestMethod() string
	SetRequestHeadIP(ip string)
	SetRequestHeadInputType(typeRequest string)
	SetRequestHeadFrom(from string)
	SetRequestHeadMethod(method string)
	SetRequestHead(request interface{})
}
type IRequestBody interface {
	GetNewRequestBody() interface{}
	GetRequestPayload() string
	SetRequestPayload(payload string)
	SetRequestBody(requestBody interface{})
}
type IRequestTail interface {
	GetNewRequestTail() interface{}
	GetRequestTailTo() string
	GetRequestTailDesIP() string
	GetRequestTailOutpputType() string
	SetRequestTailTo(to string)
	SetRequestTailDesIP(ip string)
	SetRequestTailOutputType(typeRequest string)
}
type IRequestToData interface {
	HeadToString() string
	BodyToString() string
	TailToString() string
}
