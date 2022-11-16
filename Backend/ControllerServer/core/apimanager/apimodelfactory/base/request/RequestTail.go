package base

type RequestTail struct {
	RequestTo         string
	RequestDesIP      string
	RequestOutputType string
}
type IRequestTail interface {
	getRequestTailTo() string
	getRequestTailDesIP() string
	getRequestTailOutpputType() string
	setTailTo(to string)
	setRequestTailDesIP(ip string)
	setRequestTailOutputType(typeRequest string)
}

func (requestTail *RequestTail) getRequestTailTo() string {
	return requestTail.RequestTo
}

func (requestTail *RequestTail) getRequestTailDesIP() string {
	return requestTail.RequestDesIP
}
func (requestTail *RequestTail) getRequestTailOutpputType() string {
	return requestTail.RequestOutputType
}

func (requestTail *RequestTail) setTailTo(to string) {
	requestTail.RequestTo = to
}
func (requestTail *RequestTail) setRequestTailDesIP(ip string) {
	requestTail.RequestDesIP = ip
}
func (requestTail *RequestTail) setRequestTailOutputType(typeRequest string) {
	requestTail.RequestOutputType = typeRequest
}
