package main

import (
	"runtime"

	"../../../panel/src/admin"
	"../controllers"
)

func main() {
	runtime.GOMAXPROCS(4)
	go admin.MakeAdminPanel()
	controllers.Register()
}
