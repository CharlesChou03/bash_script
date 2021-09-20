package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	r := m.Run()

	if r == 0 && testing.CoverMode() != "" {
		c := testing.Coverage() * 100
		l := 70.00
		fmt.Println("=================================================")
		fmt.Println("||               Coverage Report               ||")
		fmt.Println("=================================================")
		fmt.Printf("Cover mode: %s\n", testing.CoverMode())
		fmt.Printf("Coverage  : %.2f %% (Threshold: %.2f %%)\n\n", c, l)
		if c < l {
			fmt.Println("[Tests passed but coverage failed]")
			r = -1
		}
	}

	os.Exit(r)
}

func TestHealth(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestVersion(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/version", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
