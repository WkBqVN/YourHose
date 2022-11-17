package request

type IRequest interface {
	GetRequestBase() interface{}
	GetRequestBaseHead() interface{}
	GetRequestBaseBody() interface{}
	GetRequestBaseTail() interface{}
	SetRequestBaseHead(requestHead interface{})
	SetRequestBaseBody(requestBody interface{})
	SetRequestBaseTail(RequestTail interface{})
	SetRequestBase(request *interface{})
	GetRequestHead() interface{}
	GetRequestHeadSourceIP() string
	GetRequestHeadInputType() string
	GetRequestHeadFrom() string
	GetRequestMethod() string
	SetRequestHeadIP(ip string)
	SetRequestHeadInputType(typeRequest string)
	SetRequestHeadFrom(from string)
	SetRequestMethod(method string)
	SetRequestHead(request interface{})
	GetRequestPayload() string
	SetRequestPayload(payload string)
	GetRequestTailTo() string
	GetRequestTailDesIP() string
	GetRequestTailOutpputType() string
	SetTailTo(to string)
	SetRequestTailDesIP(ip string)
	SetRequestTailOutputType(typeRequest string)
	ToString() string
}
