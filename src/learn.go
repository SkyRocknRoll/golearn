package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"net/http"
)

func main() {
	http.HandleFunc("/host", hostFinder)
	http.HandleFunc("/ip", sourceIp)
	fmt.Println("Listening on 0.0.0.0:8000")
	err := http.ListenAndServe("0.0.0.0:8000", nil)
	if err != nil {
		log.Fatal(err)
	}

}

func hostFinder(rw http.ResponseWriter, req *http.Request) {
	hostPort := req.Host
	host, port, _ := net.SplitHostPort(hostPort)
	io.WriteString(rw, "<html>")
	io.WriteString(rw, "Host Name : "+host+" <br>")
	io.WriteString(rw, "Port Number : "+port)
	io.WriteString(rw, "</br>")
	io.WriteString(rw, req.RemoteAddr)
	io.WriteString(rw, "</html>")
	// io.WriteString(w, s)

}

func sourceIp(rw http.ResponseWriter, req *http.Request) {
	host, _, _ := net.SplitHostPort(req.RemoteAddr)
	io.WriteString(rw, "<h1> Your Ip address : "+host)
}
