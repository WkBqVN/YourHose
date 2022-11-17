package response

type IResponseBase interface {
	GetNewResponseBase() interface{}
	GetResponseBaseHead() interface{}
	GetResponseBaseBody() interface{}
	GetResponseBaseTail() interface{}
	SetResponseBaseHead(requestHead interface{})
	SetResponseBaseBody(requestBody interface{})
	SetResponseBaseTail(RequestTail interface{})
	SetResponseBase(request *interface{})
	IResponseToData
}

type IResponseHead interface {
	GetNewResponseHead() interface{}
	GetResponseHeadDesIP() string
	GetResonseHeadInputType() string
	GetResponseHeadTo() string
	GetResponseMethod() string
	SetResponseHeadDesIP(ip string)
	SetResponseHeadOutputType(typeRequest string)
	SetResponseHeadTo(from string)
	SetResponseMethod(method string)
	SetResponseHead(request interface{})
}

type IResponseBody interface {
	GetNewResponseBody() interface{}
	GetResponsePayload() string
	SetResponsePayload(payload string)
}

type IResponseTail interface {
	GetNewResponseTail() interface{}
	GetResponseTailFrom() string
	GetResponseTailSourceIP() string
	GetResponseTailOutpputType() string
	SetResponseTailFrom(from string)
	SetResponseTailDesIP(ip string)
	SetResponseTailOutputType(typeRequest string)
	SetResponseTail(response interface{})
}

type IResponseToData interface {
	HeadToString() string
	BodyToString() string
	TailToString() string
}
