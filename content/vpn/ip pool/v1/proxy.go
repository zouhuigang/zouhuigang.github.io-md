package main

import (
	"github.com/elazarl/goproxy"
	//"golang.org/x/net/proxy"
	"log"
	"net/http"
	"net/url"
)

func proxyHTTP(httpAddr string) (*goproxy.ProxyHttpServer, error) {
	proxyURL, err := url.Parse(httpAddr)
	if err != nil {
		return nil, err
	}
	log.Printf("New HTTP proxy Host: %s, Port: %v", proxyURL.Host, proxyURL)

	prox := goproxy.NewProxyHttpServer()
	prox.Verbose = true
	prox.Tr.Proxy = http.ProxyURL(proxyURL)

	return prox, nil
}

func main() {
	//proxy := goproxy.NewProxyHttpServer()
	//proxy.Verbose = true
	//139.196.48.36 6111
	proxy, _ := proxyHTTP("http://61.155.164.111:3128")

	log.Fatal(http.ListenAndServe(":8080", proxy))
}
