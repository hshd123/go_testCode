package IM

import (
	"IM_Demo/Common"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/satori/go.uuid"
)

type TcpServer struct {
	listener     *net.TCPListener      // 服务器监听器
	clients      *Common.ClientList    //客户端列表
	joinSniffer  chan *net.TCPConn     //检测到的connect
	quiteSniffer chan *Common.Client   // 客户端退出
	recvSniffer  chan Common.IMMessage // 接收到的消息
}

func StartIMServer() {
	log.Println("im server 正在器启动中。 。。。。 ")
	server := new(TcpServer)
	server.clients = Common.GetClientList()
	server.joinSniffer = make(chan *net.TCPConn)
	server.quiteSniffer = make(chan *Common.Client)
	server.recvSniffer = make(chan Common.IMMessage)

	//先关闭以前的服务
	server.closedService()
	//准备接收消息
	server.listen()

	//启动服务
	server.start()
}

func (s *TcpServer) listen() {
	go func() {
		for {
			select {
			case msg := <-s.recvSniffer:
				fmt.Println(msg)
			//有一个新的连接过来
			case conn := <-s.joinSniffer:
				s.joinClient(conn)
			case quite := <-s.quiteSniffer:
				fmt.Println(quite)
			}
		}
	}()
}

func (s *TcpServer) closedService() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		sig := <-c
		log.Printf("captured %v, stopping profiler and exiting..", sig)
		// 清除客户端连接
		for _, v := range s.clients.List {
			s.quitHandler(v)
		}
		// 退出
		os.Exit(1)
	}()
}

func (s *TcpServer) joinClient(conn *net.TCPConn) {
	c1 := new(Common.Client)
	id := uuid.NewV4().String()
	c1.Udid = id
	c1.Conn = conn
	s.clients.Lock.Lock()
	//加入到列表中
	s.clients.List[id] = c1

	//读取数据 ， 发送数据
}

func (s *TcpServer) start() {
	fmt.Println("starting the server ... port 8832")
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
			fmt.Println("accept error ", conn.RemoteAddr().String())
			continue
			conn.Close()
		}
		fmt.Println("accept new client ", conn.RemoteAddr().String())
		s.listener = listener
		s.joinSniffer <- conn
	}
	defer listener.Close()
}

func (s *TcpServer) quitHandler(clent *Common.Client) {

}
