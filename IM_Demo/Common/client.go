package Common

import (
	"net"
	"sync"
)

type Client struct {
	Udid string
	Conn *net.TCPConn
	
}

type ClientList struct {
	List map[string]*Client
	Lock sync.Mutex
}

func GetClientList() *ClientList {
	c1 := new(ClientList)
	c1.List = make(map[string]*Client)
	return c1
}

func CreateClient(key string , con net.TCPConn) *Client  {

}