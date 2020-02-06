package main

import (
	"net"
	"net/http"
)

func readUserIP(r *http.Request) string {
	ip := r.Header.Get("X-Real-Ip")
	if ip == "" {
		ip = r.Header.Get("X-Forwarded-For")
	}
	if ip == "" {
		ip, _, _ = net.SplitHostPort(r.RemoteAddr)
	}
	return ip
}

func readUserAgent(r *http.Request) string {
	return r.UserAgent()
}

// Ternary function
func Ternary(statement bool, a, b interface{}) interface{} {
	if statement {
		return a
	}
	return b
}
