// From https://gist.github.com/d-schmidt/587ceec34ce1334a5e60
package main

import (
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var host = os.Getenv("HOST")
var keepPort = os.Getenv("KEEP_PORT")

func redirect(w http.ResponseWriter, r *http.Request) {
	toHost := host
	if toHost == "" {
		toHost = r.Host
	}
	if keepPort == "" {
		toHost = strings.Split(toHost, ":")[0]
	}
	targetUrl := url.URL{
		Scheme:   "https",
		Host:     toHost,
		Path:     r.URL.Path,
		RawQuery: r.URL.RawQuery,
	}
	log.Printf("redirect to: %s", targetUrl)
	http.Redirect(w, r, targetUrl.String(), http.StatusTemporaryRedirect)
}

func main() {
	// redirect every http request to https
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	log.Println("listening on", port)
	http.ListenAndServe(":"+port, http.HandlerFunc(redirect))
}
