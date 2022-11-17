package data

import (
	"encoding/json"
	"fmt"
)

type jsonManager struct {
	jsonObj map[string]interface{}
	data    string
}

func GetJsonManager() jsonManager {
	return jsonManager{}
}

func (jsm *jsonManager) SetJsonManager(jsmNew jsonManager) {
	jsm = &jsmNew
}

func (jsm *jsonManager) GetJsonObj() map[string]interface{} {
	return jsm.jsonObj
}
func (jsm *jsonManager) SetJsonObj(jsomMap map[string]interface{}) {
	jsm.jsonObj = jsomMap
}
func (jsm *jsonManager) GetJsonData() string {
	return jsm.data
}

func (jsm *jsonManager) SetJsonData(data string) {
	jsm.data = data
}

func (jsm *jsonManager) ConvertJsonToObj(jsonString string) map[string]interface{} {
	jsm.data = jsonString
	err := json.Unmarshal([]byte(jsonString), &jsm.jsonObj)
	if err != nil {
		return jsm.jsonObj
	} else {
		return nil
	}
}

func (jsm *jsonManager) ConvertObjToJson(jsonObj map[string]interface{}) string {
	jsonStr, err := json.Marshal(jsonObj)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		jsm.data = string(jsonStr)
	}
	return jsm.data
}
