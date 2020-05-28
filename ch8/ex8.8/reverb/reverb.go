package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8000, "listen port")

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		log.Println("a new connection established")
		go handleConn(conn)
	}
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	timer := time.After(10 * time.Second)
	shout := make(chan string)

	go func() {
		input := bufio.NewScanner(c)
		for input.Scan() {
			shout <- input.Text()
		}
	}()

	for {
		select {
		case <-timer:
			log.Println("no shout in 10 seconds, disconnect client")
			close(shout)
			c.Close()
			return
		case s := <-shout:
			go echo(c, s, 1*time.Second)
		}
	}
	close(shout)
	c.Close()
}
