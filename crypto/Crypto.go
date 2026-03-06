package crypto

import (
    "fmt"
)

const globalSize = 4

func pkcs7Pad(data []byte, blockSize int) []byte {
    padding := blockSize - len(data)%blockSize
    if padding == 0 {
        padding = blockSize
    }

    result := make([]byte, len(data)+padding)
    copy(result, data)

    for i := len(data); i < len(result); i++ {
        result[i] = byte(padding)
    }

    return result
}

func pkcs7Unpad(data []byte) ([]byte, error) {
    if len(data) == 0 {
        return nil, fmt.Errorf("nul data")
    }

    padding := int(data[len(data)-1])
    if padding > len(data) {
        return nil, fmt.Errorf("pas correct padding")
    }

    for i := len(data) - padding; i < len(data); i++ {
        if data[i] != byte(padding) {
            return nil, fmt.Errorf("pas correct padding")
        } 
    }

    return data[:len(data)-padding], nil
} 

func xor(str1 [globalSize]byte, str2 [globalSize]byte) [globalSize]byte{
	var out [globalSize]byte
	for i:= 0; i < globalSize; i++ {
		out[i] = str1[i] ^ str2[i]
	}
	return out
}

func genKey(key []byte, i int) [globalSize]byte{
	var out[globalSize]byte
	var step = (len(key)-1 + int(key[i])) * 13
	for j := 0; j < globalSize; j++ {
		out[j] = key[(i + j)%len(key)] ^ key[(step + j)%len(key)]
	}
	return out
}


func Crypt(str []byte, key []byte, b bool) string{
	out := make([]byte, 0)
	var key_session [globalSize]byte
 	key_session = [globalSize]byte(key[:globalSize])
	var block [globalSize]byte
	var padded []byte
	if b{
		padded = pkcs7Pad(str, globalSize)
	}else{
		padded = str
	}

	for i := 0; i < len(padded); i += globalSize {
		block = [globalSize]byte(padded[i: i+globalSize])
		key_session = xor(key_session, genKey(key, i%len(key)))
		newBlock := xor(key_session, block)
		out = append(out, newBlock[:]...)
	}
	if !b{
		unpadded, _ := pkcs7Unpad(out)
		out = unpadded
	}
	return string(out)
}

