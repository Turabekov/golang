package main

import "fmt"

func decodeMessage(key string, message string) string {
	subsMap := make(map[byte]byte)

	var alphabet byte = 'a'
	for i := 0; i < len(key); i++ {
		if string(key[i]) == " " {
			continue
		}

		_, ok := subsMap[key[i]] 
		if ok {
			continue
		}

		subsMap[key[i]] = alphabet
		alphabet++	
	}

	decodedMsg := make([]byte, len(message))

	for i := 0; i < len(message); i++ { 
		if (message[i]) == ' ' {
			decodedMsg[i] = ' '
			continue
		}

		decodedMsg[i] = subsMap[message[i]]
	}

	return string(decodedMsg)
}

func main() {
	fmt.Println(decodeMessage("the quick brown fox jumps over the lazy dog", "vkbs bs t suepuv"))
}