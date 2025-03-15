package main

import (
	"Blog-Server/utils/computer"
	"fmt"
)

func main() {
	fmt.Println(computer.GetCpuPercent())
	fmt.Println(computer.GetMemPercent())
	fmt.Println(computer.GetDiskPercent())
}
