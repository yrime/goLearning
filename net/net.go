package net

import (
	"fmt"
	"net"
	"strconv"
	"os"
	"strings"
	"golangTest/terminal"
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
		fmt.Println("Connection from :", conn.RemoteAddr().String());
		go handle_client(conn)
	}
}

func ConnectionToServer(ip string, port int) {
	conn, err := net.Dial("tcp", ip + ":" + strconv.Itoa(port))
	if err != nil {
                fmt.Println("error connection to server: ", err.Error())
        }
	defer conn.Close()
	fmt.Printf("Connected to %s\n", conn.RemoteAddr().String())

	io := terminal.Init()
	
	go func(){
		buffer := make([]byte, 1024)
		for {
			n, _ := conn.Read(buffer)
        	        if err != nil {
                	        fmt.Println("error ridding: ", err.Error())
                	}
                	fmt.Printf("From: %s  message: %s \n", conn.RemoteAddr().String(), string(buffer[:n]))
		}
	}()

	for {
		select{
                        case val, b := <-io:
                                if b {
                                        _, err = conn.Write([]byte(val))
                                }
                	default:
				continue
		}
	}
}

func handle_client(conn net.Conn) {
	defer conn.Close()
	io := terminal.Init()

	go func() {
                buffer := make([]byte, 1024)
		for {
	                n, err := conn.Read(buffer)
        	        if err != nil {
                	        fmt.Println("error readding: ", err.Error())
	                }                
        	        if n > 0 {
				fmt.Println("f " + strings.TrimSpace(string(buffer[:n])))
                	        if strings.TrimSpace(string(buffer[:n])) == "exit" {
					fmt.Println("stopped")
                        	        conn.Close()
					os.Exit(0)
	                        }
        	                fmt.Printf("From: %s  Message: %s\n", conn.RemoteAddr().String(), string(buffer[:n]))
                	}
		}
        }()


	for {
		select{
			case val, b := <-io:
				if b {
					_, _ = conn.Write([]byte(val))
				}
			default:
				continue
		}
	}
}
