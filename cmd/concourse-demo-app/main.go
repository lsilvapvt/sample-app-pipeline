package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", IndexController)
	n := negroni.Classic()
	n.UseHandler(mux)
	n.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}

const (
	WelcomeMessage = "Welcome to my even more awesome page!"
)

func IndexController(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, WelcomeMessage)
}
