package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("PID", os.Getpid())

	sigs := make(chan os.Signal)
	done := make(chan bool)

	signal.Notify(sigs, syscall.SIGQUIT, syscall.SIGKILL)

	go func() {
		sig := <-sigs
		fmt.Println("Signal received", sig)
		done <- true
	}()

	fmt.Println("Waiting for signal...")
	<-done
	fmt.Println("Signal received.")
}