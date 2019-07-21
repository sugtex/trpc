package test

import (
	"encoding/binary"
	jsoniter "github.com/json-iterator/go"
	"net"
	"reflect"
)

//封包
func Packet(c Protocol,conn net.Conn){
	var item Shell
	item.Header ="DDF-CXK"
	item.Kernel =c
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	s,_:=json.Marshal(item)
	//整个包长度
	by:=make([]byte,4)
	binary.LittleEndian.PutUint32(by, uint32(len(s)))
	res:=append(by,s...)
	conn.Write(res)
}
//解包
func Analysis(conn net.Conn,c []byte)  {
	var shell Shell
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	json.Unmarshal(c,&shell)
	//验证头部
	if checkLegality(shell.Header)==false{
		conn.Close()
	}
	data:=&shell.Kernel.DataRequest                                 //得到数据，map[string]interface{}类型（json没法自适应到结构体上）
	temp,_:=json.Marshal(data)                                      //序列化成[]byte数组
	entity:=reflect.New(StructGroup[shell.Kernel.Mark]).Interface() //实例化出目标结构体，指针
	json.Unmarshal(temp,entity)                                     //反序列化进结构体
	//反射调用函数
	v:=reflect.ValueOf(entity)
	fun:=v.MethodByName(shell.Kernel.Function)
    pg:=[]reflect.Value{reflect.ValueOf(conn)}
    fun.Call(pg)
}
//确认合法性
func checkLegality(h string) bool{
	res:=true
	if h!="DDF-CXK"{
		res=false
	}
	return res
}
//封装回复模板
func PacketResponse(r *Response,s int,m string,d []interface{}){
    	r.Status=s
    	r.Msg=m
    	r.Data=d
}

