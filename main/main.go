package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
	HttpServerStart()
}

func HttpServerStart() {

	log.SetPrefix("Info:")
	log.SetFlags(log.Ldate | log.Llongfile)

	mux := http.NewServeMux()
	mux.HandleFunc("/", accessFunc)
	mux.HandleFunc("/healthz", healthz)

	err := http.ListenAndServe(":80", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func accessFunc(response http.ResponseWriter, request *http.Request) {
	if len(response.Header()) > 0 {
		for k, v := range response.Header() {
			log.Printf("%s=%s", k, v[0])

			response.Header().Set(k, v[0])
		}
	}

	log.Printf("\n")

	err := request.ParseForm()
	if err != nil {
		return
	}

	if len(request.Form) > 0 {
		for k, v := range request.Form {
			log.Printf("%s=%s", k, v[0])
		}
	}

	log.Printf("\n")

	err = os.Setenv("VERSION", "JDK version 1.8")
	if err != nil {
		return
	}

	name := os.Getenv("VERSION")
	log.Printf("VERSION env:", name)

	log.Printf("\n")

	ip, _, err := net.SplitHostPort(request.RemoteAddr)
	if err != nil {
		fmt.Println("err:", err)
	}

	if net.ParseIP(ip) != nil {
		fmt.Printf("ip:%s\n", ip)
		log.Println(ip)
	}

	fmt.Printf("http status Code ===>> %s\n", http.StatusOK)
	log.Println(http.StatusOK)

	response.WriteHeader(http.StatusOK)
	response.Write([]byte("Server Access,Success!"))
}

func healthz(response http.ResponseWriter, request *http.Request) {
	HealthzCode := "200"
	response.Write([]byte(HealthzCode))
}
