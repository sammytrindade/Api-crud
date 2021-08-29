package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, response *http.Request) {
		fmt.Fprintf(writer, "Hello Peter")
	})

	http.ListenAndServe(":8080", nil)

}
