package consul

import (
	"fmt"
	"testing"
)

func TestServices(t *testing.T) {
	agent := NewAgent("ip地址", 8500)
	res, _ := agent.Services()
	fmt.Println(string(res.Body))
}
