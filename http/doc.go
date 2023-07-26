// Package http provides a set of functions for working with HTTP requests and responses.
// It is an extension of the go-zero package https://github.com/zeromicro/go-zero/tree/master/rest/httpx,
// it provides a base response struct, as before, you can only respond a json string like this:
//
//	httpx.OkJson(w, message{Name: "anyone"})
//
// then you can receive a json string like this:
//
//	{"name":"anyone"}
//
// but now, you can respond a json string with base response struct like this:
//
//	JsonBaseResponse(w, message{Name: "anyone"})
//
// then you can receive a json string like this:
//
//	{"code":0,"msg":"ok","data":{"name":"anyone"}}
package http
