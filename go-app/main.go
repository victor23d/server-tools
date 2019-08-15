package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	host, _ := os.Hostname()
	log.Println(host)
	fmt.Fprintf(w, host)

	fmt.Fprintf(w, "%s", "\n")
	log.Printf("r.URL.Path: %s\n", r.URL.Path[1:])
	fmt.Fprintf(w, "r.URL.Path: %s\n", r.URL.Path[1:])
	// fmt.Fprintf(w, "%s", "\n")
	GetIP(w, r)
}

func GetIP(w http.ResponseWriter, r *http.Request) {

	XRealIP := r.Header.Get("X-Real-Ip")
	fmt.Fprintf(w, "X-Real-Ip: "+XRealIP+"\n")

	XForwardFor := r.Header.Get("X-Forwarded-For")
	fmt.Fprintf(w, "X-Forwarded-For: "+XForwardFor+"\n")

	RemoteAddr := r.RemoteAddr
	fmt.Fprintf(w, "RemoteAddr: "+RemoteAddr+"\n")
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Server start at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// 面对疾风吧！
