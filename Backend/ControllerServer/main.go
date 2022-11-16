package main

import (
	"controllerserver/sender"
	// "encoding/json"
	"fmt"
	// "log"
	// "net/http"
)

// "controllerserver/Receiver"

// "controllerserver/Core"

func main() {
	p := sender.RequestSender{}

	fmt.Println(p)

	// p.RequestBody["test"] = "test contect"
	// http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.WriteHeader(http.StatusCreated)
	// 	jsonResp, err := json.Marshal(p)
	// 	if err != nil {
	// 		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	// 	}
	// 	w.Write(jsonResp)
	// 	return
	// })

	// fmt.Printf("Starting server at port 5050\n")
	// if err := http.ListenAndServe(":5050", nil); err != nil {
	// 	log.Fatal(err)
	// }
}
