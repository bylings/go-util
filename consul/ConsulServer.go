package consul

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

type ConsulServer struct {
	host string
	port int
}

func NewConsulServer(host string, port int) *ConsulServer {
	return &ConsulServer{
		host: host,
		port: port,
	}
}

func (c *ConsulServer) Get(url string, options map[string]interface{}) (*Response, error) {
	return c.request("GET", url, options)
}

func (c *ConsulServer) Put(url string, options map[string]interface{}) (*Response, error) {
	return c.request("PUT", url, options)
}

// 封装请求
func (c *ConsulServer) request(method, url string, options map[string]interface{}) (*Response, error) {
	// 构建uri
	uri := "http://" + c.host + ":" + strconv.Itoa(c.port) + url
	//fmt.Println("url  ", uri)

	// 根据传输的参数判断
	var req *http.Request
	if options != nil {
		s, _ := json.Marshal(options)
		req, _ = http.NewRequest(method, uri, bytes.NewReader(s))
	} else {
		req, _ = http.NewRequest(method, uri, nil)
	}
	res, err := http.DefaultClient.Do(req)
	if res == nil {
		return nil, errors.New(uri + " 请求失败，请检查地址和端口：" + err.Error())
	}
	return NewResponse(res), nil
}
