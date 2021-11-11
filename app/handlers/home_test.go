package handlers

import (
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHome(t *testing.T) {
	t.Run("return application name", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()

		Home(rr, req)

		require.Contains(t, rr.Body.String(), "argocd-playground-app")
	})
}
