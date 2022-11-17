// not use so much. becaues it reduce flexible
package data

type IManager interface {
	IJsonManager
	IGraphManager
}

type IJsonManager interface {
	GetJsonManager() jsonManager
	GetJsonObj() map[string]interface{}
	GetJsonData() string
	SetJsonObj(jsomMap map[string]interface{})
	SetJsonData(data string)
	SetJsonManager(jsm *jsonManager)
}

type IGraphManager interface {
}
