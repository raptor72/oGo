package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOkRunCmd(t *testing.T) {
	expectedContent := "Bar"
	env := EnvValue{Value: expectedContent, NeedRemove: false}
	expected := make(Environment)
	os.Unsetenv("arg1")
	expected["arg1"] = env
	code := RunCmd([]string{"echo", ""}, expected)
	require.Equal(t, os.Getenv("arg1"), expectedContent)
	require.Equal(t, 0, code)
}

func TestPurgeOldValue(t *testing.T) {
	expectedContent := ""
	env := EnvValue{Value: expectedContent, NeedRemove: true}
	expected := make(Environment)
	expected["arg1"] = env
	code := RunCmd([]string{"echo"}, expected)
	require.Equal(t, os.Getenv("arg1"), expectedContent)
	require.Equal(t, 0, code)
}
