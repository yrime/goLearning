package terminal

import (
	"bufio"
	_"fmt"
	"os"
	"strings"
)

func Init() chan string{
	ioChan := make(chan string)
	go reader(ioChan)
	/*
	for msg := range ioChan {
		fmt.Println("get: " + msg);
	}
*/	
	return ioChan
}

func reader(ioChan chan string){
	re := bufio.NewReader(os.Stdin)
	for {
		text, _ := re.ReadString('\n')
		text = strings.TrimSpace(text)
	//	if text == "exit" {
	//		close(ioChan)
	//		break;
	//	}
		ioChan <- text
	}
}
