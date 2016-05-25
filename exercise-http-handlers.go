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
/* the declaration of func ListenAndServe
func ListenAndServe(addr string, handler Handler) error
*/


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
	// your http.Handle calls here
	http.Handle("/string", String("I'm a frayed knot."))
	//http://localhost:4000/sting                          I'm a frayed knot.  
	http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
	//http://localhost:4000/struct                         Hello:Gophers!  
	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}
