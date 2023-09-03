package api_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/anotherandrey/token-rest-api/internal/app/api"
)

var router = api.NewApi().Router

func TestHandleWhoAmIWithNoToken(t *testing.T) {
	recorder := httptest.NewRecorder()

	r, _ := http.NewRequest("GET", "/api/v1/whoami", nil)

	router.ServeHTTP(recorder, r)

	if recorder.Code != http.StatusUnauthorized {
		t.Errorf("Code must be %d", http.StatusUnauthorized)
	}
}

func TestHandleWhoAmIWithWrongToken(t *testing.T) {
	recorder := httptest.NewRecorder()

	r, _ := http.NewRequest("GET", "/api/v1/whoami", nil)
	r.Header.Add("Authorization", "Bearer 42")

	router.ServeHTTP(recorder, r)

	if recorder.Code != http.StatusUnauthorized {
		t.Errorf("Code must be %d", http.StatusUnauthorized)
	}
}
