package main

import (
	"log/slog"
	"os"
	"time"
)

func myFunction() {
	j := 0
	for i := 1; i < 100000000; i++ {
		j = j % i
	}
}

func main() {
	handler := slog.NewTextHandler(os.Stdout, nil)
	logger := slog.New(handler)
	logger.Debug("This is a DEBUG message")

	for i := 0; i < 5; i++ {
		now := time.Now()
		myFunction()
		elapsed := time.Since(now)
		logger.Info(
			"Observability",
			slog.Int64("time_taken", int64(elapsed)),
		)
	}
}
