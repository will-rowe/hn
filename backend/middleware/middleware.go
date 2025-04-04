package middleware

import (
	"context"
	"net/http"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// Dummy validation for bearer token
func validateToken(token string) bool {
	return token == "testtoken"
}

// HTTP Middleware for REST API
func AuthHTTPMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") || !validateToken(strings.TrimPrefix(authHeader, "Bearer ")) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// gRPC Unary Interceptor
func AuthUnaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	if md, ok := grpcHeader(ctx); ok {
		token := extractBearerToken(md["authorization"])
		if validateToken(token) {
			return handler(ctx, req)
		}
	}
	return nil, status.Error(codes.Unauthenticated, "invalid or missing token")
}

// helper to extract headers from context
func grpcHeader(ctx context.Context) (map[string][]string, bool) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, false
	}
	return md, true
}

func extractBearerToken(headers []string) string {
	for _, h := range headers {
		if strings.HasPrefix(h, "Bearer ") {
			return strings.TrimPrefix(h, "Bearer ")
		}
	}
	return ""
}
