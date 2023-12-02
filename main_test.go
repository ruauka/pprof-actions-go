package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func BenchmarkExecute(b *testing.B) {
	payload := prepare()
	// Test server
	router := httprouter.New()
	router.POST("/execute", Execute)
	// Test Request
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/execute", bytes.NewBuffer(payload))

	for i := 0; i < b.N; i++ {
		Execute(w, req, httprouter.Params{})
	}
}

func prepare() []byte {
	data, err := os.ReadFile("testdata/input.json")
	if err != nil {
		fmt.Println(err)
	}

	return data
}
