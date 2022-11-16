package core

type IMapConvert interface {
	convertMapObj(key string, value string) map[string]interface{}
	convertStringObj(key string, value string) map[string]string
	convertIntObj(key string, value int) map[string]int
}

type IRequestValidate interface {
	isNumber(value string) bool
	isChar(value string) bool
}

type IRequest interface{
	getTypeRequest() string
	getSourceIPRequest() string
	getDesIPRequest() string
	getBodyRequest() string
}