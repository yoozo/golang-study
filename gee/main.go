package main

import (
	"encoding/json"
	"fmt"
	"gee/engine"
	"log"
	"net/http"
)

var addr string = ":8080"

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatal("main panic ::: ", err)
		}
	}()

	engine := engine.New("gee frame")

	engine.AddRouter("GET", "/debug", func(w http.ResponseWriter, r *http.Request) {
		resp := make(map[string]any)

		fmt.Printf("path: %v, method: %v", r.Method, r.URL.Path)
		resp["code"] = 0
		resp["msg"] = "success"
		w.WriteHeader(http.StatusOK)
		if respJson, err := json.Marshal(resp); err != nil {
			log.Fatal("json.Marshal ::: ", err)
		} else {
			w.Write(respJson)
		}
	})
	// http.HandleFunc("/debug", func(w http.ResponseWriter, r *http.Request) {
	// 	resp := make(map[string]any)

	// 	fmt.Printf("path: %v, method: %v", r.Method, r.URL.Path)
	// 	resp["code"] = 0
	// 	resp["msg"] = "success"
	// 	w.WriteHeader(http.StatusOK)
	// 	if respJson, err := json.Marshal(resp); err != nil {
	// 		log.Fatal("json.Marshal ::: ", err)
	// 	} else {
	// 		w.Write(respJson)
	// 	}
	// })

	err := http.ListenAndServe(addr, engine)
	log.Fatal("ListenAndServe fail ::: ", err)
}
