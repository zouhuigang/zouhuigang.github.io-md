/*
代理服务器客户端，用于连接远程代理服务器
*/
package model

import (
	"github.com/elazarl/goproxy"
	"log"
	"net/http"
	"net/url"
	"time"
)

func ProxyHTTP(httpAddr string) (*goproxy.ProxyHttpServer, error) {
	proxyURL, err := url.Parse(httpAddr)
	if err != nil {
		return nil, err
	}
	log.Printf("正在使用代理地址: %v\n", proxyURL)

	prox := goproxy.NewProxyHttpServer()
	prox.Verbose = true
	prox.Tr.Proxy = http.ProxyURL(proxyURL)

	//超时
	prox.OnRequest(goproxy.DstHostIs("www.reddit.com")).DoFunc(
		func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
			h, _, _ := time.Now().Clock()
			if h >= 8 && h <= 17 {
				return r, goproxy.NewResponse(r,
					goproxy.ContentTypeText, http.StatusForbidden,
					"Don't waste your time!")
			} else {
				ctx.Warnf("clock: %d, you can waste your time...", h)
			}
			return r, nil
		})

	prox.OnResponse().DoFunc(func(resp *http.Response, ctx *goproxy.ProxyCtx) *http.Response {
		log.Printf("zzz:%v", resp)
		return resp
	})

	return prox, nil
}
