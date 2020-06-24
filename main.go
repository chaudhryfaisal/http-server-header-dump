package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var dumpEnv = false
var envMatch = ".*"

func Handler(resp http.ResponseWriter, req *http.Request) {
	requestDump, err := httputil.DumpRequest(req, true)
	if err != nil {
		fmt.Println("Errror:", err)
	}
	fmt.Println(string(requestDump))
	fmt.Printf("------------------------- %s -------------------------\n", time.Now().Format("2006-01-02 3:4:5 PM"))
	if dumpEnv {
		for _, e := range os.Environ() {
			pair := strings.SplitN(e, "=", 2)
			res1, _ := regexp.MatchString(envMatch, pair[0])
			if res1 {
				req.Header.Add(fmt.Sprintf("ENV_%s", pair[0]), pair[1])
			}
		}
	}
	req.Write(resp)
}

func main() {
	port := flag.Int("port", 8080, "Listen on port")
	includeEnv := flag.Bool("includeEnv", dumpEnv, "Include Environment")
	matchEnv := flag.String("matchEnv", envMatch, "Match ENV KEYs")
	flag.Parse()
	dumpEnv = *includeEnv
	envMatch = *matchEnv
	log.Printf("Starting http server port=%d includeEnv=%t matchEnv=%s", *port, dumpEnv, envMatch)
	if dumpEnv {
		for _, e := range os.Environ() {
			pair := strings.SplitN(e, "=", 2)
			res1, _ := regexp.MatchString(envMatch, pair[0])
			if res1 {
				log.Printf("ENV: %s=%s\n", pair[0], pair[1])
			}
		}
	}
	http.HandleFunc("/", Handler)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(*port), nil))
}
