package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	var timeout time.Duration
	flag.DurationVar(&timeout, "timeout", 10*time.Second, "connection timeout")
	flag.Parse()

	args := flag.Args()
	if len(args) < 2 || len(args) > 3 {
		//if len(args) != 2 {
		return errors.New("count of args should be two or three")
	}

	address := net.JoinHostPort(args[0], args[1])

	client := NewTelnetClient(address, timeout, os.Stdin, os.Stdout)
	err := client.Connect()
	if err != nil {
		return err
	}
	fmt.Fprintf(os.Stderr, "...Connected to %s\n", address)
	defer client.Close()

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	go func() {
		if err := client.Send(); err != nil {
			client.Print(os.Stdout, err)
		}
		client.Print(os.Stderr, "...EOF")
		cancel()
	}()

	go func() {
		if err := client.Receive(); err != nil {
			client.Print(os.Stdout, err)
		}
		client.Print(os.Stderr, "...Connection was closet by peer")
		cancel()
	}()

	<-ctx.Done()
	return nil
}
