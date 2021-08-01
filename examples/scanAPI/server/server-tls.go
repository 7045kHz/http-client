package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", tls_handler)
	http.ListenAndServeTLS(":10443", "cert.pem", "key.pem", nil)
}

func tls_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Request at %v\n", time.Now())
	fmt.Fprintf(w, "URL: %v\n", r.URL)
	fmt.Fprintf(w, "Request at %v\n", time.Now())
	fmt.Fprintf(w, "Method: %v\n", r.Method)
	fmt.Fprintf(w, "Content Length: %v\n", r.ContentLength)
	data, _ := ioutil.ReadAll(r.Body)
	fmt.Fprintf(w, "Body: %v\n", string(data))
	//fmt.Fprintf(w, "TLS: %v\n", r.TLS)
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

// Go to https://localhost:10443/ or https://127.0.0.1:10443/
// list of TCP ports:
// https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers

// Generate unsigned certificate
// go run $(go env GOROOT)/src/crypto/tls/generate_cert.go --host=somedomainname.com
// for example
// go run $(go env GOROOT)/src/crypto/tls/generate_cert.go --host=localhost

// WINDOWS
// windows may have issues with go env GOROOT
// go run %(go env GOROOT)%/src/crypto/tls/generate_cert.go --host=localhost

// instead of go env GOROOT
// you can just use the path to the GO SDK
// wherever it is on your computer
