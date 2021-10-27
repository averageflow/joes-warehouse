package app

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/averageflow/joes-warehouse/internal/domain/articles"
	"github.com/averageflow/joes-warehouse/internal/infrastructure"
	"github.com/brianvoe/gofakeit/v6"
)

func TestApplicationServer_getArticlesHandlerIsUnauthorizedWithoutToken(t *testing.T) {
	t.Parallel()

	s := NewApplicationServer(&ApplicationState{
		DB: infrastructure.MockApplicationDatabase{},
	})

	ts := httptest.NewServer(s.State.Handler)
	defer ts.Close()

	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodGet,
		fmt.Sprintf("%s/%s", ts.URL, "api/articles"),
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

func TestApplicationServer_getArticlesHandlerIsSuccessful(t *testing.T) {
	t.Parallel()

	s := NewApplicationServer(&ApplicationState{
		DB: infrastructure.MockApplicationDatabase{},
	})

	ts := httptest.NewServer(s.State.Handler)
	defer ts.Close()

	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodGet,
		fmt.Sprintf("%s/%s", ts.URL, "api/articles"),
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

func TestApplicationServer_addArticlesHandlerIsUnauthorizedWithoutToken(t *testing.T) {
	t.Parallel()

	s := NewApplicationServer(&ApplicationState{
		DB: infrastructure.MockApplicationDatabase{},
	})

	ts := httptest.NewServer(s.State.Handler)
	defer ts.Close()

	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodPost,
		fmt.Sprintf("%s/%s", ts.URL, "api/articles"),
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

func TestApplicationServer_addArticlesHandlerErrorsOnBadData(t *testing.T) {
	t.Parallel()

	s := NewApplicationServer(&ApplicationState{
		DB: infrastructure.MockApplicationDatabase{},
	})

	ts := httptest.NewServer(s.State.Handler)
	defer ts.Close()

	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodPost,
		fmt.Sprintf("%s/%s", ts.URL, "api/articles"),
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

	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("Expected status code %d, got %d", http.StatusBadRequest, resp.StatusCode)
	}
}

func TestApplicationServer_addArticlesHandlerIsSuccessful(t *testing.T) {
	t.Parallel()

	s := NewApplicationServer(&ApplicationState{
		DB: infrastructure.MockApplicationDatabase{},
	})

	ts := httptest.NewServer(s.State.Handler)
	defer ts.Close()

	var fakeData articles.RawArticleUploadRequest

	_ = gofakeit.Struct(&fakeData)

	marshaled, err := json.Marshal(fakeData)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodPost,
		fmt.Sprintf("%s/%s", ts.URL, "api/articles"),
		bytes.NewBuffer(marshaled),
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

func TestApplicationServer_deleteArticlesHandlerIsUnauthorizedWithoutToken(t *testing.T) {
	t.Parallel()

	s := NewApplicationServer(&ApplicationState{
		DB: infrastructure.MockApplicationDatabase{},
	})

	ts := httptest.NewServer(s.State.Handler)
	defer ts.Close()

	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodDelete,
		fmt.Sprintf("%s/%s", ts.URL, "api/articles/1"),
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

func TestApplicationServer_deleteArticlesHandlerErrorsOnBadData(t *testing.T) {
	t.Parallel()

	s := NewApplicationServer(&ApplicationState{
		DB: infrastructure.MockApplicationDatabase{},
	})

	ts := httptest.NewServer(s.State.Handler)
	defer ts.Close()

	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodDelete,
		fmt.Sprintf("%s/%s", ts.URL, "api/articles/a"),
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

	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("Expected status code %d, got %d", http.StatusBadRequest, resp.StatusCode)
	}
}

func TestApplicationServer_deleteArticlesHandlerIsSuccessful(t *testing.T) {
	t.Parallel()

	s := NewApplicationServer(&ApplicationState{
		DB: infrastructure.MockApplicationDatabase{},
	})

	ts := httptest.NewServer(s.State.Handler)
	defer ts.Close()

	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodDelete,
		fmt.Sprintf("%s/%s", ts.URL, "api/articles/1"),
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
