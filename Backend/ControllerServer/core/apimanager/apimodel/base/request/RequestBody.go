package request

type requestBody struct {
	requestPayload string
}

func GetNewRequestBody() requestBody {
	return requestBody{}
}
func (requestBody *requestBody) GetRequestPayload() string {
	return requestBody.requestPayload
}
func (requestBody *requestBody) SetRequestPayload(payload string) {
	requestBody.requestPayload = payload
}

func (requestBody *requestBody) SetRequestBody(requestBodyNew requestBody) {
	*requestBody = requestBodyNew
}

func (requestBody *requestBody) ToString() string {
	returnString := ""
	returnString += "PayLoad : " + requestBody.GetRequestPayload() + "\n"
	return returnString
}
