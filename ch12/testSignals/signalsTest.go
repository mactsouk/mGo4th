package mySignals

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func HandleSignal(sig os.Signal) {
	fmt.Println("handleSignal() Caught:", sig)
}

func Listener() {
	fmt.Printf("Process ID: %d\n", os.Getpid())
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs)
	start := time.Now()

	go func() {
		for {
			sig := <-sigs
			switch sig {
			case syscall.SIGUSR1:
				HandleSignal(sig)
				duration := time.Since(start)
				fmt.Println("Execution time:", duration)
			case syscall.SIGUSR2:
				HandleSignal(sig)
			default:
				fmt.Println("Caught:", sig)
			}
		}
	}()

	for {
		time.Sleep(10 * time.Second)
	}
}
