package main

import (
	"rpc-server/test"
)

func main() {
    test.RegisterRoutes()
	test.InitServer()

}
