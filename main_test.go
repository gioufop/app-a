package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler(t *testing.T) {
	// Cria uma requisição HTTP simulada
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	// Chama a função
	handler(w, req)

	// Verifica o resultado
	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)
	expected := "Hello World! I'm GO app-a!"

	if !strings.Contains(string(body), expected) {
		t.Errorf("Esperado '%s', recebido '%s'", expected, string(body))
	}
}