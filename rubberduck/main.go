package main

import (
	"log"
	"os"
	"fmt"
	"net/http"
)

var (
	version = "1.0.1"
	sysLog  *log.Logger
	errLog  *log.Logger
)

func initLogger() {
	sysLog = log.New(os.Stdout, "SYS: ", log.Ldate|log.Ltime|log.Llongfile)
	errLog = log.New(os.Stderr, "ERR: ", log.Ldate|log.Ltime|log.Llongfile)
}

func main() {
	initLogger()
	sysLog.Printf("Booting program...")
	// Root Endpoint
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		host,err := os.Hostname()
		if err!=nil {
			errLog.Fatalf("Oops something went wrong while retrieving hostname")
		}
		fmt.Fprintf(w, fmt.Sprintf("Hello folks! I am running on node: %s", host))
	})
	// Healthcheck Endpoint
	http.HandleFunc("/ping", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, fmt.Sprintf("PONG! app version: %s", version))
	})
	// Start a web server
	sysLog.Printf("Listening on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		errLog.Fatalf("Error starting server: %s", err)
	}
}
