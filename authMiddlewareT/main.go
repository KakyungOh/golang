package main

import (
	"authMiddleware/authenticate"
)

func main() {
	authenticate.NewServer().Run()
}
