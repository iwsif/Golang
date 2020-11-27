package main 

import (
	"encoding/hex"
	"time"
	"io"
    "fmt"
    "crypto/cipher"
    "crypto/aes"
    "crypto/rand"
	"log"
)

func encrypt(word_to_encrypt string) {
	key := make([]byte,32)
	io.ReadFull(rand.Reader,key)
	encoded_key := hex.EncodeToString(key)
	value_to_encrypt := []byte(word_to_encrypt)
	aes_key,_:= aes.NewCipher(key)
	algo,err := cipher.NewGCM(aes_key)
	if err != nil {
		log.Fatal(err)
	}
	nonce := make([]byte,algo.NonceSize())
	
	text := algo.Seal(nonce,nonce,value_to_encrypt,nil)
	encoded_text := hex.EncodeToString(text)

	fmt.Println("Key:" + encoded_key)
	fmt.Println("Text:" + encoded_text)
	time.Sleep(time.Second*2)
}

func decrypt(encoded_string string,encoded_key string) {
	
	encrypted_text,_:= hex.DecodeString(encoded_string)
	decoded_key,_:= hex.DecodeString(encoded_key)

	key,_:= aes.NewCipher(decoded_key)
	algo,_ := cipher.NewGCM(key)

	nonce,new_text := encrypted_text[:algo.NonceSize()],encrypted_text[algo.NonceSize():]

	value,_:= algo.Open(nil,nonce,new_text,nil)
	decoded_value := hex.EncodeToString(value)
	fmt.Println(decoded_value)
	final_value := []byte(decoded_value)
	fmt.Println("Decode this hex and you have your plain text >> " + hex.EncodeToString(final_value))
}

func main() {
    var answer string 
    fmt.Println("-----Welcome-----")
    fmt.Println("This is an encryption program(aes-algorithm)")
    fmt.Println("You want to encrypt or decrypt something?")
    fmt.Scanf("%s",&answer)
    if answer == "encrypt" {
		var word_to_encrypt string
        fmt.Println("Type in what you want to encrypt")
		fmt.Scanf("%s",&word_to_encrypt)
        fmt.Println("Starting the process...")
		time.Sleep(time.Second*2)
		fmt.Println("Done..")
		time.Sleep(time.Second*2)
        encrypt(word_to_encrypt)
    }else if answer == "decrypt" {
		var key,text string
		fmt.Println("Text:")
		fmt.Scanf("%s",&text)
		fmt.Println("Key:")
		fmt.Scanf("%s",&key)
        fmt.Println("Starting the process")
		time.Sleep(time.Second*2)
		decrypt(text,key)
    }else {
        fmt.Println("Wrong answer,type encrypt or decrypt")
	}
}


