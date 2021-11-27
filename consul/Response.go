package consul

import (
	"io/ioutil"
	"net/http"
)

type Response struct {
	Status     string
	StatusCode int
	Body       []byte
}

func NewResponse(res *http.Response) *Response {
	// 读取请求后的返回的数据 是 []byte类型
	body, _ := ioutil.ReadAll(res.Body)
	// 必须要关闭不然可能会存在内存溢出的问题
	defer res.Body.Close()
	return &Response{
		Status:     res.Status,
		StatusCode: res.StatusCode,
		Body:       body,
	}
}
