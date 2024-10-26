package main

import (
	"gobuild/luna"
)

func main() {

	engine, err := luna.New(
		luna.Config{
			ENV:         "development",
			FrontendDir: "frontend",
		},
	)
	if err != nil {
		panic(err)
	}

	engine.FrontEnd()
	engine.Server.Start(":8080")
}
