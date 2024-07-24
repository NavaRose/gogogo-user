package main

func main() {
	engine := InitEngine(RouteCreator)

	err := engine.Run(":80")
	if err != nil {
		ErrorHandle(err)
	}
}
