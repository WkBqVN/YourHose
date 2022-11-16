package base

type RequestBody struct {
	RequestPayload string
}
type IRequestbody interface {
	getRequestPayload() string
	setRequestPayload(payload string)
}

func (requestBody *RequestBody) getRequestPayload() string {
	return requestBody.RequestPayload
}
func (requestBody *RequestBody) setRequestPayload(payload string) {
	requestBody.RequestPayload = payload
}
