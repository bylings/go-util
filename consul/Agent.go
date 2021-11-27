package consul

// 对服务的操作方法
type Agent struct {
	consul *ConsulServer
	params map[string]interface{}
}

// 实例化
func NewAgent(host string, port int) *Agent {
	return &Agent{
		consul: NewConsulServer(host, port),
	}
}

// 查询所有服务
func (a *Agent) GetServices() (*Response, error) {
	return a.consul.Get("/v1/agent/services", a.params)
}

// 注册服务
func (a *Agent) RegisterService() (*Response, error) {
	return a.consul.PUT("/v1/agent/service/register", a.params)
}
