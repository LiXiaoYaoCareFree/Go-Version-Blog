package main

import (
	"Blog-Server/core"
	"fmt"
)

func main() {
	core.InitIPDB()
	fmt.Println(core.GetIpAddr("175.0.201.207"))

}
