package win

import (
	"log"
	"testing"
)

func TestNewProcessByName(t *testing.T) {
	p, err := NewProcessByName("WeChat.exe")
	if err != nil {
		log.Printf("err: %+v", err)
		return
	}

	var value int
	//err = p.ReadProcessMemory(uintptr(0x6777D9E8), 4, &value)
	//if err != nil {
	//	log.Printf("err: %+v", err)
	//	return
	//}

	var addrs = []uintptr{
		0x677700E0,
		0x6777D90C,
		0x6777D9E8,
		0x67793E4C,
		0x67795AA4,
		0x677985D4,
	}

	//log.Printf("value: %+v", value)

	value = 0x63090B11
	for _, addr := range addrs {
		err = p.WriteProcessMemory(addr, 4, &value)
		if err != nil {
			log.Printf("err: %+v", err)
			return
		}
	}
}
