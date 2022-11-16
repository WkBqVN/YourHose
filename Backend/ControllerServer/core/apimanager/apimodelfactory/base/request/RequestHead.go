package base

type RequestHead struct {
	RequestFrom      string
	RequestSourceIP  string
	RequestInputType string
}
type IRequestHead interface {
	getRequestHeadSourceIP() string
	getRequestHeadInputType() string
	getRequestHeadFrom() string
	setRequestHeadIP(ip string)
	setRequestHeadInputType(typeRequest string)
	setRequestHeadFrom(from string)
}

func (requestHead *RequestHead) getRequestHeadSourceIP() string {
	return requestHead.RequestSourceIP
}

func (requestHead *RequestHead) getRequestHeadInputType() string {
	return requestHead.RequestInputType
}
func (requestHead *RequestHead) getRequestHeadFrom() string {
	return requestHead.RequestSourceIP
}
func (requestHead *RequestHead) setRequestHeadIP(ip string) {
	requestHead.RequestSourceIP = ip
}
func (requestHead *RequestHead) setRequestHeadInputType(typeRequest string) {
	requestHead.RequestInputType = typeRequest
}

func (requestHead *RequestHead) setRequestHeadFrom(from string) {
	requestHead.RequestFrom = from
}
