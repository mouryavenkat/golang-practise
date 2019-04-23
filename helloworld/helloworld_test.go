package helloworld

import (
	"fmt"
	"testing"
)

func TestHelloworld(t *testing.T) {
	response := printHelloWorld()
	if response != "hello world" {
		t.Error(fmt.Sprintf("Expected = hello world, Got = %s\n", response))
	}
}
