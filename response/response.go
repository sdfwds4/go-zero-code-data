package response

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type rspBase struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data,omitempty"`
}

// Response 统一的响应函数，用于处理请求的响应，在逻辑处理完成后自动调用。
func Response(w http.ResponseWriter, resp any, err error) {
	var body rspBase
	if err != nil {
		body.Code = -1
		body.Msg = err.Error()
	} else {
		body.Msg = "ok"
		body.Data = resp
	}
	httpx.OkJson(w, body)
}
