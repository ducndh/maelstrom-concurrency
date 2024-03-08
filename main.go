package main

import (
	"log"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func main() {
	n := maelstrom.NewNode()
	s := &server{n: n}
	n.Handle("echo", s.echo)

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}
