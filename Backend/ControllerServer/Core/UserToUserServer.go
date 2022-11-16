// package core

// import (
// 	// "controllerserver/receiver"
// 	"bytes"
// 	"controllerserver/sender"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// )

// type objModel struct {
// 	// input  receiver.RequestReceiver
// 	output sender.RequestSender
// }

// const serverPort = 5050

// var gObj objModel

// // func GetInputRequest() receiver.RequestReceiver {
// // 	return gObj.input
// // }
// // func GetOutputRequest() sender.RequestSender {
// // 	return gObj.output
// // }

// // func SetInputRequest(in receiver.RequestReceiver) {
// // 	gObj.input = in
// // }

// func HandleRequest() {
// 	t := http.Response{
// 		Body: ioutil.NopCloser(bytes.NewBufferString("Hello World")),
// 	}

// 	buff := bytes.NewBuffer(nil)
// 	t.Write(buff)

// 	fmt.Println(buff)
// }
package Core
func something() {

}