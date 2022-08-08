package main

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadEmptyDir(t *testing.T) {
	expected := make(Environment)
	dir, err := ioutil.TempDir(".", "tempdir")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(dir)
	result, err := ReadDir(dir)
	require.NoError(t, err)
	require.Equal(t, expected, result)
}
