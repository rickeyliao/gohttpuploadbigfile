package main

import (
	"net/http"
	"github.com/rickeyliao/gohttpuploadbigfile/file"
	"log"
)

func main()  {
	mux := http.NewServeMux()
	listenport:=":50810"

	mux.Handle("/upload",file.NewFileUpLoad())
	httpserver := &http.Server{Addr: listenport, Handler: mux}

	log.Fatal(httpserver.ListenAndServe())
}
