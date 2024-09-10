package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

var total = 0

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s port\nExample: %s 4444\n\n", os.Args[0], os.Args[0])
		os.Exit(1)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		total++
		fmt.Printf("###--- Request #%06d | from %s ---###\n", total, r.RemoteAddr)
		fmt.Printf("%s %s %s\n", r.Method, r.RequestURI, r.Proto)
		for k, v := range r.Header {
			fmt.Printf("%s: %s\n", k, strings.Join(v, ";"))
		}
		if r.ContentLength > 0 {
			fmt.Print("\n")
			buf := new(bytes.Buffer)
			if _, err := io.Copy(buf, r.Body); err != nil {
				fmt.Printf("error reading body")
			}
			fmt.Printf("%s", buf.Bytes())
		}
		fmt.Print("\n###---###\n\n")
		w.WriteHeader(204)
	})
	if len(os.Args) == 5 && os.Args[2] == "tls" {
		fmt.Printf("Beep! :>\nListening on port %s (with tls: cert %s / key %s)\n", os.Args[1], os.Args[3], os.Args[4])
		if err := http.ListenAndServeTLS(fmt.Sprintf(":%s", os.Args[1]), os.Args[3], os.Args[4], nil); err != nil {
			fmt.Printf("Boop! :<\nGot an error trying to listen on port %s (with tls: cert %s / key %s): \"%s\"\n", os.Args[1], os.Args[3], os.Args[4], err.Error())
			os.Exit(1)
		}
	} else {
		fmt.Printf("Beep! :>\nListening on port %s\n", os.Args[1])
		if err := http.ListenAndServe(fmt.Sprintf(":%s", os.Args[1]), nil); err != nil {
			fmt.Printf("Boop! :<\nGot an error trying to listen on port %s: \"%s\"\n", os.Args[1], err.Error())
			os.Exit(1)
		}
	}
}
