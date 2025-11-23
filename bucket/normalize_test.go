package bucket

import "testing"

func TestNormalizeName(t *testing.T) {
	tests := []struct {
		input       string
		expected    string
		expectError bool
	}{
		{"simple-bucket", "simple-bucket", false},
		{"SIMPLE-BUCKET", "simple-bucket", false},
		{"https://example-bucket.s3.amazonaws.com", "example-bucket", false},
		{"https://example-bucket.s3.us-east-1.amazonaws.com/path/file.txt", "example-bucket", false},
		{"https://s3.amazonaws.com/example-bucket/foo", "example-bucket", false},
		{"https://s3.us-west-2.amazonaws.com/example-bucket/foo", "example-bucket", false},
		{"s3://mixed-case-bucket/path", "mixed-case-bucket", false},
		{"example-bucket.s3.amazonaws.com", "example-bucket", false},
		{"https://example-bucket.s3-website-us-east-1.amazonaws.com", "example-bucket", false},
		{"", "", true},
		{"https://example.com/not-an-s3-url", "", true},
	}

	for _, test := range tests {
		got, err := NormalizeName(test.input)
		if test.expectError {
			if err == nil {
				t.Errorf("NormalizeName(%q) expected error, got nil", test.input)
			}
			continue
		}
		if err != nil {
			t.Errorf("NormalizeName(%q) unexpected error: %v", test.input, err)
			continue
		}
		if got != test.expected {
			t.Errorf("NormalizeName(%q) = %q, expected %q", test.input, got, test.expected)
		}
	}
}
