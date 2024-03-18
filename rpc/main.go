package main

import (
	"su/common"
	"su/register"
)

func main() {
	common.NacosRegister()
	register.InitRegister()
}
