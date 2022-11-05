package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	//Listen
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}

	//Capable of handling mulitple requrest simultaneuously
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Panic(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	err := conn.SetDeadline(time.Now().Add(15 * time.Second)) //add 10 sec deadLine for the conn
	if err != nil {
		log.Println("Connection close dues to deadline reached..:(")
	}
	sc := bufio.NewScanner(conn)
	for sc.Scan() {
		ln := sc.Text()
		fmt.Println(ln)
	}
	defer conn.Close()
}
