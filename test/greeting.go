package test

import (
	"net"
)

type Greet struct {
	Msg string
}
func(this *Greet) SayHello(conn net.Conn){
	var r Response
	PacketResponse(&r,200,"",[]interface{}{"sure"})
	 p:= Protocol{Mark: PGreet,Function:"SayHelloResponse",DataResponse:r}
     Packet(p,conn)
}