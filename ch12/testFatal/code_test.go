package server

import (
	"testing"
)

func Test_map(t *testing.T) {
	key := "server"
	server, ok := DATA[key]
	if !ok {
		t.Fatalf("Key %s not found!", key)
	}

	key = "port"
	port, ok := DATA[key]
	if !ok {
		t.Fatalf("Key %s not found!", key)
	}

	t.Log("Connecting to", server, "@port", port)
}
