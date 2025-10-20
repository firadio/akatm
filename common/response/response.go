package response

import (
	"net/http"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type Response struct {
	Code    int32       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func Success(w http.ResponseWriter, data interface{}) {
	httpx.OkJson(w, Response{Code: 0, Message: "success", Data: data})
}

func Error(w http.ResponseWriter, code int32, message string) {
	httpx.OkJson(w, Response{Code: code, Message: message})
}
