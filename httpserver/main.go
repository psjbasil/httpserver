package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	_ "os"

	_ "net/http/pprof"

	"github.com/golang/glog"
)

func main() {
	glog.V(2).Info("Starting http server...")
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/healthz", healthz)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}

}

func healthz(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "ok\n")
}

func copyHeader(w http.ResponseWriter, r *http.Request) {
	for name, headers := range r.Header {
		for _, h := range headers {
			w.Header().Set(name, h)
		}
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("entering root handler")
	copyHeader(w, r)
	user := r.URL.Query().Get("user")
	if user != "" {
		io.WriteString(w, fmt.Sprintf("hello [%s]\n", user))
	} else {
		io.WriteString(w, "hello [stranger]\n")
	}
	io.WriteString(w, "===================Details of the http request header:============\n")
	r.Header.Set("VERSION", os.Getenv("VERSION"))
	for k, v := range r.Header {
		io.WriteString(w, fmt.Sprintf("%s=%s\n", k, v))
	}
	fmt.Println("client ip is:", r.RemoteAddr)
	fmt.Println("Status code is:", w.Header().Get("Status-Code"))
}
