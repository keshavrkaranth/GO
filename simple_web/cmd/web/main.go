package main

import (
	"fmt"
	"github.com/keshavrkaranth/simple_web/pkg/handlers"
	"net/http"
)

var portNumber = ":8000"

func main() {
	http.HandleFunc("/", handlers.HomePage)
	http.HandleFunc("/about", handlers.AboutPage)
	fmt.Println(fmt.Sprintf("Started listning server at port%s\n", portNumber))
	http.ListenAndServe(portNumber, nil)
}
