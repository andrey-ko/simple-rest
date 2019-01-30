package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

func getHostIps() ([]string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	ips := make([]string, 0, len(ifaces))
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			return nil, err
		}
		for _, addr := range addrs {
			ips = append(ips, addr.String())
		}
	}
	return ips, nil
}
func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("processing request from '%s'", r.RemoteAddr)
	_, err := fmt.Fprint(w, "simple rest v0.0\n")
	if err != nil {
		log.Printf("failed to write response: %s", err)
	}
	host, err := os.Hostname()
	_, err = fmt.Fprintf(w, "host name: %s\n", host)
	if err != nil {
		log.Printf("failed to write response: %s", err)
	}
	ips, err := getHostIps()
	if err != nil {
		log.Printf("failed to get ip addresses: %s", err)
	}
	_, err = fmt.Fprintf(w, "host ip: %s\n", strings.Join(ips, ", "))
	if err != nil {
		log.Printf("failed to write response: %s", err)
	}
}

func main() {
	log.Printf("pid: '%d'...", os.Getpid())
	addr := flag.String("addr", ":8080", "listening address")
	flag.Parse()

	http.HandleFunc("/", handler)

	log.Printf("listening on '%s'...", *addr)
	err := http.ListenAndServe(*addr, nil)
	log.Printf("exiting...: %s", err)
}
