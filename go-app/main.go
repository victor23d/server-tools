package main

import (
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"os"
	"strconv"
	"time"
	"flag"
)

var (
	logger, _ = zap.NewProduction()
	log       = logger.Sugar()
)

var (
	qa bool
	count = 1
)


func main() {
 	flag.BoolVar(&qa, "qa", false, "enable qa mode")
 	flag.Parse()
 	serve()

}

func serve(){
	http.HandleFunc("/", echoserver)
	port := 8080
	log.Infof("listen: %s", strconv.Itoa(port))
	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		log.Fatal(err)
	}
}

func echoserver(w http.ResponseWriter, r *http.Request) {
	if count == 11 { count = 1 }
	//if count > 10 { count = 1 }
	if qa == true{
		if count == 4 {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			count ++
			return
		}else if count ==5 {
			http.Error(w,http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
			count ++
			return
		}
	}
	err := GetInfo(w, r)
	if err != nil {
		log.Fatal(err)
	}
	err = GetIP(w, r)
	if err != nil {
		log.Fatal(err)
	}
	now := time.Now()
	fmt.Fprintf(w, "Now: %s\n", now.Format(time.RFC3339))
	fmt.Fprintf(w, "count: %d\n",count)
	log.Infof("count: %d", count)
	count ++
	log.Info("========================================")
}

func GetInfo(w http.ResponseWriter, r *http.Request) error {

	host, _ := os.Hostname()
	log.Info(host)
	_, err := w.Write([]byte("hostname: " + host + "\n"))
	if err != nil {
		return err
	}
	log.Infof("r.URL.Path: %s", r.URL.Path)
	_, err = fmt.Fprintf(w, "r.URL.Path: %s\n", r.URL.Path)
	if err = r.ParseForm(); err != nil {
		return err
	}
	log.Infof("Form %s", r.Form)
	fmt.Fprintf(w, "Form %s\n", r.Form)

	return err
}

func GetIP(w http.ResponseWriter, r *http.Request) error {

	XRealIP := r.Header.Get("X-Real-Ip")
	log.Infof("X-Real-Ip: %s", XRealIP)
	_, err := fmt.Fprintf(w, "X-Real-Ip: %s\n", XRealIP)

	XForwardFor := r.Header.Get("X-Forwarded-For")
	log.Infof("X-Forwarded-For: %s", XForwardFor)
	_, err = fmt.Fprintf(w, "X-Forwarded-For: %s\n", XForwardFor)

	RemoteAddr := r.RemoteAddr
	log.Infof("RemoteAddr: %s", RemoteAddr)
	_, err = fmt.Fprintf(w, "RemoteAddr: %s\n", RemoteAddr)

	UserAgent := r.Header.Get("User-Agent")
	log.Infof("User-Agent: %s", UserAgent)
	_, err = fmt.Fprintf(w, "User-Agent: %s\n", UserAgent)

	return err
}
