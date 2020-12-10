package parser

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	primary, args := Scan("    go   help")

	println("Primary:")
	fmt.Println(primary)

	println("\nArguments:")
	fmt.Println(args)

	if primary != "go" {
		t.Errorf("expected primary to be 'go', instead I got %v", primary)
	}

	if args[0] != "help" {
		t.Errorf("expected primary to be 'help', got %v", args[0])
	}
}
