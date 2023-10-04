package main

import (
	"github.com/verniyyy/battle-ship2/battleship2"
)

const listenPort = 9000

func main() {
	s := battleship2.NewHTTPServer(listenPort)
	s.Logger().Fatal(s.Serve())
}
