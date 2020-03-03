package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

var (
	Address = ":80"
)

func main() {
	DataInit()
	http.HandleFunc("/create/", router_api)
	http.HandleFunc("/get/", router_api)
	http.HandleFunc("/update/", router_api)
	http.HandleFunc("/delete/", router_api)
	http.HandleFunc("/count", router_api)
	if err := http.ListenAndServe(Address, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func router_api(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Cache-control", "private, max-age=0, no-cache")
	w.WriteHeader(200)
	if parts := strings.Split(r.RequestURI,"/"); len(parts) > 1 {
		fun := strings.ToLower(parts[1])
		parts = parts[2:]
		key, val := "", ""
		if len(parts) > 0 {
			key = parts[0]
			parts = parts[1:]
		}
		if len(parts) > 0 {
			val = parts[0]
			parts = parts[1:]
		}
		switch fun {
		case "create":
			DataCreate(key, val)
		case "get":
			fmt.Fprint(w, DataGet(key))
		case "update":
			DataUpdate(key, val)
		case "delete":
			DataDelete(key)
		case "count":
			fmt.Fprint(w, DataCount())
		}
	}
}

