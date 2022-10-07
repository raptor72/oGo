package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

type Telnet struct {
	address string
	timeout time.Duration
	in      io.ReadCloser
	out     io.Writer
	conn    net.Conn
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) *Telnet {
	return &Telnet{
		address: address,
		timeout: timeout,
		in:      in,
		out:     out,
	}
}

func (c *Telnet) Connect() error {
	conn, err := net.DialTimeout("tcp", c.address, c.timeout)
	if err != nil {
		return err
	}
	c.conn = conn
	return nil
}

func (c *Telnet) Close() error {
	return c.conn.Close()
}

func (c *Telnet) Send() error {
	_, err := io.Copy(c.conn, c.in)
	return err
}

func (c *Telnet) Receive() error {
	_, err := io.Copy(c.out, c.conn)
	return err
}

func (c *Telnet) Print(w io.Writer, s interface{}) {
	fmt.Fprintln(w, s)
}
