package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopyIrregularFile(t *testing.T) {
	tests := []struct {
		from string
	}{
		{"/dev/urandom"},
		{"/dev/null"},
		{"/dev/tty"},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.from, func(t *testing.T) {
			err := Copy(tc.from, "/tmp/out.txt", 0, 0)
			require.Truef(t, errors.Is(err, ErrUnsupportedFile), "actual error %q", err)
		})
	}
}

func TestOffsetExceedFileSize(t *testing.T) {
	tests := []struct {
		str string
		i64 int64
	}{
		{"6618", 6618},
		{"10000", 10000},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.str, func(t *testing.T) {
			err := Copy("testdata/input.txt", "/tmp/out.txt", tc.i64, 0)
			require.Truef(t, errors.Is(err, ErrOffsetExceedsFileSize), "actual error %q", err)
		})
	}
}

func TestBadSourceFile(t *testing.T) {
	tests := []struct {
		from string
	}{
		{"/tmp/unexpected.file"},
		{"/root/root"},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.from, func(t *testing.T) {
			err := Copy(tc.from, "/tmp/out.txt", 0, 0)
			require.ErrorContains(t, err, ErrBadSourceFile.Error())
		})
	}
}

func TestBadDestination(t *testing.T) {
	tests := []struct {
		to string
	}{
		{"/dir/dir/file"},
		{"/dir/dir/"},
		{"./dir/"},
		{"./dir/file"},
		{"/root/"},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.to, func(t *testing.T) {
			err := Copy("testdata/input.txt", tc.to, 0, 0)
			require.ErrorContains(t, err, ErrBadDestination.Error())
		})
	}
}
