package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	port := flag.Int("port", 8000, "the listening port of clock")
	flag.Parse()

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatal(err)
	}

	timeZone := os.Getenv("TZ")
	if len(timeZone) == 0 {
		timeZone = "Asia/Shanghai"
	}

	log.Printf("%s clock is listening on port %d", timeZone, *port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		go handleConn(conn, timeZone)
	}

}

func handleConn(c net.Conn, timeZone string) {
	defer c.Close()
	for {
		now, err := timeIn(time.Now(), timeZone)
		if err != nil {
			return
		}

		_, err = io.WriteString(c, fmt.Sprintf("%s time: %s", timeZone, now.Format("15:04:05\n")))
		if err != nil {
			return
		}

		time.Sleep(1 * time.Second)
	}
}

func timeIn(t time.Time, name string) (time.Time, error) {
	loc, err := time.LoadLocation(name)
	if err != nil {
		return t, err
	}

	t = t.In(loc)

	return t, nil
}
