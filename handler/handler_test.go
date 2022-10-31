package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/require"
)

func TestHealthCheck(t *testing.T) {
	// Test server
	router := httprouter.New()
	router.POST("/health", HealthCheck)
	// Test Request
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/health", nil)

	HealthCheck(w, r, httprouter.Params{})

	require.Equal(t, http.StatusOK, w.Code)
}

func TestAdd(t *testing.T) {
	TestTable := []struct {
		jsonIn       []byte
		expectedCode int
		//expectedResp map[string]int
		testName string
	}{
		{
			jsonIn: []byte(`{
				"number_one": 10,
				"number_two": 5
			}`),
			expectedCode: http.StatusOK,
			//expectedResp: map[string]int{"result": 15},
			testName: "test_1_OK",
		},
		{
			jsonIn: []byte(`{
				"number_one": 10,
				"number_two": 5,
			}`),
			expectedCode: http.StatusOK,
			testName:     "test_2_err decode",
		},
	}

	for _, testCase := range TestTable {
		t.Run(testCase.testName, func(t *testing.T) {
			payload := bytes.NewBuffer(testCase.jsonIn)
			// Test server
			router := httprouter.New()
			router.POST("/add", Add)
			// Test Request
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/add", payload)

			Add(w, r, httprouter.Params{})

			require.Equal(t, testCase.expectedCode, w.Code)
			//require.Equal(t, testCase.expectedResp, w.Body.String())
		})
	}
}
