package go_zero_common_response

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type Resp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data,omitempty"`
}

func Response(w http.ResponseWriter, resp any, err error) {
	var r Resp
	if err != nil {
		r.Code = -1
		r.Msg = err.Error()
	} else {
		r.Msg = "Success"
		r.Data = resp
	}
	httpx.OkJson(w, r)
}
