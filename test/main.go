package main

import (
	"KAnsible/ansible"
	"fmt"
)

func main() {
	revMsg := make(chan string)
	ok := ansible.RunPlaybook(revMsg,"data/kubespray/inventory/hosts.yaml",
										"data/kubespray/test.yaml")
	if ok {
		fmt.Println("执行成功")
	} else {
		fmt.Println("执行失败")
	}
}
