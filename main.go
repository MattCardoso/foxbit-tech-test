package main

import (
	"github.com/mattcardoso/foxbit-tech-test/web"
)

func main() {
	srv := web.NewServer()
	srv.Run()
}
