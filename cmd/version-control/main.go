package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"
	"time"

	"storj.io/storj/internal/version"
)

var (
	ver version.Version_Info
)

func handleGet(w http.ResponseWriter, r *http.Request) {
	var xfor string

	// Only handle GET Requests
	if r.Method == "GET" {
		if xfor = r.Header.Get("X-Forwarded-For"); xfor == "" {
			xfor = r.RemoteAddr
		}
		log.Printf("Request from: %s for %s", r.RemoteAddr, xfor)

		w.Header().Set("Content-Type", "application/json")
		response, err := ver.Marshal()
		if err != nil {
			w.WriteHeader(500)
		}

		w.WriteHeader(200)
		_, err = w.Write(response)
		if err != nil {
			log.Printf("error writing response to client: %v", err)
		}
	}
}

func main() {
	// Flags to specify required Version
	addr := flag.String("listen", "0.0.0.0:8080", "Defines Listen Address of Webserver")
	stimestamp := flag.String("timestamp", strconv.FormatInt(time.Now().UnixNano(), 10), "Sets Timestamp of Build")
	scommit := flag.String("commit", "", "Sets CommitHash of Build")
	sversion := flag.String("version", "v0.1.0", "Sets required Version Number")
	srelease := flag.Bool("release", false, "Sets if version is an official release")
	flag.Parse()

	if flag.Parsed() {
		ver = version.Version_Info{
			Timestamp:  *stimestamp,
			CommitHash: *scommit,
			Version:    *sversion,
			Release:    *srelease,
		}
	}

	log.Printf("setting version info to: %v", ver)
	http.HandleFunc("/", handleGet)
	log.Println("starting Webserver")

	// Not pretty but works..
	log.Fatal(http.ListenAndServe(*addr, nil))
}
