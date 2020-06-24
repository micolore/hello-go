package main

import (
	"fmt"
	"os/exec"
)

func  main(){
	cmd  := exec.Command("bash","-c","df -hl")
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n",string(out))

	// 执行perl 脚本文件
	cmd=exec.Command("bash","-c","perl hl.pl")
	out, err = cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n",string(out))
}