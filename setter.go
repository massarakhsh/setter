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
	http.HandleFunc("/create/", router_create)
	http.HandleFunc("/get/", router_get)
	http.HandleFunc("/update/", router_update)
	http.HandleFunc("/delete/", router_delete)
	http.HandleFunc("/count", router_count)
	if err := http.ListenAndServe(Address, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func router_create(w http.ResponseWriter, r *http.Request) {
	_,key,val := marshalIn(r)
	DataCreate(key, val)
	marshalOut(w, "")
}

func router_get(w http.ResponseWriter, r *http.Request) {
	_,key,_ := marshalIn(r)
	val := DataGet(key)
	marshalOut(w, val)
}

func router_update(w http.ResponseWriter, r *http.Request) {
	_,key,val := marshalIn(r)
	DataUpdate(key, val)
	marshalOut(w, "")
}

func router_delete(w http.ResponseWriter, r *http.Request) {
	_,key,_ := marshalIn(r)
	DataDelete(key)
	marshalOut(w, "")
}

func router_count(w http.ResponseWriter, r *http.Request) {
	marshalIn(r)
	val := fmt.Sprint(DataCount())
	marshalOut(w, val)
}

func marshalIn(r *http.Request) (string,string,string) {
	fun,key,val := "","",""
	if parts := strings.Split(r.RequestURI,"/"); len(parts) > 1 {
		fun = strings.ToLower(parts[1])
		parts = parts[2:]
		if len(parts) > 0 {
			key = parts[0]
			parts = parts[1:]
		}
		if len(parts) > 0 {
			val = parts[0]
			parts = parts[1:]
		}
	}
	return fun,key,val
}

func marshalOut(w http.ResponseWriter, answer string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Cache-control", "private, max-age=0, no-cache")
	w.WriteHeader(200)
	fmt.Fprint(w, answer)
}

