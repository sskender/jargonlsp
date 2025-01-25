package state

import (
	"testing"
)

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

func TestGetTokenWithUnderscore(t *testing.T) {
	text := "hello _world from_ Jargon_LSP"

	var i uint
	for i = 6; i <= 11; i++ {
		token := getTokenFromText(text, i)

		if token == nil {
			t.Fatalf("got invalid token for i = %d: nil", i)
		}

		if *token != "_world" {
			t.Fatalf("got invalid token for i = %d: '%s'", i, *token)
		}
	}

	for i = 13; i <= 17; i++ {
		token := getTokenFromText(text, i)

		if token == nil {
			t.Fatalf("got invalid token for i = %d: nil", i)
		}

		if *token != "from_" {
			t.Fatalf("got invalid token for i = %d: '%s'", i, *token)
		}
	}

	for i = 19; i <= 28; i++ {
		token := getTokenFromText(text, i)

		if token == nil {
			t.Fatalf("got invalid token for i = %d: nil", i)
		}

		if *token != "Jargon_LSP" {
			t.Fatalf("got invalid token for i = %d: '%s'", i, *token)
		}
	}
}

func TestGetTokenWithNumbers(t *testing.T) {
	text := "hello world0 from Jargon1_0_ LSP"

	var i uint
	for i = 6; i <= 11; i++ {
		token := getTokenFromText(text, i)

		if token == nil {
			t.Fatalf("got invalid token for i = %d: nil", i)
		}

		if *token != "world0" {
			t.Fatalf("got invalid token for i = %d: '%s'", i, *token)
		}
	}

	for i = 18; i <= 27; i++ {
		token := getTokenFromText(text, i)

		if token == nil {
			t.Fatalf("got invalid token for i = %d: nil", i)
		}

		if *token != "Jargon1_0_" {
			t.Fatalf("got invalid token for i = %d: '%s'", i, *token)
		}
	}
}

func TestGetTokenWithValidSpecialCharacter(t *testing.T) {
	text := "var $token = \"hello world\""

	var i uint
	for i = 5; i <= 9; i++ {
		token := getTokenFromText(text, i)

		if token == nil {
			t.Fatalf("got invalid token for i = %d: nil", i)
		}

		if *token != "token" {
			t.Fatalf("got invalid token for i = %d: '%s'", i, *token)
		}
	}
}

func TestGetTokenWithSurrounding(t *testing.T) {
	text := "hello [\"world\"] token"

	var i uint
	for i = 8; i <= 12; i++ {
		token := getTokenFromText(text, i)

		if token == nil {
			t.Fatalf("got invalid token for i = %d: nil", i)
		}

		if *token != "world" {
			t.Fatalf("got invalid token for i = %d: '%s'", i, *token)
		}
	}

	var token *string

	i = 7
	token = getTokenFromText(text, i)
	if token != nil {
		t.Fatalf("got invalid token for i = %d: '%s'", i, *token)
	}

	i = 13
	token = getTokenFromText(text, i)
	if token != nil {
		t.Fatalf("got invalid token for i = %d: '%s'", i, *token)
	}
}
