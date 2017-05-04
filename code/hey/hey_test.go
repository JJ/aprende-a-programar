package main

import (
	"testing"
	"strings"
)

func TestPage(t *testing.T) {
	page := MainPage() 
	if strings.LastIndex(page, "<") < 0 {
		t.Error("Estructura incorrecta")
	}
}

