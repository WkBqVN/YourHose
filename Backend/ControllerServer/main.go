package main

import (
	"controllerserver/core/apimanager/apimodel/base/request"
	"controllerserver/core/data"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", readRequest)
	err := http.ListenAndServe(":5050", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func readRequest(w http.ResponseWriter, r *http.Request) {
	requestIn := request.GetNewRequestBase()
	requestHead := requestIn.GetRequestBaseHead()
	requestBody := requestIn.GetRequestBaseBody()
	requestTail := requestIn.GetRequestBaseTail()

	requestHead.SetRequestHeadFrom(r.Host)
	requestHead.SetRequestHeadInputType(r.Method)
	requestHead.SetRequestHeadIP("ip")
	requestHead.SetRequestHeadMethod(r.Method)

	requestIn.SetRequestBaseHead(requestHead)
	requestIn.SetRequestBaseTail(requestTail)
	requestIn.SetRequestBaseBody(requestBody)

	// io.WriteString(w, requestIn.ToString())
	testMap := map[string]interface{}{
		"a": "bc",
	}
	jsstr := `{"a":"b"}`
	jsm := data.GetJsonManager()

	s := jsm.ConvertObjToJson(testMap)

	jsm.ConvertJsonToObj(jsstr)

	fmt.Println("dfas: ")
	fmt.Println(jsm.GetJsonObj())

	// var x request.IRequest
	// x.GetRequestBase()

	io.WriteString(w, s)
}
