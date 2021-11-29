package consul

import (
	"fmt"
	"testing"
)

func TestRegisterServer(t *testing.T) {
	servers := map[string]interface{}{"ID": "go_consul_test", "Name": "go1_test", "Address": "127.0.0.1", "Port": 8300}

	fmt.Println("servers : ", servers)

	agent := NewAgent("host地址", 8550)
	res, err := agent.RegisterService(servers)
	fmt.Println("res : ", res)
	fmt.Println("err : ", err)
}

func TestServices(t *testing.T) {
	agent := NewAgent("host地址", 8550)
	res, _ := agent.GetServices()
	fmt.Println(string(res.Body))
}
