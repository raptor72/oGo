package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
	ErrDadSourceFile         = errors.New("source file is not exists")
	ErrBadDestination        = errors.New("bad destination to copy")
)

type filePath struct {
	absPath  string
	filename string
	size     int64
}

func checkCondition(fromPath, toPath string, offset int64) (filePath, filePath, error) {
	from, to := filePath{}, filePath{}
	if srcInfo, err := os.Stat(fromPath); err != nil {
		return from, to, err
	} else {
		if !srcInfo.Mode().IsRegular() {
			return from, to, ErrUnsupportedFile
		}
		if srcInfo.IsDir() {
			return from, to, ErrDadSourceFile
		}
		if offset > srcInfo.Size() {
			return from, to, ErrOffsetExceedsFileSize
		}
		from.absPath, err = filepath.Abs(path.Dir(fromPath))
		if err != nil {
			return from, to, err
		}
		from.filename = srcInfo.Name()
		from.size = srcInfo.Size()
	}

	to.absPath = filepath.Dir(toPath)
	to.filename = filepath.Base(toPath)

	if filepath.Base(toPath) == "." || strings.HasSuffix(toPath, "/") {
		to.filename = from.filename
	}
	return from, to, nil
}

func Copy(fromPath, toPath string, offset, limit int64) error {
	from, to, err := checkCondition(fromPath, toPath, offset)
	if err != nil {
		return err
	}

	srcReader, err := os.Open(path.Join(from.absPath, from.filename))
	if err != nil {
		return err
	}

	if offset != 0 {
		_, err := srcReader.Seek(offset, os.SEEK_SET)
		if err != nil {
			return err
		}
	}

	dstPath := filepath.Join(to.absPath, to.filename)
	fmt.Println("dstpath!: ", dstPath)
	_, err = os.Create(dstPath)
	if err != nil {
		err = fmt.Errorf("%s. Detail: %w, Third error", ErrBadDestination, err)
		return err
	}

	dstWriter, err := os.OpenFile(dstPath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}

	if limit+offset > from.size {
		limit = from.size - offset
	}

	if limit > from.size || limit == 0 {
		limit = from.size
	}

	byteLen, err := io.CopyN(dstWriter, srcReader, limit)
	if err != nil {
		fmt.Println("io.CopyN error")
		return err
	}
	fmt.Printf("Copy %v bytes\n", byteLen)

	return nil
}
