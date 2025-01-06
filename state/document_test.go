package state

import "testing"

// TODO test mixed cases

func TestGetTokenFromTextEmpty(t *testing.T) {
	text := "hello world from Jargon LSP"

	var i uint = 5
	token := getTokenFromText(text, i)

	if token != nil {
		t.Fatalf("got invalid token for i = %d: '%s'", i, *token)
	}
}

func TestGetTokenFromTextBeginning(t *testing.T) {
	text := "hello world from Jargon LSP"

	var i uint
	for i = 0; i <= 4; i++ {
		token := getTokenFromText(text, i)

		if token == nil {
			t.Fatalf("got invalid token for i = %d: nil", i)
		}

		if *token != "hello" {
			t.Fatalf("got invalid token for i = %d: '%s'", i, *token)
		}
	}
}

func TestGetTokenFromTextMiddleSimple(t *testing.T) {
	text := "hello world from Jargon LSP"

	var i uint
	for i = 6; i <= 10; i++ {
		token := getTokenFromText(text, i)

		if token == nil {
			t.Fatalf("got invalid token for i = %d: nil", i)
		}

		if *token != "world" {
			t.Fatalf("got invalid token for i = %d: '%s'", i, *token)
		}
	}
}

func TestGetTokenFromTextMiddleCapitalized(t *testing.T) {
	text := "hello world from Jargon LSP"

	var i uint
	for i = 17; i <= 22; i++ {
		token := getTokenFromText(text, i)

		if token == nil {
			t.Fatalf("got invalid token for i = %d: nil", i)
		}

		if *token != "Jargon" {
			t.Fatalf("got invalid token for i = %d: '%s'", i, *token)
		}
	}
}

func TestGetTokenFromTextEnd(t *testing.T) {
	text := "hello world from Jargon LSP"

	var i uint
	for i = 24; i <= 26; i++ {
		token := getTokenFromText(text, i)

		if token == nil {
			t.Fatalf("got invalid token for i = %d: nil", i)
		}

		if *token != "LSP" {
			t.Fatalf("got invalid token for i = %d: '%s'", i, *token)
		}
	}
}
