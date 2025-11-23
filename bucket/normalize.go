package bucket

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
)

// NormalizeName accepts a raw bucket identifier (bucket name, URL, or ARN-like host) and returns a canonical bucket name.
func NormalizeName(raw string) (string, error) {
	trimmed := strings.TrimSpace(raw)
	if trimmed == "" {
		return "", errors.New("empty bucket identifier")
	}

	if extracted, ok := extractNameFromURL(trimmed); ok {
		extracted = strings.ToLower(extracted)
		if IsValidS3BucketName(extracted) {
			return extracted, nil
		}
	}

	candidate := strings.ToLower(trimmed)
	if IsValidS3BucketName(candidate) {
		return candidate, nil
	}

	return "", fmt.Errorf("unable to parse bucket name from %q", raw)
}

func extractNameFromURL(identifier string) (string, bool) {
	candidates := []string{identifier}
	if !strings.Contains(identifier, "://") {
		candidates = append(candidates, "https://"+identifier)
	}

	for _, candidate := range candidates {
		u, err := url.Parse(candidate)
		if err != nil {
			continue
		}

		// s3://bucket-name/object
		if u.Scheme == "s3" && u.Host != "" {
			return strings.ToLower(u.Host), true
		}

		host := strings.ToLower(u.Host)
		if host == "" {
			continue
		}
		host = trimPort(host)

		if name := virtualHostBucket(host); name != "" {
			return name, true
		}

		if pathBucket := firstPathSegment(u.Path); pathBucket != "" && strings.HasPrefix(host, "s3") {
			return strings.ToLower(pathBucket), true
		}
	}
	return "", false
}

func virtualHostBucket(host string) string {
	markers := []string{".s3.", ".s3-", ".s3.amazonaws.com", ".s3-website", ".s3.dualstack."}
	for _, marker := range markers {
		if idx := strings.Index(host, marker); idx > 0 {
			return host[:idx]
		}
	}
	return ""
}

func firstPathSegment(path string) string {
	trimmed := strings.Trim(path, "/")
	if trimmed == "" {
		return ""
	}
	parts := strings.SplitN(trimmed, "/", 2)
	return parts[0]
}

func trimPort(host string) string {
	if h, _, found := strings.Cut(host, ":"); found {
		return h
	}
	return host
}
