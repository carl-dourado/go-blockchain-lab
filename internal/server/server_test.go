package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/carl-dourado/go-blockchain-lab/internal/blockchain"
)

func TestAPICreateMineAndValidate(t *testing.T) {
	handler := NewHandler(blockchain.NewStore(t.TempDir()))

	createBody := bytes.NewBufferString(`{"author":"carlos","data":"api record"}`)
	createRequest := httptest.NewRequest(http.MethodPost, "/records", createBody)
	createRequest.Header.Set("Content-Type", "application/json")
	createResponse := httptest.NewRecorder()
	handler.ServeHTTP(createResponse, createRequest)

	if createResponse.Code != http.StatusCreated {
		t.Fatalf("expected create status 201, got %d with body %s", createResponse.Code, createResponse.Body.String())
	}

	mineRequest := httptest.NewRequest(http.MethodPost, "/mine", nil)
	mineResponse := httptest.NewRecorder()
	handler.ServeHTTP(mineResponse, mineRequest)

	if mineResponse.Code != http.StatusCreated {
		t.Fatalf("expected mine status 201, got %d with body %s", mineResponse.Code, mineResponse.Body.String())
	}

	validateRequest := httptest.NewRequest(http.MethodGet, "/validate", nil)
	validateResponse := httptest.NewRecorder()
	handler.ServeHTTP(validateResponse, validateRequest)

	if validateResponse.Code != http.StatusOK {
		t.Fatalf("expected validate status 200, got %d", validateResponse.Code)
	}

	var payload struct {
		Valid bool `json:"valid"`
	}
	if err := json.NewDecoder(validateResponse.Body).Decode(&payload); err != nil {
		t.Fatalf("decode validate response failed: %v", err)
	}

	if !payload.Valid {
		t.Fatal("expected chain to validate")
	}
}

func TestAPIMineRejectsEmptyPendingQueue(t *testing.T) {
	handler := NewHandler(blockchain.NewStore(t.TempDir()))
	request := httptest.NewRequest(http.MethodPost, "/mine", nil)
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	if response.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", response.Code)
	}
}

func TestExplorerPageServed(t *testing.T) {
	handler := NewHandler(blockchain.NewStore(t.TempDir()))
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", response.Code)
	}

	contentType := response.Header().Get("Content-Type")
	if !strings.Contains(contentType, "text/html") {
		t.Fatalf("expected html content type, got %q", contentType)
	}

	if !strings.Contains(response.Body.String(), "Go Blockchain Lab") {
		t.Fatal("expected explorer page to contain app title")
	}

	if !strings.Contains(response.Body.String(), "chain-canvas") {
		t.Fatal("expected explorer page to contain chain canvas")
	}
}
