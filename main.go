package main

import (
	// input output
	"fmt"
	// komunikasi http client and server
	"net/http"
	// module yang sudah dibuat
	"datahandle"
)

func main() {

	// register alamat local url yang akan dihandle
	http.HandleFunc("/", datahandle.HomeUrl)
	http.HandleFunc("/student", datahandle.HandleUrl)

	// print output
    fmt.Println("Listening on 127.0.0.1:8080...")
	
	// menjalankan http global untuk listening pada request port :8080
    http.ListenAndServe(":8080", nil)
}