package main

import (
	"fmt"
	"fs/Plugins"
	"fs/common"
	"time"
)

func main() {
	start := time.Now()
	var Info common.HostInfo
	common.Flag(&Info)
	common.Parse(&Info)
	Plugins.Scan(Info)
	t := time.Now().Sub(start)
	fmt.Printf("[*] 扫描结束,耗时: %s\n", t)
}
