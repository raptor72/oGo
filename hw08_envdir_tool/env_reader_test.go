package main

import (
	"bufio"
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

func TestReadOk(t *testing.T) {
	expectedContent := "Bar"
	env := EnvValue{Value: expectedContent, NeedRemove: false}

	dir, err := ioutil.TempDir(".", "tempdir")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(dir)
	tmpOut, err := os.CreateTemp(dir, "out")
	if err != nil {
		log.Println(err)
	}
	expected := make(Environment)

	tmpOutInfo, err := os.Stat(tmpOut.Name())
	if err != nil {
		log.Println(err)
	}
	expected[tmpOutInfo.Name()] = env
	os.Unsetenv(tmpOutInfo.Name())
	datawriter := bufio.NewWriter(tmpOut)
	datawriter.WriteString(expectedContent)
	datawriter.Flush()
	result, err := ReadDir(dir)
	require.NoError(t, err)
	require.Equal(t, expected, result)
}

func TestRemoveEmpty(t *testing.T) {
	expectedContent := ""
	env := EnvValue{Value: expectedContent, NeedRemove: true}

	dir, err := ioutil.TempDir(".", "tempdir")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(dir)
	tmpOut, err := os.CreateTemp(dir, "out")
	if err != nil {
		log.Println(err)
	}

	expected := make(Environment)

	tmpOutInfo, err := os.Stat(tmpOut.Name())
	if err != nil {
		log.Println(err)
	}
	expected[tmpOutInfo.Name()] = env
	os.Unsetenv(tmpOutInfo.Name())
	datawriter := bufio.NewWriter(tmpOut)
	datawriter.WriteString(expectedContent)
	datawriter.Flush()
	result, err := ReadDir(dir)
	require.NoError(t, err)
	require.Equal(t, expected, result)
}
