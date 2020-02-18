package main

import "net/http"
import "net/http/httputil"
import "net/url"
import "log"
import "os"

func main() {
	if len(os.Args) < 3 {
		log.Fatalf("Usage %s https://path/to/proxy/to :port", os.Args[0])
	}

	target, err := url.Parse(os.Args[1])
	if err != nil {
		log.Fatalf("Error: Could Not Parse URL")
	}

	reverseProxy := httputil.NewSingleHostReverseProxy(target)
	handler := func(w http.ResponseWriter, r *http.Request) {
		r.Host = target.Host
		reverseProxy.ServeHTTP(w, r)
	}

	log.Fatal(http.ListenAndServe(os.Args[2], http.HandlerFunc(handler)))
}
