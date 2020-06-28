package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/wsq1220/myBalance/balance"
)

func main() {
	// insts := make([]*balance.Instance)
	var insts []*balance.Instance
	for i := 0; i < 16; i++ {
		host := fmt.Sprintf("192.168.%d.%d", rand.Intn(255), rand.Intn(255))
		one := balance.NewInstance(host, 8080)
		insts = append(insts, one)
	}

	// var balancer balance.Balancer
	// var balanceName string
	var balanceName = "random"
	if len(os.Args) > 1 {
		balanceName = os.Args[1]
	}

	// if balanceName == "random" {
	// 	balancer = &balance.RandomBalance{}
	// 	fmt.Println("use random balancer")
	// } else if balanceName == "roundrobin" {
	// 	balancer = &balance.RoundRobinBalance{}
	// 	fmt.Println("use roundrobin balancer")
	// }
	// balancer := balance.RandomBalance{}
	// balancer := balance.RoundRobinBalance{}
	for {
		inst, err := balance.DoBalance(balanceName, insts)
		if err != nil {
			fmt.Println("dobalance failed, err:", err)
		}
		fmt.Println(inst)
		time.Sleep(time.Second)
	}
}
