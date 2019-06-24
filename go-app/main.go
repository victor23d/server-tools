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

	log.Printf("%s", r.URL.Path[1:])
	fmt.Fprintf(w, "%s", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// 面对疾风吧！
