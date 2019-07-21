package test

import (
	"encoding/binary"
	"fmt"
	"net"
)

func InitServer(){

	//监听端口
	listener,_:=net.Listen("tcp","127.0.0.1:7879")
	defer listener.Close()
	//客户端连接
	for {
		conn,err:=listener.Accept()
		if err!=nil{

			continue
		}
		//处理客户端消息
		go handleMessage(conn)
	}
}
//处理客户端消息
func handleMessage(conn net.Conn) {
	buf:=make([]byte,1024)//最大接受区
	var cache []byte//剩余
	for{
		//每次接收的都放在接受区中
		length,err:=conn.Read(buf)
		if err!=nil{
			fmt.Println("读取失败",err)
			conn.Close()
			break
		}
		cache=append(cache,buf[:length]...)
		all:=len(cache)
		//判断缓存区是否空和分包
		if all==0||all<=4{
			continue
		}
		//读取单体包
		slength:=binary.LittleEndian.Uint32(cache[:4])
		//判断是否能组成一个包体
		if uint32(all)<4+slength{
			continue
		}
		content:=cache[4:4+slength]
		Analysis(conn,content)
		cache=cache[4+slength:]
	}
}
