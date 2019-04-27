package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	intrpsig := make(chan os.Signal)
	signal.Notify(intrpsig, os.Interrupt)
	tick := time.Tick(1 * time.Second)
	for {
		select {
		case <-tick:
			fmt.Println(time.Now())
		case <-intrpsig:
			fmt.Println("Yeee!!! Received interrupt signal")
			os.Exit(1)
		}
	}

}
