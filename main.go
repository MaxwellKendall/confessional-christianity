package main

import (
	"fmt"
	"net/http"
)

func handler(rw http.ResponseWriter, req *http.Request) {
	bs := []byte("YO")
	fmt.Println("hi", req.Method)
	fmt.Println("hi", req.Method)
	rw.Write(bs)
}

func main() {
	http.HandleFunc("/test", handler)
	http.ListenAndServe(":1517", nil)
}
