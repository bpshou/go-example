package main

import (
	"gin_app/app/core"
	"gin_app/app/router"
)

func init() {
	core.Init()
}

func main() {
	router.Run()
}
