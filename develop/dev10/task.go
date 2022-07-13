package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Client struct {
	conn    net.Conn
	ctx     context.Context
	cancel  context.CancelFunc
	signal  chan os.Signal
	data    chan string
	stop    chan int
	addr    string
	timeout time.Duration
}

func GetClient(addr string, timeout time.Duration) Client {
	c := Client{addr: addr, timeout: timeout, signal: make(chan os.Signal, 1), data: make(chan string), stop: make(chan int)}
	c.ctx, c.cancel = context.WithCancel(context.Background())
	return c
}

func (c *Client) Connect() error {
	var err error
	dialer := net.Dialer{Timeout: c.timeout}
	c.conn, err = dialer.Dial("tcp", c.addr)
	fmt.Println(c.addr, err)
	if err != nil {
		return err
	}
	fmt.Println("Successfully connected to:", c.addr)
	return nil
}

func (c *Client) catchSignal() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	<-ch
	fmt.Println("Получили сигнал остановки")
	c.stop <- 0
}

func (c *Client) readData() {
	for {
		select {
		case <-c.ctx.Done():
			return
		default:
			if err := c.conn.SetReadDeadline(time.Now().Add(time.Millisecond * 10)); err != nil {
				log.Fatal(err)
			}
			data, err := ioutil.ReadAll(c.conn)
			if err != nil {
				if err == io.EOF {
					c.stop <- 0
					log.Fatal(err)
					return
				}
				if netErr, ok := err.(net.Error); ok && !netErr.Timeout() {
					c.stop <- 0
					log.Fatal(err)
					return
				}
			}
			if len(data) == 0 {
				break
			}
			fmt.Print("New message:", string(data))
		}
	}
}

func (c *Client) writeData() {
	for {
		select {
		case <-c.ctx.Done():
			break
		default:
			reader := bufio.NewReader(os.Stdin)
			newData, err := reader.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					c.stop <- 0
					log.Fatal(err)
					return
				}
				c.stop <- 0
				log.Fatal(err)
				return
			}
			if _, err := c.conn.Write([]byte(newData)); err != nil {
				log.Println(err)
				return
			}
		}
	}
}

func (c *Client) Run() {
	err := c.Connect()
	if err != nil {
		log.Fatal("error")
		return
	}
	go c.catchSignal()
	go c.readData()
	go c.writeData()
	<-c.stop
	c.cancel()
	err = c.conn.Close()
	if err != nil {
		return
	}
	return
}

func FlagParse() (string, string) {
	timeout := flag.String("timeout", "10s", "timeout for connect")
	flag.Parse()
	if len(flag.Args()) < 2 {
		fmt.Println("Please enter 2 args")
		return "", ""
	}
	addr := flag.Arg(0) + ":" + flag.Arg(1)
	return *timeout, addr
}

func main() {
	timeInput, addr := FlagParse()
	if timeInput == "" || addr == "" {
		return
	}
	fmt.Println(timeInput, addr)
	timeout, err := time.ParseDuration(timeInput)
	if err != nil {
		log.Fatal(err)
	}
	client := GetClient(addr, timeout)
	client.Run()
}
