// WHAT
package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net"
	"time"
)

var filedataCh = make(chan []byte)

func fileMon(filename string) {
outerLoop:
	for {
		filedata, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Fatal(err)
		}

		timeout := time.After(10 * time.Second)
		for {
			select {
			case filedataCh <- filedata:
			case <-timeout:
				continue outerLoop
			}
		}
	}
}

func main() {
	filename := flag.String("filename", "", "Filename to be dumping")
	listen := flag.String("listen", ":8080", "Address to listen on")
	flag.Parse()

	if *filename == "" {
		log.Fatal("No filename specified")
	}

	go fileMon(*filename)

	log.Printf("Listening on %s", *listen)
	l, err := net.Listen("tcp", *listen)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	log.Printf("Connection from %s", conn.RemoteAddr())
	conn.Write(<-filedataCh)
	conn.Close()
}
