package service_auth

import (
	"context"
	"strings"
)

type requestMetadataKey struct{}

type requestMetadata struct {
	ip        string
	userAgent string
}

func withRequestMetadata(ctx context.Context, ip string, userAgent string) context.Context {
	return context.WithValue(ctx, requestMetadataKey{}, requestMetadata{
		ip:        strings.TrimSpace(ip),
		userAgent: strings.TrimSpace(userAgent),
	})
}

func WithRequestMetadataForHandler(ctx context.Context, ip string, userAgent string) context.Context {
	return withRequestMetadata(ctx, ip, userAgent)
}

func requestIPFromContext(ctx context.Context) string {
	meta, ok := ctx.Value(requestMetadataKey{}).(requestMetadata)
	if !ok || meta.ip == "" {
		return "unknown"
	}
	return meta.ip
}

func requestUserAgentFromContext(ctx context.Context) string {
	meta, ok := ctx.Value(requestMetadataKey{}).(requestMetadata)
	if !ok || meta.userAgent == "" {
		return "unknown"
	}
	return meta.userAgent
}
