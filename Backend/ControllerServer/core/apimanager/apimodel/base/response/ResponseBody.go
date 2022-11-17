package response

type responseBody struct {
	responsePayload string
}
func GetNewResponseBody() responseBody{
	return responseBody{}
}
func (responseBody *responseBody) GetResponsePayload() string {
	return responseBody.responsePayload
}
func (responseBody *responseBody) SetResponsePayload(payload string) {
	responseBody.responsePayload = payload
}
func (responseBody *responseBody) ToString() string{
	returnString := ""
	returnString += "Payload: " + responseBody.responsePayload + "\n"
	return returnString
}
