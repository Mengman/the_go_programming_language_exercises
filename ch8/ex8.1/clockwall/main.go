package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
)

type _Clock struct {
	name string
	host string
	port int64
}

func (c _Clock) hostName() string {
	return fmt.Sprintf("%s:%d", c.host, c.port)
}

func main() {
	var clocks []_Clock
	for _, arg := range os.Args[1:] {
		clock, err := createClock(arg)
		if err != nil {
			log.Println(err)
			continue
		}

		clocks = append(clocks, clock)
	}

	timeChans := make(map[string]chan string)
	for _, clock := range clocks {
		timeChans[clock.name] = make(chan string)
		go readTime(clock, timeChans[clock.name])
	}

	displayTime(clocks, timeChans)
}

func createClock(inputStr string) (clock _Clock, err error) {
	inputFormatError := fmt.Errorf("incorrect clock input expected: ClockName=host:port")
	args := strings.Split(inputStr, "=")
	if len(args) != 2 {
		err = inputFormatError
		return
	}

	name := args[0]
	hostInfo := strings.Split(args[1], ":")
	if len(hostInfo) != 2 {
		err = inputFormatError
		return
	}

	host := hostInfo[0]
	port, err := strconv.ParseInt(hostInfo[1], 10, 64)
	if err != nil {
		return
	}

	clock = _Clock{
		name: name,
		host: host,
		port: port,
	}

	return
}

func readTime(clock _Clock, timeChan chan<- string) {
	conn, err := net.Dial("tcp", clock.hostName())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	input := bufio.NewScanner(conn)
	for input.Scan() {
		timeChan <- input.Text()
	}
}

func displayTime(clocks []_Clock, timeChans map[string]chan string) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.AlignRight|tabwriter.Debug)

	for {
		timeTable := make(map[string]string)
		for clockName, timeChan := range timeChans {
			timeStrs := strings.Split(<-timeChan, "time:")
			timeTable[clockName] = strings.TrimSpace(timeStrs[1])
		}

		sbTitle := strings.Builder{}
		sbRecord := strings.Builder{}
		for _, clock := range clocks {

			sbTitle.WriteString(clock.name)
			sbTitle.WriteString("\t")

			sbRecord.WriteString(timeTable[clock.name])
			sbRecord.WriteString("\t")
		}

		fmt.Fprintln(w, sbTitle.String())
		fmt.Fprintln(w, sbRecord.String())
		w.Flush()
	}
}
