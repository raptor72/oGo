package main

import (
	"errors"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCopyIrregularFile(t *testing.T) {
	tests := []struct {
		from string
	}{
		{"/dev/urandom"},
		{"/dev/null"},
	}
	//for _, tc := range tests {
	//	err := Copy(tc.from, "/tmp", 0, 0)
	//	//fmt.Println(err)
	//	require.Truef(t, errors.Is(err, ErrUnsupportedFile), "actual error %q", err)
	//}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.from, func(t *testing.T) {
			err := Copy(tc.from, "/tmp", 0, 0)
			require.Truef(t, errors.Is(err, ErrUnsupportedFile), "actual error %q", err)
		})
	}
}
