package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"strconv"
)

func handler(w http.ResponseWriter, req *http.Request) {
	req.Write(w)// Save a copy of this request for debugging.
	requestDump, err := httputil.DumpRequest(req, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(requestDump))
	fmt.Fprintf(w, "RemoteAddr: %s", req.RemoteAddr)
	fmt.Println("---------------------------------------------------------------")
}

func main() {
	port := flag.Int("port", 8080, "Listen on port")
	flag.Parse()

	http.HandleFunc("/", handler)
	log.Println("Starting http server on port:", *port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(*port), nil))
}
