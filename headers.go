package genshin

import "net/http"

func getHeaders(server string, ds string, cookie string) http.Header {
	switch server {
	case "os":
		return http.Header{
			"DS": {ds},
    		"Origin": {"https://webstatic-sea.hoyolab.com"},
    		"Referer": {"https://webstatic-sea.hoyolab.com/"},
			"Accept": {"application/json, text/plain, */*"},
			"Content-Type": {"application/json;charset=utf-8"},
			"Accept-Language": {"zh-CN,en-US;q=0.8"},
			"x-rpc-language": {"en-us"},
			"x-rpc-app_version": {"1.5.0"},
			"x-rpc-client_type": {"5"},
			"User-Agent": {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.93 Safari/537.36"},
			"Cookie": {cookie},
		}
	default:
		return http.Header{}
	}
}