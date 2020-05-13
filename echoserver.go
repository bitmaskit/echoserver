package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"
)

func main() {
	port := os.Args[1]
	fmt.Printf("Starting server on port: %s\n", port)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Save a copy of request for debug
		requestDump, err := httputil.DumpRequest(r, true)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(requestDump))
		fmt.Fprintf(w, string(requestDump))
	})
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
}
