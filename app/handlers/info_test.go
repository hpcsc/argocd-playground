package handlers

import (
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestInfoHandler_ServeHTTP(t *testing.T) {
	t.Run("return error when version file not found", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/info", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()

		filePath := "not-existing"
		handler := NewInfoHandler(filePath)
		handler.ServeHTTP(rr, req)

		require.Equal(t, http.StatusInternalServerError, rr.Code)
		require.Contains(t, rr.Body.String(), "failed to read version")
	})

	t.Run("return error when version file cannot be unmarshalled", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/info", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()

		filePath := "testdata/invalid-version.json"
		handler := NewInfoHandler(filePath)
		handler.ServeHTTP(rr, req)

		require.Equal(t, http.StatusInternalServerError, rr.Code)
		require.Contains(t, rr.Body.String(), "invalid file format")
	})

	t.Run("return version and commit when version file is valid", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/info", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()

		filePath := "testdata/valid-version.json"
		handler := NewInfoHandler(filePath)
		handler.ServeHTTP(rr, req)

		require.Equal(t, http.StatusOK, rr.Code)
		require.Contains(t, rr.Body.String(), "v1.2.3")
		require.Contains(t, rr.Body.String(), "abcd1234")
	})

	t.Run("return current environment from environment variable", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/info", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		os.Setenv("ENVIRONMENT", "test-env")
		defer os.Unsetenv("ENVIRONMENT")

		filePath := "testdata/valid-version.json"
		handler := NewInfoHandler(filePath)
		handler.ServeHTTP(rr, req)

		require.Equal(t, http.StatusOK, rr.Code)
		require.Contains(t, rr.Body.String(), "test-env")
	})
}
