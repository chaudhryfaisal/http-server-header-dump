package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"strconv"
	"time"
)

func handler(w http.ResponseWriter, req *http.Request) {
	requestDump, err := httputil.DumpRequest(req, true)
	if err != nil {
		fmt.Println("Errror:", err)
	}
	fmt.Println(string(requestDump))
	fmt.Printf("------------------------- %s -------------------------\n", time.Now().Format("2006-01-02 3:4:5 PM"))
	req.Write(w) // Save a copy of this request for debugging.
}

func main() {
	port := flag.Int("port", 8081, "Listen on port")
	flag.Parse()

	http.HandleFunc("/", handler)
	log.Println("Starting http server on port:", *port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(*port), nil))
}
