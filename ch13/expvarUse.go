package main

import (
	"expvar"
	"fmt"
	"net/http"
)

func main() {
	// Define an integer variable to be exposed
	intVar := expvar.NewInt("intVar")
	intVar.Set(1234)

	// Register a custom function with expvar
	expvar.Publish("customFunction", expvar.Func(func() interface{} {
		return "Hi from Mastering Go!"
	}))

	// Register an additional HTTP handler for expvar
	http.Handle("/debug/expvars", expvar.Handler())

	// Start an HTTP server to expose variables
	go func() {
		fmt.Println("HTTP server listening on :8080")
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			fmt.Println("Error starting HTTP server:", err)
		}
	}()

	// Your application logic here...
	// You can update the variable over time
	intVar.Add(10)

	// Keep the main goroutine running
	select {}
}
