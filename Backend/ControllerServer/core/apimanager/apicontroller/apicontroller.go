package apicontroller

import (
	"controllerserver/core/apimanager/apimodel/base/response"
	"controllerserver/core/apimanager/apimodel/receiver"
	"controllerserver/core/apimanager/apimodel/sender"
	"log"
	"net/http"
)

type Controller struct {
	receiver   receiver.RequestReceiver
	controller int
	sender     sender.ResponseSender
	dataMap    map[string]string
}

func GetNewController() Controller {
	return Controller{}
}
//	func readRequest(w http.ResponseWriter, r *http.Request) {
//		requestIn := request.GetNewRequestBase()
//
// x.GetRequestBase()
//
//		io.WriteString(w, s)
//	}
//
// http.HandleFunc("/", readRequest)

func (controller *Controller) StartServerUp(port string) {
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

// receiver head
// requestFrom      string
// requestSourceIP  string
// requestInputType string
// requestMethod    string

// receiver tail
// requestTo         string
// requestDesIP      string
// requestOutputType string

// response head
// responseTo      string
// responseDesIP   string
// responseOutType string
// responseMethod  string

// resposne tail
// responseFrom       string
// responseSouceIP    string
// responseOutputType string
// tail should be setting up to db and take data from head to mapping
func (controller *Controller) runProcess(receiverParam receiver.RequestReceiver) {
	senderObj := sender.GetNewResponseSender()
	// receiver split element
	receiverHead := receiverParam.GetRequestObj().GetRequestBaseHead()
	receiverTail := receiverParam.GetRequestObj().GetRequestBaseTail()
	receiverBody := receiverParam.GetRequestObj().GetRequestBaseBody()
	// create element store data
	head := response.GetNewResponseHead()
	body := response.GetNewResponseBody()
	tail := response.GetNewResponseTail()
	// assgin value to element
	head.SetResponseHeadTo(receiverTail.GetRequestTailTo())
	head.SetResponseHeadDesIP(receiverTail.GetRequestTailDesIP())
	body.SetResponsePayload(receiverBody.GetRequestPayload())
	tail.SetResponseTailFrom(receiverHead.GetRequestHeadFrom())
	tail.SetResponseTailSourceIP("ip")
	// assign element to obj
	senderObj.GetRequestObj().SetResponseBaseHead(head)
	senderObj.GetRequestObj().SetResponseBaseBody(body)
	senderObj.GetRequestObj().SetResponseBaseTail(tail)
}

