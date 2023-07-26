package http

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zeromicro/go-zero/core/logx"
	errorx "github.com/zeromicro/x/errors"
	"github.com/zeromicro/x/test"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestMain(m *testing.M) {
	logx.Disable()
	m.Run()
}

func TestJsonBaseResponse(t *testing.T) {
	executor := test.NewExecutor[any, testWriterResult](comparisonOption)
	executor.Add([]test.Data[any, testWriterResult]{
		{
			Name:  "code-msg-pointer",
			Input: errorx.New(1, "test"),
			Want: testWriterResult{
				code:        200,
				writeString: `{"code":1,"msg":"test"}`,
			},
		},
		{
			Name:  "code-msg-struct",
			Input: errorx.CodeMsg{Code: 1, Msg: "test"},
			Want: testWriterResult{
				code:        200,
				writeString: `{"code":1,"msg":"test"}`,
			},
		},
		{
			Name:  "status.Status",
			Input: status.New(codes.OK, "ok"),
			Want: testWriterResult{
				code:        200,
				writeString: `{"code":0,"msg":"ok"}`,
			},
		},
		{
			Name:  "error",
			Input: errors.New("test"),
			Want: testWriterResult{
				code:        200,
				writeString: `{"code":-1,"msg":"test"}`,
			},
		},
		{
			Name:  "struct",
			Input: message{Name: "anyone"},
			Want: testWriterResult{
				code:        200,
				writeString: `{"code":0,"msg":"ok","data":{"name":"anyone"}}`,
			},
		},
	}...)
	executor.RunE(t, func(a any) (testWriterResult, error) {
		w := &tracedResponseWriter{headers: make(map[string][]string)}
		JsonBaseResponse(w, a)
		return w.result()
	})
}

func TestJsonBaseResponseCtx(t *testing.T) {
	executor := test.NewExecutor[any, testWriterResult](comparisonOption)
	executor.Add([]test.Data[any, testWriterResult]{
		{
			Name:  "code-msg-pointer",
			Input: errorx.New(1, "test"),
			Want: testWriterResult{
				code:        200,
				writeString: `{"code":1,"msg":"test"}`,
			},
		},
		{
			Name:  "code-msg-struct",
			Input: errorx.CodeMsg{Code: 1, Msg: "test"},
			Want: testWriterResult{
				code:        200,
				writeString: `{"code":1,"msg":"test"}`,
			},
		},
		{
			Name:  "status.Status",
			Input: status.New(codes.OK, "ok"),
			Want: testWriterResult{
				code:        200,
				writeString: `{"code":0,"msg":"ok"}`,
			},
		},
		{
			Name:  "error",
			Input: errors.New("test"),
			Want: testWriterResult{
				code:        200,
				writeString: `{"code":-1,"msg":"test"}`,
			},
		},
		{
			Name:  "struct",
			Input: message{Name: "anyone"},
			Want: testWriterResult{
				code:        200,
				writeString: `{"code":0,"msg":"ok","data":{"name":"anyone"}}`,
			},
		},
	}...)
	executor.RunE(t, func(a any) (testWriterResult, error) {
		w := &tracedResponseWriter{headers: make(map[string][]string)}
		JsonBaseResponseCtx(context.TODO(), w, a)
		return w.result()
	})
}

var comparisonOption = test.WithComparison[any, testWriterResult](func(t *testing.T, expected, actual testWriterResult) {
	assert.Equal(t, expected.code, actual.code)
	assert.Equal(t, expected.writeString, actual.writeString)
})
