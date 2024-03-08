package mySignals

import (
	"fmt"
	"syscall"
	"testing"
	"time"
)

func TestAll(t *testing.T) {
	go Listener()
	time.Sleep(time.Second)
	test_SIGUSR1()
	time.Sleep(time.Second)
	test_SIGUSR2()
	time.Sleep(time.Second)
	test_SIGHUP()
	time.Sleep(time.Second)
}

func test_SIGUSR1() {
	fmt.Println("Sending syscall.SIGUSR1")
	syscall.Kill(syscall.Getpid(), syscall.SIGUSR1)
}

func test_SIGUSR2() {
	fmt.Println("Sending syscall.SIGUSR2")
	syscall.Kill(syscall.Getpid(), syscall.SIGUSR2)
}

func test_SIGHUP() {
	fmt.Println("Sending syscall.SIGHUP")
	syscall.Kill(syscall.Getpid(), syscall.SIGHUP)
}
