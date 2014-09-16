package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c)
	sig := (<-c).(syscall.Signal)
	fmt.Printf("%#x (%s)\n", int(sig), sig)
}
