package balance

import (
	"fmt"
	"hash/crc32"
	"math/rand"
)

type HashBalance struct {
}

func init() {
	RegisterBalancer("hash", &HashBalance{})
	fmt.Println("register hash balance...")
}

func (p *HashBalance) DoBalance(insts []*Instance, key ...string) (inst *Instance, err error) {
	var defKey string = fmt.Sprintf("%d", rand.Int())
	if len(key) > 0 {
		defKey = key[0]
	}

	lens := len(insts)
	if lens == 0 {
		fmt.Println("No backend nstance1")
		return
	}
	// crcTable := crc64.MakeTable(crc64.ECMA)
	// hashVal := crc64.Checksum([]byte(defKey), crcTable)
	crcTable := crc32.MakeTable(crc32.IEEE)
	hashVal := crc32.Checksum([]byte(defKey), crcTable)
	index := int(hashVal) % lens
	inst = insts[index]

	return
}
