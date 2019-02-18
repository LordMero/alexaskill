package main

import (
	"fmt"
	"log"
	"net"
)

func main() {

	l, e := net.Listen("tcp",":8080")

	defer l.Close()
	if e != nil{
		log.Panic(e)
	}

	for {c, e := l.Accept()
		if e != nil{
			log.Panic(e)
		}
	fmt.Fprint(c, "\n Hello Graziano\n\n")
	c.Close()
}





}
