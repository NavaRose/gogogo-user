package main

func main() {
	engine := core.InitEngine(RouteCreator)

	err := engine.Run(":80")
	if err != nil {
		ErrorHandle(err)
	}
}
