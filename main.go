package Api_crud

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		
	})
	http.ListenAndServe(":8080", nil)

}
