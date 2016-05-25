package main

import (
	"fmt"
	"log"
	"net/http"
)

type String string

type Struct struct {
	Greeting string
	Punct    string
	Who      string
}

/*definition of Handler ineterface
type Handler interface{
	serveHTTP(ResposeWriter, *Request)//router
}*/

func (h String) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, h)
}
func (h Struct) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, h.Greeting+h.Punct+h.Who)
}

func main() {
	a := Struct{"Hello", ":", "Gophers!"}
	http.Handle("/string", String("I'm a frayed knot."))
	http.Handle("/struct", &a)
	// your http.Handle calls here
	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}
