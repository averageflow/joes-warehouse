package app

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/averageflow/joes-warehouse/internal/infrastructure"
)

func TestApplicationServer_getTransactionsHandlerIsUnauthorizedWithoutToken(t *testing.T) {
	t.Parallel()

	s := NewApplicationServer(&ApplicationState{
		DB: infrastructure.MockApplicationDatabase{},
	})

	ts := httptest.NewServer(s.State.Handler)
	defer ts.Close()

	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodGet,
		fmt.Sprintf("%s/%s", ts.URL, "api/transactions"),
		nil,
	)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusUnauthorized {
		t.Fatalf("Expected status code %d, got %d", http.StatusUnauthorized, resp.StatusCode)
	}
}

func TestApplicationServer_getTransactionsHandlerIsSuccessful(t *testing.T) {
	t.Parallel()

	s := NewApplicationServer(&ApplicationState{
		DB: infrastructure.MockApplicationDatabase{},
	})

	ts := httptest.NewServer(s.State.Handler)
	defer ts.Close()

	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodGet,
		fmt.Sprintf("%s/%s", ts.URL, "api/transactions"),
		nil,
	)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", dangerouslyHardcodedAuthToken))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}
}
