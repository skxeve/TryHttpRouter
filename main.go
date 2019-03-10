package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Hello HttpRouter Index")
}

func Number(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Fprintf(w, "Number parameter is %s", params.ByName("id"))
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/number/:id", Number)

	log.Fatal(http.ListenAndServe(":8080", router))
}
