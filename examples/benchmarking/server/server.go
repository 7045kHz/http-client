package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	data, _ := ioutil.ReadAll(r.Body)

	fmt.Printf("Request at %v\n", time.Now())
	fmt.Fprintf(w, "URL: %v\n", r.URL)
	fmt.Fprintf(w, "Request at %v\n", time.Now())
	fmt.Fprintf(w, "Method: %v\n", r.Method)
	fmt.Fprintf(w, "Content Length: %v\n", r.ContentLength)

	fmt.Fprintf(w, "Body: %v\n", string(data))
	fmt.Fprintf(w, "TLS Version: %v\n", r.TLS.Version)
	fmt.Fprintf(w, "TLS Protocol: %v\n", r.TLS.NegotiatedProtocol)
	fmt.Fprintf(w, "Host: %v\n", r.Host)
	fmt.Fprintf(w, "URL User: %v\n", r.URL.User)
	fmt.Fprintf(w, "Remove Address: %v\n", r.RemoteAddr)
	fmt.Fprintf(w, "Request URI: %v\n", r.RequestURI)
	fmt.Fprintf(w, "Response: %v\n", r.Response)
	for k, v := range r.Header {
		fmt.Printf("%v: %v\n", k, v)
		fmt.Fprintf(w, "%v:%v\n", k, v)
	}
}

func main() {
	fmt.Printf("Starting Server at %v\n", time.Now())
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9090", nil)
}
