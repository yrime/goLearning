package net

import (
	"fmt"
	"net"
	"strconv"
//	"os"
)

func CreateServer(ip string, port int) error {
	listener, err := net.Listen("tcp", ip + ":" + strconv.Itoa(port))
	if err != nil {
		return err
	}
	defer listener.Close()

	fmt.Println("Server listening on: " + ip + ":" + strconv.Itoa(port))
	
	for {
		conn,err := listener.Accept()
		if err != nil {
			fmt.Println("Error accept: " + err.Error())
			continue
		}
		go handle_client(conn)
	}
}

func ConnectionToServer(ip string, port int) {
	conn, err := net.Dial("tcp", ip + ":" + strconv.Itoa(port))
	if err != nil {
                fmt.Println("error connection to server: ", err.Error())
        }
	defer conn.Close()
	message := "Hellow"
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("error send: ", err.Error())
	}
	
}

func handle_client(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("error ridding: ", err.Error())
	}
	fmt.Printf("Received: %s\n", string(buffer[:n]))
}
