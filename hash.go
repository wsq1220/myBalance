package main

import (
	"fmt"
	"hash/crc64"
	"math/rand"

	"github.com/wsq1220/myBalance/balance"
)

type HashBalance struct {
}

func init() {
	balance.RegisterBalancer("hash", &HashBalance{})
	fmt.Println("register hash balance...")
}

func (p *HashBalance) DoBalance(insts []*balance.Instance, key ...string) (inst *balance.Instance, err error) {
	var defKey string = fmt.Sprintf("%d", rand.Int())
	if len(key) > 0 {
		defKey = key[0]
	}

	lens := len(insts)
	if lens == 0 {
		fmt.Println("No backend nstance1")
		return
	}
	crcTable := crc64.MakeTable(crc64.ECMA)
	hashVal := crc64.Checksum([]byte(defKey), crcTable)
	index := int(hashVal) % lens
	inst = insts[index]

	return
}
