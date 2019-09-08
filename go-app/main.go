package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"path"
	"runtime"
	"strconv"
)

var (
	log = logrus.New()
)

func main() {
	http.HandleFunc("/", echoserver)
	port := 8080
	log.Println("listen: " + strconv.Itoa(port))
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
	log.Println(host)
	log.Printf("r.URL.Path: %s\n", r.URL.Path[1:])
	_, err := w.Write([]byte("hostname: " + host + "\n"))
	//if _, err = w.Write([]byte(host)); err != nil {
	//	return err
	//}
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(w, "r.URL.Path: %s\n", r.URL.Path[1:])
	return err

}

func GetIP(w http.ResponseWriter, r *http.Request) error {

	XRealIP := r.Header.Get("X-Real-Ip")
	_, err := fmt.Fprintf(w, "X-Real-Ip: %s\n", XRealIP)
	log.Printf("X-Real-Ip: %s\n", XRealIP)

	XForwardFor := r.Header.Get("X-Forwarded-For")
	_, err = fmt.Fprintf(w, "X-Forwarded-For: %s\n", XForwardFor)
	log.Printf("X-Forwarded-For: %s\n", XForwardFor)

	RemoteAddr := r.RemoteAddr
	_, err = fmt.Fprintf(w, "RemoteAddr: %s\n", RemoteAddr)
	log.Printf("RemoteAddr: %s\n", RemoteAddr)
	// add line to another request or consider use structured log zap
	log.Println()
	return err
}

func SetLog(log *logrus.Logger) *logrus.Logger {
	log.SetReportCaller(true)
	log.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", filename, f.Line)
		},
	}
	return log
}

func init() {
	SetLog(log)
}

// 面对疾风吧！
