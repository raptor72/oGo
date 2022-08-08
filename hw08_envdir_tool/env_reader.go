package main

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	envArray := make(Environment)

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if !file.Mode().IsRegular() {
			continue
		}

		var envVar EnvValue

		fileReader, err := os.Open(dir + "/" + file.Name())
		if err != nil {
			return nil, err
		}
		r := bufio.NewReader(fileReader)
		line, _, err := r.ReadLine()
		if err != nil {
			if errors.Is(err, io.EOF) {
				// Файл пустой, переменная окружения должна быть удалены
				envVar.NeedRemove = true
			} else {
				return nil, nil
			}
		}
		// Удаляем терминальные нули
		line = bytes.ReplaceAll(line, []byte("\000"), []byte("\n"))
		// Удаляем пробелы табы и знак равно в конце строки
		value := strings.TrimRight(string(line), "\t =")

		if len(value) == 0 {
			envVar.NeedRemove = true
		}
		envVar.Value = value
		envArray[file.Name()] = envVar
	}
	return envArray, nil
}
