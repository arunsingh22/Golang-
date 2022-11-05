package main

import (
	"io"
	"net"
)

func main() {

	//Step 1: Listen
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer li.Close()

	//Step 2: Accept the conn
	for {
		conn, err := li.Accept()
		if err != nil {
			panic(err)
		}

		//Step3: Write back and close.
		io.WriteString(conn, "\nWeclome Client to TCP server..how may i help you :)")
		conn.Close()
	}
}
