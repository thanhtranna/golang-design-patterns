package main

import (
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	number := rand.Intn(10)

	t := AContext{
		CurrentState: &Ask{},
		Number:       number,
	}

	t.prntState()

	for t.CurrentState.Exec(&t) {
		t.prntState()
	}
}
