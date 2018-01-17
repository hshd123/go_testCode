package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func main() {
	con, err := net.Dial("tcp", "localhost:8832")
	if err != nil {
		fmt.Println("connect error ", err.Error())
	}

	for {
		time.Sleep(3.0)
		s := "client write " + string(time.Now().String()+"\n")
		con.Write([]byte(s))
		time.Sleep(3.0 * time.Second)
		
		r1 := bufio.NewReader(con)
		b1 := make([]byte, 4096)
		r1.Read(b1)
		if l1 := len(b1); l1 > 0 {
			fmt.Println("read succ ", string(b1))
		}
	}

}
