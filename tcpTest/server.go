package main

import (
	"fmt"
	"net"
	"time"
)

func main() {

	fmt.Println("starting the server ...")
	//listener, err := net.Listen("tcp", "localhost:8832")
	tcpadd, _ := net.ResolveTCPAddr("tcp", "localhost:8832")
	listener, err := net.ListenTCP("tcp", tcpadd)
	if err != nil {
		fmt.Println("error listening ", err.Error())
		return
	}
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Println("Error accepting ", err.Error())
			continue
			conn.Close()
		}
		go joinnerHandle(conn)
	}
	listener.Close()
}

func joinnerHandle(con *net.TCPConn) {
	fmt.Println("accept client ip ", con.RemoteAddr().String())
	for {
		con.SetDeadline()
		buf := make([]byte, 4096)
		len, err := con.Read(buf)
		if err != nil {
			fmt.Println("read error ", err.Error())
			con.Close()
			fmt.Println("client closed ip ", con.RemoteAddr().String())
			break
		}
		fmt.Println("Received data: %v", string(buf[:len]))
		if len > 0 {
			s1 := "recv   " + string(buf[:len]) + "\n"
			wLen, err := con.Write([]byte(s1))
			if err != nil {
				fmt.Println("write error !", err.Error())
			} else {
				fmt.Println("write succ , write len ", wLen)
			}
			time.Sleep(3.0 * time.Second)
		}
	}
}
