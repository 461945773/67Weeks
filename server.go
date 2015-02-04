// server.go

package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Server go at port:1314")
	fileServer := http.FileServer(http.Dir("./"))
	err := http.ListenAndServe(":1314", fileServer)
	if err != nil {
		fmt.Println(err)
	}
}
