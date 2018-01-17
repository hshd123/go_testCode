package main

import "encoding/json"
import "fmt"

type Server struct {
	ServerName string
	ServerIP   string
}

type Serverslice struct {
	Servers []Server
}

func main()  {
	
	//json - > struct
	var s1 Serverslice
	str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	json.Unmarshal([]byte(str), &s1)
	fmt.Println(s1.Servers[0].ServerIP)
	
	//struct - > json
	s2 := Server{
		ServerIP:"192.168.1.2",
		ServerName:"localhost",
	}
	b1, err := json.Marshal(s2)
	fmt.Println("ss " , string(b1))
	if err != nil {
		 fmt.Println(err.Error())
	}
}