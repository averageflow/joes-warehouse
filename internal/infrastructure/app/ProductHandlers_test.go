package app

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/averageflow/joes-warehouse/internal/domain/products"
	"github.com/averageflow/joes-warehouse/internal/infrastructure"
	"github.com/brianvoe/gofakeit/v6"
)

func TestApplicationServer_getProductsHandlerIsUnauthorizedWithoutToken(t *testing.T) {
	t.Parallel()

	s := NewApplicationServer(&ApplicationState{
		DB: infrastructure.MockApplicationDatabase{},
	})

	ts := httptest.NewServer(s.State.Handler)
	defer ts.Close()

	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodGet,
		fmt.Sprintf("%s/%s", ts.URL, "api/products"),
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

func TestApplicationServer_getProductsHandlerIsSuccessful(t *testing.T) {
	t.Parallel()

	s := NewApplicationServer(&ApplicationState{
		DB: infrastructure.MockApplicationDatabase{},
	})

	ts := httptest.NewServer(s.State.Handler)
	defer ts.Close()

	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodGet,
		fmt.Sprintf("%s/%s", ts.URL, "api/products"),
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

func TestApplicationServer_addProductsHandlerIsUnauthorizedWithoutToken(t *testing.T) {
	t.Parallel()

	s := NewApplicationServer(&ApplicationState{
		DB: infrastructure.MockApplicationDatabase{},
	})

	ts := httptest.NewServer(s.State.Handler)
	defer ts.Close()

	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodPost,
		fmt.Sprintf("%s/%s", ts.URL, "api/products"),
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

func TestApplicationServer_addProductsHandlerErrorsOnBadData(t *testing.T) {
	t.Parallel()

	s := NewApplicationServer(&ApplicationState{
		DB: infrastructure.MockApplicationDatabase{},
	})

	ts := httptest.NewServer(s.State.Handler)
	defer ts.Close()

	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodPost,
		fmt.Sprintf("%s/%s", ts.URL, "api/products"),
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

func TestApplicationServer_addProductsHandlerIsSuccessful(t *testing.T) {
	t.Parallel()

	s := NewApplicationServer(&ApplicationState{
		DB: infrastructure.MockApplicationDatabase{},
	})

	ts := httptest.NewServer(s.State.Handler)
	defer ts.Close()

	var fakeData products.RawProductUploadRequest

	_ = gofakeit.Struct(&fakeData)

	marshaled, err := json.Marshal(fakeData)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodPost,
		fmt.Sprintf("%s/%s", ts.URL, "api/products"),
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

func TestApplicationServer_deleteProductsHandlerIsUnauthorizedWithoutToken(t *testing.T) {
	t.Parallel()

	s := NewApplicationServer(&ApplicationState{
		DB: infrastructure.MockApplicationDatabase{},
	})

	ts := httptest.NewServer(s.State.Handler)
	defer ts.Close()

	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodDelete,
		fmt.Sprintf("%s/%s", ts.URL, "api/products/1"),
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

func TestApplicationServer_deleteProductsHandlerErrorsOnBadData(t *testing.T) {
	t.Parallel()

	s := NewApplicationServer(&ApplicationState{
		DB: infrastructure.MockApplicationDatabase{},
	})

	ts := httptest.NewServer(s.State.Handler)
	defer ts.Close()

	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodDelete,
		fmt.Sprintf("%s/%s", ts.URL, "api/products/a"),
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

func TestApplicationServer_deleteProductsHandlerIsSuccessful(t *testing.T) {
	t.Parallel()

	s := NewApplicationServer(&ApplicationState{
		DB: infrastructure.MockApplicationDatabase{},
	})

	ts := httptest.NewServer(s.State.Handler)
	defer ts.Close()

	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodDelete,
		fmt.Sprintf("%s/%s", ts.URL, "api/products/1"),
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

func TestApplicationServer_sellProductsHandlerIsUnauthorizedWithoutToken(t *testing.T) {
	t.Parallel()

	s := NewApplicationServer(&ApplicationState{
		DB: infrastructure.MockApplicationDatabase{},
	})

	ts := httptest.NewServer(s.State.Handler)
	defer ts.Close()

	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodPatch,
		fmt.Sprintf("%s/%s", ts.URL, "api/products/sell"),
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

func TestApplicationServer_sellProductsHandlerErrorsOnBadData(t *testing.T) {
	t.Parallel()

	s := NewApplicationServer(&ApplicationState{
		DB: infrastructure.MockApplicationDatabase{},
	})

	ts := httptest.NewServer(s.State.Handler)
	defer ts.Close()

	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodPatch,
		fmt.Sprintf("%s/%s", ts.URL, "api/products/sell"),
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

func TestApplicationServer_sellProductsHandlerErrorsDueToNoProduct(t *testing.T) {
	t.Parallel()

	s := NewApplicationServer(&ApplicationState{
		DB: infrastructure.MockApplicationDatabase{},
	})

	ts := httptest.NewServer(s.State.Handler)
	defer ts.Close()

	fakeData := products.SellProductRequest{
		Data: []products.SellProductRequestItems{
			{
				ProductID: gofakeit.Int64(),
				Amount:    gofakeit.Int64(),
			},
		},
	}

	marshaled, err := json.Marshal(fakeData)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodPatch,
		fmt.Sprintf("%s/%s", ts.URL, "api/products/sell"),
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

	if resp.StatusCode != http.StatusUnprocessableEntity {
		t.Fatalf("Expected status code %d, got %d", http.StatusUnprocessableEntity, resp.StatusCode)
	}
}
