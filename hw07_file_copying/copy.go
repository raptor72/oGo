package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
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
	if _, err := os.Stat(toPath); errors.Is(err, os.ErrNotExist) {
		return from, to, ErrBadDestination
	} else {
		to.absPath, err = filepath.Abs(path.Dir(toPath))
		if err != nil {
			return from, to, err
		}
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
		fmt.Println("srcReader err")
		return err
	}

	//if offset != 0 {
	//	_, err = srcReader.Seek(offset, 0)
	//	if err != nil {
	//		return err
	//	}
	//}

	dstPath := path.Join(to.absPath, from.filename)
	newFile, err := os.Create(dstPath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(newFile.Name())

	dstWriter, err := os.OpenFile(newFile.Name(), os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("dstWriter error")
		return err
	}

	if limit == 0 || limit > from.size {
		limit = from.size
	}

	byteLen, err := io.CopyN(dstWriter, srcReader, limit)
	if err != nil {
		fmt.Println("io.CopyN error")
		return err
	}
	fmt.Printf("Copy %v bytes\n", byteLen)

	//fmt.Println(from.absPath)
	//fmt.Println(from.filename)
	//fmt.Println(to.absPath)
	//fmt.Println(to.filename)

	return nil
}
