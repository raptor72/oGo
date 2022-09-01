package hw10programoptimization

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/mailru/easyjson"
)

type User struct {
	Email string
}

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	dmn := strings.Join([]string{".", domain}, "")
	u, err := getUsers(r, dmn)
	if err != nil {
		return nil, fmt.Errorf("get users error: %w", err)
	}
	return *u, nil
}

func getUsers(r io.Reader, domain string) (*DomainStat, error) {
	mapCounter := make(DomainStat)

	err := readByChunks(r, &mapCounter, domain)
	if err != nil {
		return nil, err
	}

	return &mapCounter, nil
}

func processByteLine(line []byte, mapCounter DomainStat, domain string) error {
	var user User

	if err := easyjson.Unmarshal(line, &user); err != nil {
		return err
	}

	if strings.HasSuffix(user.Email, domain) {
		_, mail, _ := strings.Cut(user.Email, "@")
		mail = strings.ToLower(mail)
		mapCounter[mail]++
	}
	return nil
}

func readByChunks(r io.Reader, mapCounter *DomainStat, domain string) error {
	reader := bufio.NewReader(r)

	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			if !errors.Is(err, io.EOF) {
				return err
			}
		}

		err = processByteLine(line, *mapCounter, domain)

		if err != nil {
			if errors.Is(err, io.EOF) {
				return nil
			}
		}
	}
}
