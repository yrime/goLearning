package main

import (
	"fmt"
	"os"
	"strings"
	"bufio"
)

func reader(ioChan chan string, d chan bool){
        re := bufio.NewReader(os.Stdin)
        for {
		fmt.Printf("in done1\n")
                text, _ := re.ReadString('\n') //block operation
                text = strings.TrimSpace(text)
                if text == "exit" {
                      // close(ioChan)
			d <- true
                       	break;
                }
                ioChan <- text
        }
}

func print(ioChan chan string, d chan bool){
	for {
		fmt.Printf("in done2\n")
		val, _/*b*/ := <-ioChan  //block operation
	//	if b{
			fmt.Printf("f: %s\n", val)
	//	}
		if val ==  "exit" {
		       //close(ioChan)
			d <- true
                        return;
                }

	}
}

func multiMatrix(){
	m1 := [2][2]int{{0, 1}, {1, 0}}
	m2 := [1][2]int{{2,3}}
	m3 := [1][2]int{}

	for i := 0; i < len(m3); i++ {
		for j := 0; j < len(m3[0]); j++ {
			m3[i][j] = func(i int, j int) int {
					c := 0
					for k := 0; k < len(m2[0]); k++ {
						c += (m2[i][k] * m1[k][j])
					}
					return c
				}(i, j)
		}
	}
	for i := 0; i < len(m3); i++ {
		for j := 0; j < len(m3[0]); j++ {
			fmt.Printf("m3[%d][%d] = %d ", i, j,  m3[i][j])
		}
		fmt.Printf("\n")
	}
}

func main(){
	fmt.Printf("hellow\n")
	
	multiMatrix()


	done1 := make(chan bool)
	done2 := make(chan bool)
	ioChan := make(chan string)
	b := false
	defer close(ioChan)
	go reader(ioChan, done1)
	go print(ioChan, done2)
	for{
		select{
			case <-done1:
				fmt.Printf("done1\n")
				b = true
				return
			case <-done2:
				fmt.Printf("done2\n")
				b = true
				return
			default:
				if b {
					fmt.Printf("d")
				}
				continue
		}
	}
	
/*
	for {
                select{
                        case val, b := <-ioChan:
                                if b {
                                       fmt.Printf("f: %s\n", val)
                                }
                        default:
                                continue
                }
        }
*/
}

