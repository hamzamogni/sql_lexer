package main

type Token int

const (
	// Special tokens
	ILLEGAL Token = iota
	EOF
	WS

	// Literals
	IDENT

	// Misc caracters
	ASTERIX
	COMMA

	// Keywords
	SELECT
	FROM
)
