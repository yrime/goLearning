package main

import (
    "fmt"
    "golangTest/crypto"
    "golangTest/net"
    "os"
)

func fileCrypt(filename string, key []byte) {
	file, err := os.Open(filename)
        if err != nil {
                fmt.Printf("error of read file: %v\n", err)
                return
        }
	encFile, err := os.Create(filename+"enc.bin")
	if err != nil {
	        fmt.Printf("error of create file: %s  %v\n", filename+"enc.bin",  err)
        	return
    	}
    	defer encFile.Close()
        defer file.Close()
        buffer := make([]byte, 1024)
        
        for {
                n, err := file.Read(buffer)
                if n > 0 {
                        fmt.Printf("reading %d bytes: %s\n", n, buffer[:n])
                        enc_str := crypto.Crypt(buffer[:n], key, true)
			enc_str_toFile := []byte(enc_str)
			written, err := encFile.Write(enc_str_toFile)
	                if err != nil {
            			fmt.Printf("error write to file: %v\n", err)
                		return
		        }
			fmt.Println("write bytes %v", written)

                        fmt.Println("buffer: " +  string(buffer))

			fmt.Printf("2. %%q: %q\n", enc_str)

                        dec_str := crypto.Crypt([]byte(enc_str), key, false)
                        fmt.Println ("deccryption file: " +dec_str)

                }
        
                if err != nil {
                        if err.Error() == "EOF" {
                                break 
                        }
                        fmt.Printf("error: %v\n", err)
                        break
                }
        }

}

func main(){
	var filename string
	var server bool
	var key = []byte{
                        0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07,
                        0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F,
                        0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17,
                        0x18, 0x19, 0x1A, 0x1B, 0x1C, 0x1D, 0x1E, 0x1F,
                    }

	if len(os.Args) > 1{
		filename = os.Args[1]

		if os.Args[2] == "server"{
			server = true
		}else{
			server = false
		}

		fileCrypt(filename, key)
	}else{
		str := "asdqwezxcrtyfghvbnuiojklm,.qwesdxcvfghfgh"
		data := []byte(str)
		enc_str := crypto.Crypt(data, key, true)
		fmt.Println("string: " +  string(data))
		fmt.Println ("encryption: " +enc_str)
		dec_str := crypto.Crypt([]byte(enc_str), key, false)
		fmt.Println ("deccryption: " +dec_str)

	}
	if server{
		err := net.CreateServer("127.0.0.1" , 14880)
		fmt.Println("error of creation server" + err.Error())
	}else{
		net.ConnectionToServer("127.0.0.1" , 14880)
		//fmt.Println("error of connection to server" + err.Error())
	}
}

