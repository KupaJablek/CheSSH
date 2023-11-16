package main

import (
	"os"
	"testing"
)

func TestInputStringParserValidArgs(t *testing.T) {

	fName := "InputStringParser"

	os.Args = append(os.Args, "Test 1")

	test1 := []string{"main.go", "--hotseat", "-p1", "Player1", "-p2", "Player2"}
	if !InputStringParser(test1) {
		t.Fatalf(`%s(%s) = FALSE. expected result: TRUE`, fName, test1)
	}

	test2 := []string{"main.go", "--host", "-ip", "127.0.0.1"}
	if !InputStringParser(test2) {
		t.Fatalf(`%s(%s) = FALSE. expected result: TRUE`, fName, test2)
	}
}

func TestInputStringParserInvalidArgs(t *testing.T) {

	fName := "InputStringParser"

	os.Args = append(os.Args, "Test 1")

	test1 := []string{"main.go", "--hotseat", "-p1", "--p3"}
	if InputStringParser(test1) {
		t.Fatalf(`%s(%s) = TRUE. expected result: FALSE`, fName, test1)
	}

	test2 := []string{"main.go", "--hotseat", "-port", "127.0.0.1"}
	if InputStringParser(test2) {
		t.Fatalf(`%s(%s) = TRUE. expected result: FALSE`, fName, test2)
	}

	test3 := []string{"main.go", "--hotseat", "port", "ip"}
	if InputStringParser(test3) {
		t.Fatalf(`%s(%s) = TRUE. expected result: FALSE`, fName, test3)
	}
}
