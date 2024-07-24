package main

import (
	"github.com/NavaRose/gogogo-core/core"
)

func main() {
	engine := core.InitEngine(RouteCreator)
	err := engine.Run(":80")
	if err != nil {
		ErrorHandle(err)
	}
}
