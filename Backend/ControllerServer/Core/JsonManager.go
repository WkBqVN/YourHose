package core

type jsonManager interface {
	genMapObj(key string, obj map[string]string) map[string]interface{}
	genStringObj(key string, value string) map[string]string
	genIntObj(key string, value int) map[string]int
}
