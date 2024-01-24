package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"testing-gin-api/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetAccountsRoute(t *testing.T) {
	r := gin.Default()
	SetupAPI(r)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/accounts", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	accounts := mocks.GetAccounts()
	response := QueryAccountsResponse{
		Accounts: accounts,
	}
	expected, _ := json.Marshal(response)

	assert.Equal(t, string(expected), w.Body.String())
}

func TestGetSortedAccountsRoute(t *testing.T) {
	r := gin.Default()
	SetupAPI(r)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/accounts?sort=address", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	accounts := mocks.GetSortedAccounts()
	response := QueryAccountsResponse{
		Accounts: accounts,
		Sort:     "address",
	}
	expected, _ := json.Marshal(response)
	assert.Equal(t, string(expected), w.Body.String())
}
