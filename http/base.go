package http

import (
	"context"
	"net/http"

	"github.com/htquangg/x/errors"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
)

// BaseResponse is the base response struct.
type BaseResponse[T any] struct {
	// Code represents the business code, not the http status code.
	Code int `json:"code"`
	// Msg represents the business message, if Code = BusinessCodeOK,
	// and Msg is empty, then the Msg will be set to BusinessMsgOk.
	Msg string `json:"msg"`
	// Data represents the business data.
	Data T `json:"data,omitempty"`
}

func Success(ctx context.Context, w http.ResponseWriter, v any) {
	httpx.OkJsonCtx(ctx, w, wrapSuccessResponse(v))
}

func Error(ctx context.Context, w http.ResponseWriter, v any) {
	httpx.OkJsonCtx(ctx, w, wrapErrorResponse(v))
}

func wrapSuccessResponse(v any) BaseResponse[any] {
	var resp BaseResponse[any]

	resp.Code = BusinessCodeOK
	resp.Msg = BusinessMsgOk
	switch data := v.(type) {
	case *status.Status:
		resp.Data = data.Message()
	default:
		resp.Data = v
	}

	return resp
}

func wrapErrorResponse(v any) BaseResponse[any] {
	var resp BaseResponse[any]

	switch data := v.(type) {
	case *errors.CodeMsg:
		resp.Code = data.Code
		resp.Msg = data.Msg
	case errors.CodeMsg:
		resp.Code = data.Code
		resp.Msg = data.Msg
	case *status.Status:
		resp.Code = BusinessCodeError
		resp.Msg = data.Message()
	case error:
		resp.Code = BusinessCodeError
		resp.Msg = data.Error()
	default:
		resp.Code = BusinessCodeError
		resp.Msg = BusinessMsgError
	}

	return resp
}
