package main

import (
	"fmt"
	"github.com/keshavrkaranth/mongoDb/router"
	"log"
	"net/http"
)

func main() {
	fmt.Println("MONGO DB API")
	r := router.Router()
	fmt.Println("Server listening Started")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listning at port 4000")
}
