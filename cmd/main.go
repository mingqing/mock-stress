package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"

	"github.com/mingqing/mock-stress/modeler/cpu"
	"github.com/mingqing/mock-stress/modeler/memory"
)

func main() {
	cpu.UserTime(100, 7200*time.Second)

	memory.Allocate(80*1024*1024, 4096*1024, 120*time.Second)
}

func test() {
	src, err := os.Open("/dev/zero")
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	buf := make([]byte, 4096*1024*1024)

	_, err = io.Copy(ioutil.Discard, src)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	ioutil.Discard.Write(buf)
}
