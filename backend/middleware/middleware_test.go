package middleware_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/will-rowe/hn/backend/middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func TestAuthHTTPMiddleware(t *testing.T) {
	// define a test handler that will only run if middleware passes
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	tests := []struct {
		name           string
		authHeader     string
		expectedStatus int
	}{
		{"Valid Token", "Bearer testtoken", http.StatusOK},
		{"Invalid Token", "Bearer wrongtoken", http.StatusUnauthorized},
		{"Missing Header", "", http.StatusUnauthorized},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/", nil)
			if tt.authHeader != "" {
				req.Header.Set("Authorization", tt.authHeader)
			}
			rr := httptest.NewRecorder()

			handler := middleware.AuthHTTPMiddleware(testHandler)
			handler.ServeHTTP(rr, req)

			if rr.Code != tt.expectedStatus {
				t.Errorf("expected %d, got %d", tt.expectedStatus, rr.Code)
			}
		})
	}
}

func TestAuthUnaryInterceptor(t *testing.T) {
	tests := []struct {
		name         string
		md           metadata.MD
		expectedCode string
		expectError  bool
	}{
		{"Valid Token", metadata.Pairs("authorization", "Bearer testtoken"), "OK", false},
		{"Invalid Token", metadata.Pairs("authorization", "Bearer badtoken"), "Unauthenticated", true},
		{"No Token", metadata.Pairs(), "Unauthenticated", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := metadata.NewIncomingContext(context.Background(), tt.md)

			handler := func(ctx context.Context, req interface{}) (interface{}, error) {
				return "OK", nil
			}

			resp, err := middleware.AuthUnaryInterceptor(ctx, nil, &grpc.UnaryServerInfo{FullMethod: "/test"}, handler)

			if tt.expectError && err == nil {
				t.Errorf("expected error, got nil")
			}

			if !tt.expectError && err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if !tt.expectError && resp != "OK" {
				t.Errorf("expected response 'OK', got %v", resp)
			}
		})
	}
}
