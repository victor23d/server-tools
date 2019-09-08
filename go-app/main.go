package main

import (
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"os"
	"strconv"
)

var (
	logger, _ = zap.NewProduction()
	log       = logger.Sugar()
)

func main() {
	http.HandleFunc("/", echoserver)
	port := 8080
	log.Infof("listen: %s", strconv.Itoa(port))
	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		log.Fatal(err)
	}
}

func echoserver(w http.ResponseWriter, r *http.Request) {
	err := GetInfo(w, r)
	if err != nil {
		log.Fatal(err)
	}
	err = GetIP(w, r)
	if err != nil {
		log.Fatal(err)
	}
}

func GetInfo(w http.ResponseWriter, r *http.Request) error {

	host, _ := os.Hostname()
	log.Info(host)
	log.Infof("r.URL.Path: %s\n", r.URL.Path)
	_, err := w.Write([]byte("hostname: " + host + "\n"))
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(w, "r.URL.Path: %s\n", r.URL.Path)
	if err = r.ParseForm(); err != nil {
		return err
	}
	log.Infof("Form %s\n", r.Form)
	fmt.Fprintf(w, "Form %s\n", r.Form)

	return err
}

func GetIP(w http.ResponseWriter, r *http.Request) error {

	XRealIP := r.Header.Get("X-Real-Ip")
	_, err := fmt.Fprintf(w, "X-Real-Ip: %s\n", XRealIP)
	log.Infof("X-Real-Ip: %s\n", XRealIP)

	XForwardFor := r.Header.Get("X-Forwarded-For")
	_, err = fmt.Fprintf(w, "X-Forwarded-For: %s\n", XForwardFor)
	log.Infof("X-Forwarded-For: %s\n", XForwardFor)

	RemoteAddr := r.RemoteAddr
	_, err = fmt.Fprintf(w, "RemoteAddr: %s\n", RemoteAddr)
	log.Infof("RemoteAddr: %s\n", RemoteAddr)
	return err
}

func init() {
	// 面对疾风吧！
}
