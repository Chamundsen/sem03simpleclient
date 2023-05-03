package main

import (
	"net"
	"log"
	"os"
"github.com/Chamundsen/is105sem03/mycrypt"
)

func main() {
	conn, err := net.Dial("tcp", "172.17.0.5:6060")
	if err != nil {
		log.Fatal(err)
	}
    
	log.Println("os.Args[1] = ", os.Args[1])


 	  kryptertMelding := mycrypt.Krypter([]rune(os.Args[1]), 
mycrypt.ALF_SEM03, 4)
    

    

    _, err = conn.Write([]byte(string(kryptertMelding)))
    if err != nil {
        log.Fatal(err)
    }
	
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	} 
	kryptertRespons := buf[:n]
	log.Printf("kryptert respons: %s", kryptertRespons)
}
