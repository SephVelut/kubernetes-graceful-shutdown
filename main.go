package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	tick := time.NewTicker(time.Second * 600).C
	sigChan := make(chan os.Signal, 1)

	signal.Notify(sigChan, syscall.SIGTERM)

	for {
		select {
		case <-sigChan:
			fmt.Println("killed")
			os.Exit(0)
		case <-tick:
			fmt.Println("finished")
			os.Exit(0)
		}
	}
}
