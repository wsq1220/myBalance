package balance

import "fmt"

type BalancerMgr struct {
	allBalancer map[string]Balancer
}

func init() {
	
}

var (
	mgr = BalancerMgr{
		allBalancer: make(map[string]Balancer),
	}
)

func (p *BalancerMgr) registerBalancer(name string, b Balancer) {
	p.allBalancer[name] = b
}

func RegisterBalancer(name string, b Balancer) {
	mgr.registerBalancer(name, b)
}

func DoBalance(name string, insts []*Instance) (inst *Instance, err error) {
	balancer, ok := mgr.allBalancer[name]
	if !ok {
		err = fmt.Errorf("Not found %s balancer", name)
		return
	}

	fmt.Printf("use %s balancer", name)
	inst, err = balancer.DoBalance(insts)
	if err != nil {
		fmt.Println("balancer do balanlance failed, err:", err)
		return
	}

	return
}
