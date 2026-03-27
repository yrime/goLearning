package main

import (
	"net"
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
	"db/db"
)

type uneConn struct{
	conn net.Conn
	user db.User
	status bool
}


var dbConn = []uneConn{}


func initStdio() chan string{
        ioChan := make(chan string)
        go reader(ioChan)
        return ioChan
}

func reader(ioChan chan string){
        re := bufio.NewReader(os.Stdin)
        for {
                text, _ := re.ReadString('\n')
                text = strings.TrimSpace(text)
                ioChan <- text
        }
}

func registry(l string, pwd string, n string) string{
	return db.AddUser(l, pwd, n)
}

func auth(l string, pwd string) bool {
	return db.CheckUser(l, pwd)
}

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
		dbConn = append(dbConn, uneConn{conn:conn, status:false})
		go handle_client(dbConn)
	}
}

func handle_client(dbConn uneConn) {
	conn := dbConn.conn
	status := dbConn.status
	defer conn.Close()
//	io := initStdio()
	ioNet := make(chan string)

	go func(ioNet chan string) {
                buffer := make([]byte, 1024)
		for {
	                n, err := conn.Read(buffer)
        	        if err != nil {
                	        fmt.Println("error readding: ", err.Error())
	                }                
        	        if n > 0 {
				ioNet <- string(buffer[:n])
                	}
		}
        }(ioNet)


	for {
		select{
			case val, b := <-ioNet:
				if !status {
					_, _ = conn.Write([]byte("u dont authorizate,\n\
for auth send message as Auth{[login], [password],\n\
for registy send message as Reg{[login], [password], [u name] "))
				}
			default:
				continue
		}
	}
}

func ConnectionToServer(ip string, port int) {
	conn, err := net.Dial("tcp", ip + ":" + strconv.Itoa(port))
	if err != nil {
                fmt.Println("error connection to server: ", err.Error())
        }
	defer conn.Close()
	fmt.Printf("Connected to %s\n", conn.RemoteAddr().String())

	io := initStdio()
	
	go func(){
		buffer := make([]byte, 1024)
		for {
			n, _ := conn.Read(buffer)
        	        if err != nil {
                	        fmt.Println("error readding: ", err.Error())
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


func main() {

	var server bool

	if len(os.Args) > 1 {
		if os.Args[1] == "server" {
			server = true
		}else{
			server = false
		}
	}else{
		server = false
	}

	if server {
		err := CreateServer("127.0.0.1" , 14880)
		fmt.Println("error of creation server" + err.Error())
	}else{
		ConnectionToServer("127.0.0.1" , 14880)
	}
}
