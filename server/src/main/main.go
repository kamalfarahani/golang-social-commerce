package main

import (
	"runtime"

	"kamal/server/social-commerce/panel/src/admin"
	"kamal/server/social-commerce/server/src/controllers"
)

func main() {
	runtime.GOMAXPROCS(4)
	go admin.MakeAdminPanel()
	controllers.Register()
}
