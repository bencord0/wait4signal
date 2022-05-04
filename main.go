package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	// A buffered channel to receive signals
	signals := make(chan os.Signal, 1)

	// Notify on all signals
	signal.Notify(signals)

	// Block until any signal
	sig := <-signals
	fmt.Println("Signal:", sig)
}
