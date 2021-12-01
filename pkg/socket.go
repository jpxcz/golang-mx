package engineConnection

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

// Creates the tcp socket connection to an engine
func GenerateEngineConn(url string) (conn net.Conn, err error) {
	conn, err = net.Dial("tcp", url)
	if err != nil {
		return nil, err
	}

	defer conn.Close()
	return conn, nil
}

// Read the messages incoming from the tcp connection and pushing them into the channel
// to read the right message after delimiter has been called
func MessagesReader(r io.Reader, channel chan<- string) {
	bufReader := bufio.NewReader(r)
	for {
		msg, err := bufReader.ReadString(';')
		if err != nil {
			log.Printf("error checking the message %s", err)
		}
		channel <- msg
	}
}

// Write to the TCP engine
func MessageWritter(w io.Writer, s string) error {
	fmt.Printf("writting string to engine %s", s)
	buff := []byte(s)
	_, err := w.Write(buff)
	return err
}

// Generate the Client and read messages
func Client(url string) {
	read := make(chan string)
	tcpSocket, err := GenerateEngineConn(url)
	if err != nil {
		fmt.Print("error on client!")
	}
	// write here the messages

	go MessagesReader(tcpSocket, read)

	for {
		msg := <-read
		fmt.Print(msg)
	}
}
