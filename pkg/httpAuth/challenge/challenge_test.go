package challenge

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/spf13/viper"
)

func TestRefreshChallenge_PathTraversal(t *testing.T) {
	// Create a mock ConsoleMe server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the request path is what we expect
		if r.URL.Path != "/noauth/v1/challenge_generator/test" {
			t.Errorf("unexpected request path: got %v want %v",
				r.URL.Path, "/noauth/v1/challenge_generator/test")
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, `{"challenge_url": "http://localhost", "polling_url": "http://localhost"}`)
	}))
	defer server.Close()

	// Set the viper config
	viper.Set("consoleme_url", server.URL)
	viper.Set("challenge_settings.user", "test/../../../../..")

	// Call the function
	err := RefreshChallenge()
	if err != nil {
		// We expect an error because the polling will fail, but we are only interested in the request path.
		// A more sophisticated test would mock the polling as well.
	}
}
