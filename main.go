// main package
package main

import (
	"encoding/json"
	"log"
	"net/http"
	_ "net/http/pprof" //nolint:gosec
	"os"

	"github.com/gorilla/handlers"
	"github.com/julienschmidt/httprouter"

	"pprof-actions-go/internal/logic"
	"pprof-actions-go/internal/request"
	"pprof-actions-go/internal/response"
)

// Execute - основная функция скрипта.
func Execute(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Парсинг входящего JSON
	req := &request.Request{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Валидация объекта структуры Request.
	if err := request.ValidateStruct(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// создание словаря данных.
	data := logic.NewData(req)
	// вызов методов логики.
	data.LocalCount()
	data.ResultCount()
	// создание объекта ответа.
	resp := response.NewResponse(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp) //nolint:errcheck,gosec
}

func Health(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"}) //nolint:errcheck,gosec
}

func main() {
	router := httprouter.New()

	router.GET("/health", Health)
	router.POST("/execute", Execute)

	loggedRouter := handlers.LoggingHandler(os.Stdout, router)

	log.Println("Starting server...")

	log.Fatal(http.ListenAndServe(":8000", loggedRouter)) //nolint:gosec
}
