package request

type requestBody struct {
	requestPayload string
}

func (requestBody *requestBody) GetRequestPayload() string {
	return requestBody.requestPayload
}
func (requestBody *requestBody) SetRequestPayload(payload string) {
	requestBody.requestPayload = payload
}

func (requestBody *requestBody) ToString() string {
	returnString := ""
	returnString += "PayLoad : " + requestBody.GetRequestPayload() + "\n"
	return returnString
}
