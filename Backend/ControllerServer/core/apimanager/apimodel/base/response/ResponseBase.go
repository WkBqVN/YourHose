package response

type ResponseBase struct {
	head responseHead
	body responseBody
	tail responseTail
}
